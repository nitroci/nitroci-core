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
	pkgCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"
)

type iContext interface {
	load() error
	validate() error
}

type context struct {
	iContext
}

func (c *context) createContext() error {
	if err := c.load(); err != nil {
		return err
	}
	if err := c.validate(); err != nil {
		return err
	}
	return nil
}

type ContextInput struct {
	workingDirectory string
	profile string
	environment string
	workspaceDepth int
	verbose bool
}

func CreateContext(contextInput ContextInput, enableWorkspace bool) (pkgCtx.RuntimeContexter, error) {
	var runtimeCtx pkgCtx.RuntimeContexter
	if enableWorkspace {
		runtimeCtx = newRuntimeWorkspaceContext(contextInput)
	} else {
		runtimeCtx = newRuntimeContext(contextInput)
	}
	ctx := context {
		iContext: runtimeCtx.(iContext),
	}
	ctx.createContext()
	return runtimeCtx, nil
}
