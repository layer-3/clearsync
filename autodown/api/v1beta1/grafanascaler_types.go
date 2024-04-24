/*
Copyright 2024.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GrafanaScalerSpec defines the desired state of GrafanaScaler
type GrafanaScalerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// DeploymentName is the name of the Deployment to scale
	DeploymentName string `json:"deploymentName"`
}

// GrafanaScalerStatus defines the observed state of GrafanaScaler
type GrafanaScalerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Scaled bool `json:"scaled"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GrafanaScaler is the Schema for the grafanascalers API
type GrafanaScaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GrafanaScalerSpec   `json:"spec,omitempty"`
	Status GrafanaScalerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GrafanaScalerList contains a list of GrafanaScaler
type GrafanaScalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GrafanaScaler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GrafanaScaler{}, &GrafanaScalerList{})
}
