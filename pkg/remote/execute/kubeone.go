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

// kubeoneExecuter implements the Executer for commands
// of KubeOne.
type kubeOneExecuter struct {
	cfg config.Config
}

// NewKubeOneExecuter returns the Executer for KubeOne.
func NewKubeOneExecuter() Executer {
	return &kubeOneExecuter{}
}

// Init implements Executer.
func (e *kubeOneExecuter) Init(cfg config.Config) Executer {
	e.cfg = cfg

	return e
}

// HandlesCommand implements Executer.
func (e *kubeOneExecuter) HandlesCommand() bool {
	return false
}

// Do implements Executer.
func (e *kubeOneExecuter) Do() error {
	return fmt.Errorf("done nothing")
}
