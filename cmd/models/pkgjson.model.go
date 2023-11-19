package models

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
