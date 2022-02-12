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
	pkgCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"
	pkgIntBuilder "github.com/nitroci/nitroci-core/pkg/internal/builder"
	pkgIntComponents "github.com/nitroci/nitroci-core/pkg/internal/components"
)

func CreateAndInitalizeContext(workspaceType string) (pkgCtx.RuntimeContexter, error) {
	coreBuilder := pkgIntBuilder.GetCoreBuilder(workspaceType)
	director := pkgIntBuilder.NewDirector(coreBuilder)
	ctx := director.BuildContext()
	component := pkgIntComponents.CreateChainOfComponents(ctx)
	component.Execute(ctx)
	return ctx, nil
}
