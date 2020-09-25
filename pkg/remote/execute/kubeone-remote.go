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

package execute

import (
	"fmt"

	"github.com/kubermatic-labs/kubeone-remote/pkg/config"
)

// kubeoneRemoteExecuter implements the Executer for commands
// of KubeOne Remote.
type kubeOneRemoteExecuter struct {
	cfg config.Config
}

// NewKubeOneRemoteExecuter returns the Executer for KubeOne Remote.
func NewKubeOneRemoteExecuter() Executer {
	return &kubeOneRemoteExecuter{}
}

// Init implements Executer.
func (e *kubeOneRemoteExecuter) Init(cfg config.Config) Executer {
	e.cfg = cfg

	return e
}

// HandlesCommand implements Executer.
func (e *kubeOneRemoteExecuter) HandlesCommand() bool {
	return false
}

// Do implements Executer.
func (e *kubeOneRemoteExecuter) Do() error {
	return fmt.Errorf("done nothing")
}
