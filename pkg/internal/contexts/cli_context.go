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
	"fmt"
	"os"
	"runtime"
	"strings"

	pkgCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"
	pkgFilepath "github.com/nitroci/nitroci-core/pkg/core/extensions/filepath"
	pkgOs "github.com/nitroci/nitroci-core/pkg/core/extensions/os"
	pkgRegistries "github.com/nitroci/nitroci-core/pkg/core/registries"
	pkgIntConfigs "github.com/nitroci/nitroci-core/pkg/internal/configs"
)

const (
	EXTENSION_INI_TYPE = "INI"
	EXTENSION_YML_TYPE = "YML"
	EXTENSION_TYPES    = EXTENSION_INI_TYPE + "," + EXTENSION_YML_TYPE
)

type CliContext struct {
	WorkingDirectory string
	Profile          string
	Environment      string
	WorkspaceDepth   int
	Settings         map[string]string
	Verbose          bool
	Goos             string
	Goarch           string
}

// Creational functions

func (c *CliContext) load() error {
	// Load globacl config configurations
	configEnvVal := pkgOs.GetEnvOrFunc(ENV_NAME_CONFIG, func(s string) string {
		home, _ := os.UserHomeDir()
		return fmt.Sprintf("%v/.nitroci/config.ini", home)
	})
	configPDesc, err := pkgFilepath.GetFilePathDescription(configEnvVal, strings.Split(EXTENSION_TYPES, ","), false)
	if err != nil {
		return err
	}
	c.Settings[pkgCtx.CFG_NAME_CONFIG_PATH] = configPDesc.Path
	c.Settings[pkgCtx.CFG_NAME_CONFIG_HOME] = configPDesc.Home
	c.Settings[pkgCtx.CFG_NAME_CONFIG_FILE] = configPDesc.FileName
	c.Settings[pkgCtx.CFG_NAME_CONFIG_TYPE] = configPDesc.FileExtension
	// Load cache configurations
	chacheEnvVal := pkgOs.GetEnvOrFunc(ENV_NAME_CACHE_HOME, func(s string) string {
		return fmt.Sprintf("%v/cache", c.Settings[pkgCtx.CFG_NAME_CONFIG_HOME])
	})
	cachePDesc, err := pkgFilepath.GetDirPathDescription(chacheEnvVal, false)
	if err != nil {
		return err
	}
	c.Settings[pkgCtx.CFG_NAME_CACHE_PATH] = cachePDesc.Home
	c.Settings[pkgCtx.CFG_NAME_CACHE_PLUGINS_PATH] = fmt.Sprintf("%v/plugins", cachePDesc.Home)
	c.Settings[pkgCtx.CFG_NAME_CACHE_BITS_PATH] = fmt.Sprintf("%v/bits", cachePDesc.Home)
	// Load plugins configurations
	pluginRegistryKey := pkgOs.GetEnvOrDefault(ENV_NAME_PLUGINS_REGISTRY, CFG_DEFVAL_PLUGINS_REGISTRY_GITHUB_URL)
	if !pkgRegistries.IsValidRegistryKey(pluginRegistryKey) {
		return fmt.Errorf("%v is not a valid registry key", pluginRegistryKey)
	}
	c.Settings[pkgCtx.CFG_NAME_PLUGINS_REGISTRY] = pluginRegistryKey
	// Load workspace configurations
	c.Settings[pkgCtx.CFG_NAME_WKS_FILE_FOLDER] = pkgOs.GetEnvOrDefault(ENV_NAME_WKS_FILE_FOLDER, CFG_DEFVAL_WKS_FILE_FOLDER)
	c.Settings[pkgCtx.CFG_NAME_WKS_FILE_NAME] = pkgOs.GetEnvOrDefault(ENV_NAME_WKS_FILE_FOLDER, CFG_DEFVAL_WKS_FILE_NAME)
	// Load bits configurations
	bitsRegistryKey := pkgOs.GetEnvOrDefault(ENV_NAME_BITS_REGISTRY, CFG_DEFVAL_BITS_REGISTRY_GITHUB_URL)
	if !pkgRegistries.IsValidRegistryKey(bitsRegistryKey) {
		return fmt.Errorf("%v is not a valid registry key", bitsRegistryKey)
	}
	c.Settings[pkgCtx.CFG_NAME_BITS_REGISTRY] = bitsRegistryKey
	// Ensure os configuration
	pkgIntConfigs.EnsureConfiguration(configEnvVal)
	err = pkgOs.MkdirInMap(c.Settings, []string{pkgCtx.CFG_NAME_CONFIG_HOME, pkgCtx.CFG_NAME_CACHE_PATH, pkgCtx.CFG_NAME_CACHE_PLUGINS_PATH, pkgCtx.CFG_NAME_CACHE_BITS_PATH})
	if err != nil {
		return err
	}
	return nil
}

func (c *CliContext) validate() error {
	if len(c.WorkingDirectory) == 0 || len(c.Profile) == 0 || c.WorkspaceDepth < 0 {
		return errors.New("invalid cli context")
	}
	return nil
}

func newCliContext(contextInput ContextInput) *CliContext {
	return &CliContext{
		WorkingDirectory: contextInput.workingDirectory,
		Profile:          contextInput.profile,
		Environment:      contextInput.environment,
		WorkspaceDepth:   contextInput.workspaceDepth,
		Settings:         map[string]string{},
		Verbose:          contextInput.verbose,
		Goos:             runtime.GOOS,
		Goarch:           runtime.GOARCH,
	}
}