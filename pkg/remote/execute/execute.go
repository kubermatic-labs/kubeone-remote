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

// Package execute provides the functions to execute the individual
// KubeOne Remote commands on the remote node.
package execute

import (
	"fmt"

	"github.com/kubermatic-labs/kubeone-remote/pkg/config"
)

// Executer defines the interface the executing type for the individual
// commands based on the addressed tool (kubeone-emote, kubeone, helm,
// kubectl, ...).
type Executer interface {
	// Do executes the command.
	Do(cfg config.Config) error
}

// Switch determines which Executer implementation later is responsible
// for the command execution.
func Switch(cfg config.Config) (Executer, error) {
	return nil, fmt.Errorf("not yet implemented")
}
