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
	"golang.org/x/crypto/ssh"

	"github.com/kubermatic-labs/kubeone-remote/pkg/config"
)

// PrepareSSHCommand creates the command to be run on the controller
// node based on the config.
func PrepareSSHCommand(cfg config.Config) string {
	return ""
}

// PrepareSSHSession establishes the SSH session to run the command.
// It has to be closed by the caller.
func PrepareSSHSession(cfg config.Config) (*ssh.Session, error) {
	return &ssh.Session{}, nil
}
