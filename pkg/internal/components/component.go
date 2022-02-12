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
	pkgCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"
)

type component interface {
	Execute(pkgCtx.RuntimeContexter)
	setNext(component) component
}

type baseComponent struct {
	next component
}

func (c *baseComponent) setNext(next component) component {
	c.next = next
	return next
}

func CreateChainOfComponents(ctx pkgCtx.RuntimeContexter) component {
	var next component
	// Initialize folders
	next = &globalFoldersComponent{}
	if ctx.IsWorkspaceRequired() {
		next = next.setNext(&localFoldersComponent{})
	}
	// Initialize cache
	next = next.setNext(&globalCacheComponent{})
	if ctx.IsWorkspaceRequired() {
		next = next.setNext(&localCacheComponent{})
	}
	// Initialize plugins
	next = next.setNext(&globalPluginsComponent{})
	if ctx.IsWorkspaceRequired() {
		next = next.setNext(&localPluginsComponent{})
	}
	// Initialize bits
	next = next.setNext(&globalBitsComponent{})
	if ctx.IsWorkspaceRequired() {
		next = next.setNext(&localBitsComponent{})
	}
	return 	next
}
