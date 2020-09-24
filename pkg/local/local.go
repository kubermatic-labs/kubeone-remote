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

// Package local provides the functions needed for the setup of
// KubeOne-Remote from the local node on the remote controller
// node. In the second step it executes the command remotely via
// SSH.
package local

import (
	"fmt"

	"github.com/kubermatic-labs/kubeone-remote/pkg/config"
)

// Setup manages all needed steps on remote side to make KubeOne-Remote
// working there.
func Setup() (config.Config, error) {
	// TODO
	// - Parse flags
	// - Check command
	// - Read config
	// - Exec pre-run tasks
	// - Copy KubeOne Remote
	return config.Config{}, nil
}

// Execute performes the remote kubeone-remote command with all needed
// environment variables and parameters in a tmux via SSH.
func Execute(cfg config.Config) error {
	cmd := PrepareSSHCommand(cfg)

	session, err := PrepareSSHSession(cfg)
	if err != nil {
		return fmt.Errorf("establishing SSH session failed: %v", err)
	}
	defer session.Close()

	err = session.Run(cmd)
	if err != nil {
		return fmt.Errorf("running SSH command failed: %v", err)
	}

	return nil
}
