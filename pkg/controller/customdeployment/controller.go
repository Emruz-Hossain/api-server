
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

package customdeployment

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"api-server/pkg/apis/customdp/v1alpha1"
	"api-server/pkg/controller/sharedinformers"
	listers "api-server/pkg/client/listers_generated/customdp/v1alpha1"
)

// +controller:group=customdp,version=v1alpha1,kind=CustomDeployment,resource=customdeployments
type CustomDeploymentControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about CustomDeployment
	lister listers.CustomDeploymentLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *CustomDeploymentControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing customdeployments labels
	c.lister = arguments.GetSharedInformers().Factory.Customdp().V1alpha1().CustomDeployments().Lister()
}

// Reconcile handles enqueued messages
func (c *CustomDeploymentControllerImpl) Reconcile(u *v1alpha1.CustomDeployment) error {
	// Implement controller logic here
	log.Printf("Running reconcile CustomDeployment for %s\n", u.Name)
	return nil
}

func (c *CustomDeploymentControllerImpl) Get(namespace, name string) (*v1alpha1.CustomDeployment, error) {
	return c.lister.CustomDeployments(namespace).Get(name)
}
