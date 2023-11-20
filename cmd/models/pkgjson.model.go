package models

import (
	"encoding/json"
	"fmt"
)

type Dependency struct {
	Key   string
	Value string
}

type Script struct {
	Key   string
	Value string
}

type PackageJsonConfig struct {
	Name            string
	Dependencies    map[string]string
	DevDependencies map[string]string
	Scripts         map[string]string
}

type PackageJsonMarshalled struct {
	Name            string
	Dependencies    string
	DevDependencies string
	Scripts         string
}

func (p *PackageJsonConfig) AddDependencies(deps []Dependency) error {
	for _, d := range deps {
		p.Dependencies[d.Key] = d.Value
	}

	return nil
}

func (p *PackageJsonConfig) AddDevDependencies(deps []Dependency) error {
	for _, d := range deps {
		p.DevDependencies[d.Key] = d.Value
	}

	return nil
}

func (p *PackageJsonConfig) AddScripts(scripts []Script) error {
	for _, s := range scripts {
		p.Scripts[s.Key] = s.Value
	}

	return nil
}

func (p *PackageJsonConfig) MarshallData() (*PackageJsonMarshalled, error) {
	// FIXME: Think about Error Wrapping -- or something to better expose all possible errors in this chain

	name, nameErr := json.MarshalIndent(p.Name, "\t", "")
	if nameErr != nil {
		fmt.Printf("Unable to marshal package.json name: (%v)", nameErr)
		return nil, nameErr
	}

	deps, depsErr := json.MarshalIndent(p.Dependencies, "\t", "\t")
	if depsErr != nil {
		fmt.Printf("Unable to marshal package.json dependencies: (%v)", depsErr)
		return nil, depsErr
	}

	devDeps, devDepsErr := json.MarshalIndent(p.DevDependencies, "\t", "\t")
	if devDepsErr != nil {
		fmt.Printf("Unable to marshal package.json devDependencies: (%v)", devDepsErr)
		return nil, devDepsErr
	}

	scripts, scriptsErr := json.MarshalIndent(p.Scripts, "\t", "\t")
	if scriptsErr != nil {
		fmt.Printf("Unable to marshal package.json scripts : (%v)", scriptsErr)
		return nil, scriptsErr
	}

	pkgJsonMarshalled := PackageJsonMarshalled{
		Name:            string(name),
		Dependencies:    string(deps),
		DevDependencies: string(devDeps),
		Scripts:         string(scripts),
	}

	return &pkgJsonMarshalled, nil
}
