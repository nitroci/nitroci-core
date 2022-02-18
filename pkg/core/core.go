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
package core

import (
	pkgCCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"
	pkgIntCtx "github.com/nitroci/nitroci-core/pkg/internal/contexts"
	pkgIntBuilders "github.com/nitroci/nitroci-core/pkg/internal/builders"
)

func createAndInitalizeContext(workspaceType string, ctxInput pkgCCtx.CoreContextBuilderInput) (pkgCCtx.CoreContexter, error) {
	coreBuilder := pkgIntBuilders.GetCoreBuilder(workspaceType)
	director := pkgIntBuilders.NewDirector(coreBuilder)
	return director.BuildContext(ctxInput), nil
}

func CreateAndInitalizeWorspaceContext(ctxInput pkgCCtx.CoreContextBuilderInput) (pkgCCtx.CoreContexter, error) {
	return createAndInitalizeContext(pkgIntCtx.CORE_BUILDER_WORKSPACE_TYPE, ctxInput)
}

func CreateAndInitalizeWorspacelessContext(ctxInput pkgCCtx.CoreContextBuilderInput) (pkgCCtx.CoreContexter, error) {
	return createAndInitalizeContext(pkgIntCtx.CORE_BUILDER_WORKSPACELESS_TYPE, ctxInput)
}
