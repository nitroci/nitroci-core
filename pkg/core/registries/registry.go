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
package registries

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	pkgHttp "github.com/nitroci/nitroci-core/pkg/core/extensions/http"
	pkgOs "github.com/nitroci/nitroci-core/pkg/core/extensions/os"
	pkgTar "github.com/nitroci/nitroci-core/pkg/core/extensions/tar"
	pkgTemplate "github.com/nitroci/nitroci-core/pkg/core/extensions/template"
	pkgYaml "github.com/nitroci/nitroci-core/pkg/core/extensions/yaml"
)

const (
	REGISTRY_GITHUB_TYPE = "GITHUB"
	REGISTRY_GIT_TYPE    = "GIT"
	REGISTRY_PATH_TYPE   = "PATH"
	REGISTRY_TYPES       = REGISTRY_GITHUB_TYPE + "," + REGISTRY_GIT_TYPE + "," + REGISTRY_PATH_TYPE
)

type TemplateData struct {
	Version string
	Goos    string
	Goarch  string
}

type dependency struct {
	name    string
	version string
}

type Registry struct {
	Key          string
	Uri          string
	Type         string
	dependencies map[string]*dependency
}

type RegistryMap struct {
	Goos           string
	Garch          string
	GlobalCache    string
	WorkspaceCache string
	regisitres     map[string]*Registry
}

func IsValidRegistryKey(registryKey string) bool {
	registryDef := strings.Split(registryKey, "+")
	if len(registryDef) == 0 || len(registryDef) > 2 || len(registryDef[0]) == 0 || len(registryDef[1]) == 0 {
		return false
	}
	if registryDef[0] != REGISTRY_GITHUB_TYPE && registryDef[0] != REGISTRY_GIT_TYPE && registryDef[0] != REGISTRY_PATH_TYPE {
		return false
	}
	return true
}

func CreateRegistryMap(globalCache string, workspaceCache string, goos string, goarch string) *RegistryMap {
	registryMap := &RegistryMap{}
	registryMap.Goos = goos
	registryMap.Garch = goarch
	registryMap.GlobalCache = globalCache
	registryMap.WorkspaceCache = workspaceCache
	registryMap.regisitres = map[string]*Registry{}
	return registryMap
}

func (r *RegistryMap) addRegistry(registryKey string) error {
	if !IsValidRegistryKey(registryKey) {
		return fmt.Errorf("%v is not a valid registry key", registryKey)
	}
	registryDef := strings.Split(registryKey, "+")
	if _, ok := r.regisitres[registryKey]; ok {
		return nil
	}
	r.regisitres[registryKey] = &Registry{
		Key:          registryKey,
		Uri:          registryDef[1],
		Type:         REGISTRY_GITHUB_TYPE,
		dependencies: map[string]*dependency{},
	}
	return nil
}

func GetPackageName(name string, version string) string {
	return fmt.Sprintf("%v@%v", name, version)
}

func (r *Registry) AddDependency(name string, version string) error {
	if len(name) == 0 || len(version) == 0 {
		return errors.New("invalid dependency")
	}
	key := GetPackageName(name, version)
	if _, ok := r.dependencies[key]; !ok {
		r.dependencies[key] = &dependency{}
	}
	r.dependencies[key].name = name
	r.dependencies[key].version = version
	return nil
}

func (r *RegistryMap) AddDependency(registryKey string, name string, version string) error {
	err := r.addRegistry(registryKey)
	if err != nil {
		return err
	}
	if _, ok := r.regisitres[registryKey]; !ok {
		return errors.New("unknown registry key")
	}
	return r.regisitres[registryKey].AddDependency(name, version)
}

func (r *RegistryModel) findPackage(dep *dependency) (*PackageDef, error) {
	for _, pkg := range r.Packages {
		if pkg.Name == dep.name {
			return &pkg, nil
		}
	}
	return nil, fmt.Errorf("dependency %v cannot be found", dep.name)
}

func (r *Registry) downloadGitHubRegistryModel() (*RegistryModel, error) {
	pluginsUrl := fmt.Sprintf("%v/nitro-plugins.yml", r.Uri)
	httpResult, err := pkgHttp.HttpGet(pluginsUrl)
	if err != nil {
		return nil, err
	}
	registryModel := &RegistryModel{}
	_, err = pkgYaml.LoadYamlBytes(httpResult.Body, &registryModel)
	if err != nil {
		return nil, err
	}
	return registryModel, nil
}

func (r *Registry) downloadUrlDependency(dep *dependency, packageDef *PackageDef, registryMap *RegistryMap, registryModel *RegistryModel) error {
	folderName := GetPackageName(dep.name, dep.version)
	workspaceFolderName := path.Join(registryMap.WorkspaceCache, folderName)
	globalFolderName := path.Join(registryMap.GlobalCache, folderName)
	if _, err := os.Stat(workspaceFolderName); err == nil {
		return nil
	} else if _, err := os.Stat(globalFolderName); err == nil {
		err := pkgOs.CopyDir(globalFolderName, workspaceFolderName)
		if err != nil {
			return err
		}
		return nil
	}
	replacements := registryModel.Settings.Os.Replacements
	goos := strings.ToLower(registryMap.Goos)
	goarc := strings.ToLower(registryMap.Garch)
	if len(replacements[goos]) > 0 {
		goos = replacements[goos]
	}
	if len(replacements[goarc]) > 0 {
		goarc = replacements[goarc]
	}
	templateData := TemplateData{
		Version: dep.version,
		Goos:    goos,
		Goarch:  goarc,
	}
	releaseUrl, err := pkgTemplate.ExecuteString(packageDef.Release, templateData)
	if err != nil {
		return err
	}
	httpResult, err := pkgHttp.HttpGet(releaseUrl)
	if err != nil {
		return err
	}
	if httpResult.FileExtension != "tar.gz" {
		return fmt.Errorf("%v is an invalid extension format", httpResult.FileExtension)
	}
	fileName := fmt.Sprintf("%v.%v", folderName, httpResult.FileExtension)
	globalPath := path.Join(registryMap.GlobalCache, fileName)
	err = os.WriteFile(globalPath, httpResult.Body, 0644)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(httpResult.Body)
	pkgOs.MkdirInArray([]string{globalFolderName})
	err = pkgTar.Untar(globalFolderName, reader)
	if err != nil {
		return err
	}
	err = pkgOs.CopyDir(globalFolderName, workspaceFolderName)
	if err != nil {
		return err
	}
	return nil
}

func (rm *RegistryMap) Download(onDownload func(string), onFailedDownload func(string)) error {
	for _, r := range rm.regisitres {
		var err error
		var registryModel *RegistryModel
		switch r.Type {
		case REGISTRY_GITHUB_TYPE:
			registryModel, err = r.downloadGitHubRegistryModel()
			if err != nil {
				onFailedDownload(fmt.Sprintf("registry %v cannot be found", r.Key))
				return err
			}
		default:
			onFailedDownload(fmt.Sprintf("registry %v not implemented yet", r.Type))
			return fmt.Errorf("registry %v not implemented yet", r.Type)
		}
		for k, d := range r.dependencies {
			pkg, err := registryModel.findPackage(d)
			if err != nil {
				onFailedDownload(fmt.Sprintf("package %v - %v cannot be found", k, r.Key))
				return err
			}
			err = r.downloadUrlDependency(d, pkg, rm, registryModel)
			if err != nil {
				onFailedDownload(fmt.Sprintf("package %v - %v cannot be downloaded", k, r.Key))
				return err
			}
			onDownload(fmt.Sprintf("downloaded package %v - %v", k, r.Key))
		}
	}
	return nil
}
