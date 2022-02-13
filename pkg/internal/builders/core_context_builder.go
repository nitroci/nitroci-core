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
package builders

import (
	pkgCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"
)

type CoreContextBuilder interface {
	createCoreContext()
	initializeCoreContext()
	getCoreContext() pkgCtx.CoreContexter
}

func GetCoreBuilder(builderType string) CoreContextBuilder {
	if builderType == pkgCtx.CORE_BUILDER_WORKSPACE_TYPE {
		return newWorkspaceBuilder()
	}
	if builderType == pkgCtx.CORE_BUILDER_WORKSPACELESS_TYPE {
		return newWorkspacelessBuilder()
	}
	return nil
}