/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"forklift.konveyor.io/os-populator/pkg/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/imagedata"
	populator_machinery "github.com/kubernetes-csi/lib-volume-populator/populator-machinery"
)

const (
	prefix     = "forklift.konveyor.io"
	mountPath  = "/mnt/"
	devicePath = "/dev/block"
)

var version = "unknown"

func main() {
	var (
		mode             string
		identityEndpoint string
		username         string
		password         string
		region           string
		domainName       string
		tenantName       string
		imageID          string

		fileName     string
		httpEndpoint string
		metricsPath  string
		masterURL    string
		kubeconfig   string
		imageName    string
		showVersion  bool
		namespace    string
	)
	klog.InitFlags(nil)
	// Main arg
	flag.StringVar(&mode, "mode", "", "Mode to run in (controller, populate)")
	flag.StringVar(&identityEndpoint, "endpoint", "", "endpoint URL (https://openstack.example.com:5000/v2.0)")
	flag.StringVar(&username, "username", "", "Openstack username")
	flag.StringVar(&password, "password", "", "Openstack password")
	flag.StringVar(&region, "region", "", "Openstack region")
	flag.StringVar(&domainName, "domain", "", "Openstack domain")
	flag.StringVar(&tenantName, "tenant", "", "Openstack tenant")

	flag.StringVar(&imageID, "image-id", "", "Openstack image ID")
	flag.StringVar(&fileName, "file-name", "", "Filename to populate")

	// Populate args

	// Controller args
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&imageName, "image-name", "", "Image to use for populating")
	// Metrics args
	flag.StringVar(&httpEndpoint, "http-endpoint", "", "The TCP network address where the HTTP server for diagnostics, including metrics and leader election health check, will listen (example: `:8080`). The default is empty string, which means the server is disabled.")
	flag.StringVar(&metricsPath, "metrics-path", "/metrics", "The HTTP path where prometheus metrics will be exposed. Default is `/metrics`.")
	// Other args
	flag.BoolVar(&showVersion, "version", false, "display the version string")
	flag.StringVar(&namespace, "namespace", "konveyor-forklift", "Namespace to deploy controller")
	flag.Parse()

	if showVersion {
		fmt.Println(os.Args[0], version)
		os.Exit(0)
	}

	switch mode {
	case "controller":
		const (
			groupName  = "forklift.konveyor.io"
			apiVersion = "v1beta1"
			kind       = "OpenstackVolumePopulator"
			resource   = "openstackvolumepopulators"
		)
		var (
			gk  = schema.GroupKind{Group: groupName, Kind: kind}
			gvr = schema.GroupVersionResource{Group: groupName, Version: apiVersion, Resource: resource}
		)
		populator_machinery.RunController(masterURL, kubeconfig, imageName, httpEndpoint, metricsPath,
			namespace, prefix, gk, gvr, mountPath, devicePath, getPopulatorPodArgs)
	case "populate":
		populate(fileName, identityEndpoint, username, password, region, domainName, tenantName, imageID)
	default:
		klog.Fatalf("Invalid mode: %s", mode)
	}
}

func getPopulatorPodArgs(rawBlock bool, u *unstructured.Unstructured) ([]string, error) {
	var openstackPopulator v1beta1.OpenstackVolumePopulator
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &openstackPopulator)
	if nil != err {
		return nil, err
	}
	args := []string{"--mode=populate"}
	if rawBlock {
		args = append(args, "--file-name="+devicePath)
	} else {
		args = append(args, "--file-name="+mountPath+"+disk.img")
	}

	args = append(args, "--endpoint="+openstackPopulator.Spec.IdentityURL)
	args = append(args, "--image-id="+openstackPopulator.Spec.ImageID)
	args = append(args, "--username="+openstackPopulator.Spec.Username)
	args = append(args, "--password="+openstackPopulator.Spec.Password)
	args = append(args, "--region="+openstackPopulator.Spec.Region)
	args = append(args, "--domain="+openstackPopulator.Spec.Domain)
	args = append(args, "--tenant="+openstackPopulator.Spec.Tenant)

	return args, nil
}

func populate(fileName, endpoint, username, password, region, domainName, tenantName, imageID string) {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: endpoint,
		Username:         username,
		Password:         password,
		DomainName:       domainName,
		TenantName:       tenantName,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		klog.Fatal(err)
	}

	imageService, err := openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{Region: region})
	if err != nil {
		klog.Fatal(err)
	}

	result, err := imagedata.Download(imageService, imageID).Extract()
	if err != nil {
		klog.Fatal(err)
	}

	// TODO do bufferred reading
	content, err := ioutil.ReadAll(result)
	if err != nil {
		klog.Fatal(err)
	}
	if strings.HasSuffix(fileName, "disk.img") {
		f, err := os.Create(fileName)
		if err != nil {
			klog.Fatal(err)
		}
		defer f.Close()

		_, err = f.Write(content)
		if err != nil {
			klog.Fatal(err)
		}
	} else {
		f, err := os.Open(fileName)
		if err != nil {
			klog.Fatal(err)
		}
		defer f.Close()

		_, err = f.Write(content)
		klog.Fatal(err)
	}
}
