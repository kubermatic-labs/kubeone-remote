// Copyright 2020 The KubeOne-Remote Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package remote provides the functions needed on the controller
// node to execute the different commands for KubeOne and the other
// tools.
package remote

import (
	"fmt"

	"github.com/kubermatic-labs/kubeone-remote/pkg/config"
	"github.com/kubermatic-labs/kubeone-remote/pkg/remote/setup"
)

// Setup performs all needed steps for installation and configuration
// on the remote node.
func Setup(cfg config.Config) error {
	err := setup.InstallKubeOneRemote(cfg)
	if err != nil {
		return fmt.Errorf("setup failed installing KubeOne Remote: %v", err)
	}
	err = setup.CheckoutBranches(cfg)
	if err != nil {
		return fmt.Errorf("setup failed checking out needed branches: %v", err)
	}
	err = setup.PreRunRemoteTasks(cfg)
	if err != nil {
		return fmt.Errorf("setup failed pre-running remote tasks: %v", err)
	}
	return nil
}
