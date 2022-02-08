/*
Copyright 2021 The NitroCI Authors.

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
package contexts

import (
	"errors"
)

type WorkspaceContext struct {
	WorkspacePath       string
	WorkspaceHome       string
	WorkspaceFileFolder string
	WorkspaceFile       string
	Version             int
	Id                  string
	Name                string
}

// Contexts specific functions

func (v *WorkspaceContext) validateWorkspaceContext() error {
	if len(v.Id) == 0 {
		return errors.New("invalid workspace")
	}
	return nil
}
