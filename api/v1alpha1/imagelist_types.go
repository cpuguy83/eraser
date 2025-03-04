/*
Copyright 2021.
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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ImageListSpec defines the desired state of ImageList
type ImageListSpec struct {
	// The list of vulnerable images to delete if non-running.
	Images []string `json:"images"`
}

type NodeResult struct {
	// node name of image results
	Name string `json:"name"`
	// list of images searched and their results
	Images []NodeCleanUpResult `json:"images"`
}

// ImageListStatus defines the observed state of ImageList
type ImageListStatus struct {
	// Information when the job was completed.
	Timestamp *metav1.Time `json:"timestamp"`

	// list of node names and their images with respective results
	Node []NodeResult `json:"node"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ImageList is the Schema for the imagelists API
type ImageList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImageListSpec   `json:"spec,omitempty"`
	Status ImageListStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ImageListList contains a list of ImageList
type ImageListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageList `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImageList{}, &ImageListList{})
}
