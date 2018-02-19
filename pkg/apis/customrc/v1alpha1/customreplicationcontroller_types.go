
  /*
  Copyright 2017 The Kubernetes Authors.

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
	"log"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"api-server/pkg/apis/customrc"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomReplicationController
// +k8s:openapi-gen=true
// +resource:path=customreplicationcontrollers,strategy=CustomReplicationControllerStrategy
type CustomReplicationController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomReplicationControllerSpec   `json:"spec,omitempty"`
	Status CustomReplicationControllerStatus `json:"status,omitempty"`
}

// CustomReplicationControllerSpec defines the desired state of CustomReplicationController
type CustomReplicationControllerSpec struct {
	Replicas int32 `json:"replicas"`
	Template v1.PodTemplate `json:"template"`
}

// CustomReplicationControllerStatus defines the observed state of CustomReplicationController
type CustomReplicationControllerStatus struct {
	AvailableReplicas int32 `json:"available_replicas"`
	CreatingReplicas int32	`json:"creating_replicas"`
	TerminatingReplicas int32 `json:"terminating_replicas"`
}

// Validate checks that an instance of CustomReplicationController is well formed
func (CustomReplicationControllerStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*customrc.CustomReplicationController)
	log.Printf("Validating fields for CustomReplicationController %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default CustomReplicationController field values
func (CustomReplicationControllerSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*CustomReplicationController)
	// set default field values here
	log.Printf("Defaulting fields for CustomReplicationController %s\n", obj.Name)
}
