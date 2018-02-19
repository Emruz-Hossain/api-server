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

package customreplicationcontroller

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"api-server/pkg/apis/customrc/v1alpha1"
	listers "api-server/pkg/client/listers_generated/customrc/v1alpha1"
	"api-server/pkg/controller/sharedinformers"
)

// +controller:group=customrc,version=v1alpha1,kind=CustomReplicationController,resource=customreplicationcontrollers
type CustomReplicationControllerControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about CustomReplicationController
	lister listers.CustomReplicationControllerLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *CustomReplicationControllerControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing customreplicationcontrollers labels
	c.lister = arguments.GetSharedInformers().Factory.Customrc().V1alpha1().CustomReplicationControllers().Lister()
}

// Reconcile handles enqueued messages
func (c *CustomReplicationControllerControllerImpl) Reconcile(u *v1alpha1.CustomReplicationController) error {
	// Implement controller logic here
	log.Printf("Running reconcile CustomReplicationController for %s\n", u.Name)

	crc := u.DeepCopy()

	log.Println("Required: %v		| Available: %v	| Creating: %v	| Terminating: %v", crc.Spec.Replicas, crc.Status.AvailableReplicas, crc.Status.CreatingReplicas, crc.Status.TerminatingReplicas)
	if crc.Status.AvailableReplicas+crc.Status.CreatingReplicas < crc.Spec.Replicas {
		log.Println("We need to create new pod.")
	}
	return nil
}

func (c *CustomReplicationControllerControllerImpl) Get(namespace, name string) (*v1alpha1.CustomReplicationController, error) {
	return c.lister.CustomReplicationControllers(namespace).Get(name)
}
