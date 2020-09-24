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

// Package config provides the configuration for the work of
// KubeOne Remote. It also provides the functionality to read the
// configuration from an io.Reader.
package config

import (
	"io"
)

// Flag is the combination of a flag name and an optional value.
type Flag struct {
	Name  string
	Value *string
}

// Config contains all needed parameters.
type Config struct {
	Flags []Flag
}

// Read retrieves data from an io.Reader and unmarshals it into the
// Config structure.
func Read(in io.Reader) (Config, error) {
	return Config{}, nil
}
