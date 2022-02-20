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
package components

import (
	pkgCCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"
)

type GlobalFoldersComponent struct {
	BaseComponent
}

func (c *GlobalFoldersComponent) Execute(ctx pkgCCtx.RuntimeContexter) error {
	if c.next == nil {
		return nil
	}
	return c.next.Execute(ctx)
}

type LocalFoldersComponent struct {
	BaseComponent
}

func (c *LocalFoldersComponent) Execute(ctx pkgCCtx.RuntimeContexter) error {
	if c.next == nil {
		return nil
	}
	return c.next.Execute(ctx)
}
