package v1beta1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type OpenstackVolumePopulator struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec OpenstackVolumePopulatorSpec `json:"spec"`
}

type OpenstackVolumePopulatorSpec struct {
	IdentityURL string `json:"identityUrl"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	ImageID     string `json:"imageId"`
	Region      string `json:"region"`
	Domain      string `json:"domain"`
	Tenant      string `json:"tenant"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type OpenstackVolumePopulatorList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`
	Items         []OpenstackVolumePopulator `json:"items"`
}
