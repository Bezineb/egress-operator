/*

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

package v1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExternalServiceSpec defines the desired state of ExternalService
type ExternalServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// DnsName is a DNS name target for the external service
	DnsName string `json:"dnsName,omitempty"`

	// Ports is a list of ports on which the external service may be called
	Ports []ExternalServicePort `json:"ports,omitempty"`

	// Replicas is the number of gateways to run. Defaults to 3
	// +optional
	Replicas int32 `json:"replicas,omitempty"`

	// If true, add a `egress.monzo.com/hijack-dns: true` label to produced Service objects
	// CoreDNS can watch this label and decide to rewrite DnsName -> clusterIP
	HijackDns bool `json:"hijackDns,omitempty"`
}

type ExternalServicePort struct {
	// The protocol (TCP or UDP) which traffic must match. If not specified, this
	// field defaults to TCP.
	// +optional
	Protocol *v1.Protocol `json:"protocol,omitempty"`

	// The port on the given protocol.
	Port int32 `json:"port,omitempty"`
}

// ExternalServiceStatus defines the observed state of ExternalService
type ExternalServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster

// ExternalService is the Schema for the externalservices API
type ExternalService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExternalServiceSpec   `json:"spec,omitempty"`
	Status ExternalServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalServiceList contains a list of ExternalService
type ExternalServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExternalService{}, &ExternalServiceList{})
}
