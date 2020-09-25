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

package local

import (
	"github.com/kubermatic-labs/kubeone-remote/pkg/config"
)

// IsValidCommand checks if the given command is a valid one.
func IsValidCommand(cmd string) bool {
	return false
}

// PreRunLocalTasks performs initial configured tasks.
func PreRunLocalTasks(cfg config.Config) error {
	return nil
}

// InstallKubeOneRemote installs a KubeOne Remote copy on the controller
// node for remote execution.
func InstallKubeOneRemote(cfg config.Config) error {
	return nil
}
