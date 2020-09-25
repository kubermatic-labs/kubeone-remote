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

// Package setup provides the functionality to setup KubeOne Remote and
// all needed tools on remote side.
package setup

import (
	"github.com/kubermatic-labs/kubeone-remote/pkg/config"
)

// InstallKubeOneRemote sets up KubeOne Remote on the remote node.
func InstallKubeOneRemote(cfg config.Config) error {
	return nil
}

// CheckoutBranches cares for the ceckout of the possible needed branches
// for Helm and KubeOne Remote.
func CheckoutBranches(cfg config.Config) error {
	return nil
}

// PreRunRemoteTasks performs initial configured tasks.
func PreRunRemoteTasks(cfg config.Config) error {
	return nil
}
