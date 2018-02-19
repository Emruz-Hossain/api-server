
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

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"api-server/pkg/apis/customdp"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomDeployment
// +k8s:openapi-gen=true
// +resource:path=customdeployments,strategy=CustomDeploymentStrategy
type CustomDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomDeploymentSpec   `json:"spec,omitempty"`
	Status CustomDeploymentStatus `json:"status,omitempty"`
}

// CustomDeploymentSpec defines the desired state of CustomDeployment
type CustomDeploymentSpec struct {
}

// CustomDeploymentStatus defines the observed state of CustomDeployment
type CustomDeploymentStatus struct {
}

// Validate checks that an instance of CustomDeployment is well formed
func (CustomDeploymentStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*customdp.CustomDeployment)
	log.Printf("Validating fields for CustomDeployment %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default CustomDeployment field values
func (CustomDeploymentSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*CustomDeployment)
	// set default field values here
	log.Printf("Defaulting fields for CustomDeployment %s\n", obj.Name)
}
