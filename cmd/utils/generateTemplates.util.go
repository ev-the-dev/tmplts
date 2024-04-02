package utils

import (
	"fmt"
	"os"
	"sync"
	"text/template"

	"github.com/ev-the-dev/confits/cmd/models"
	ct "github.com/ev-the-dev/confits/templates"
)

// TODO: If package.json already exists: unmarshal, add these changes, then marshal again. That way we don't overwrite entire files.

// Init Config to be written later
var pkgJsonConfig = models.PackageJsonConfig{
	Name:            "",
	Dependencies:    make(map[string]string),
	DevDependencies: make(map[string]string),
	Scripts:         make(map[string]string),
}

var wg sync.WaitGroup

func GenerateAll(userAnswers *models.UserAnswers, cwd string) {
	// Init Package.json File
	w, err := os.Create(fmt.Sprintf("%s/package.json", cwd))
	if err != nil {
		fmt.Printf("\nUnable to create package.json file: (%v)", err)
	}
	defer wg.Wait()
	defer w.Close()

	pkgJsonConfig.Name = userAnswers.AppName

	// TS CONFIG
	if userAnswers.Typescript {
		fmt.Println("\n### Generating TS Config ###")
		wg.Add(1)
		go GenerateTsConfig(cwd)
		pkgJsonConfig.AddScripts(userAnswers.ListTypescriptScripts())
		pkgJsonConfig.AddDevDependencies(userAnswers.ListTypescriptDevDependencies())
	}

	// JEST CONFIG
	if userAnswers.Jest {
		fmt.Println("\n### Generating Jest Config ###")
		wg.Add(1)
		go GenerateJestConfig(cwd)
		pkgJsonConfig.AddScripts(userAnswers.ListJestScripts())
		pkgJsonConfig.AddDevDependencies(userAnswers.ListJestDevDependencies())
	}

	// ESBUILD CONFIG
	if userAnswers.EsBuild {
		pkgJsonConfig.AddScripts(userAnswers.ListEsBuildScripts())
		pkgJsonConfig.AddDevDependencies(userAnswers.ListEsBuildDevDependencies())
	}

	// SWC CONFIG
	if userAnswers.Swc {
		fmt.Println("\n### Generating SWC Config ###")
		wg.Add(1)
		go GenerateSwcConfig(cwd)
		pkgJsonConfig.AddScripts(userAnswers.ListSWCScripts())
		pkgJsonConfig.AddDevDependencies(userAnswers.ListSWCDevDependencies())
	}

	// ES LINT
	if userAnswers.EsLint {
		fmt.Println("\n### Generating ESLint Config ###")
		wg.Add(1)
		go GeneratEsLintConfig(cwd)
		pkgJsonConfig.AddScripts(userAnswers.ListEsLintScripts())
		pkgJsonConfig.AddDevDependencies(userAnswers.ListEsLintDevDependencies())
	}

	// Populate Package.json
	pkgJsonTmpl := template.Must(template.New("pkgjson").Parse(ct.PkgJsonTemplate))
	pkgJsonMarshalled, err := pkgJsonConfig.MarshallData()
	if err != nil {
		fmt.Printf("Unable to marshal package json file: (%v)", err)
	}
	pkgJsonTmpl.Execute(w, pkgJsonMarshalled)
}

func GeneratEsLintConfig(cwd string) error {
	w, err := os.Create(fmt.Sprintf("%s/eslint.config.js", cwd))
	if err != nil {
		fmt.Printf("\nUnable to create eslint.config.ts file: (%v)", err)
		return err
	}
	defer w.Close()

	eslintTmpl := template.Must(template.New("eslintconfig").Parse(ct.EslintConfigTemplate))

	eslintTmpl.Execute(w, "")

	fmt.Println("\n~~~Done Generating ESLint Config~~~")
	wg.Done()
	return nil
}

func GenerateJestConfig(cwd string) error {
	// TODO: If ESBuild is chosen over SWC, need to change test runner in jestconfig template
	w, err := os.Create(fmt.Sprintf("%s/jest.config.ts", cwd))
	if err != nil {
		fmt.Printf("\nUnable to create jest.config.ts file: (%v)", err)
		return err
	}
	defer w.Close()

	jestTmpl := template.Must(template.New("jestconfig").Parse(ct.JestConfigTemplate))

	jestTmpl.Execute(w, "")

	fmt.Println("\n~~~Done Generating Jest Config~~~")
	wg.Done()
	return nil
}

func GenerateSwcConfig(cwd string) error {
	w, err := os.Create(fmt.Sprintf("%s/.swcrc", cwd))
	if err != nil {
		fmt.Printf("\nUnable to create .swcrc file: (%v)", err)
		return err
	}
	defer w.Close()

	swcTmpl := template.Must(template.New("swcconfig").Parse(ct.SwcConfigTemplate))

	swcTmpl.Execute(w, "")

	fmt.Println("\n~~~Done Generating SWC Config~~~")
	wg.Done()
	return nil
}

func GenerateTsConfig(cwd string) error {
	w, err := os.Create(fmt.Sprintf("%s/tsconfig.json", cwd))
	if err != nil {
		fmt.Printf("\nUnable to create tsconfig.json file: (%v)", err)
		return err
	}
	defer w.Close()

	tsTmpl := template.Must(template.New("tsconfig").Parse(ct.TsConfigTemplate))
	tsBuildTmpl := template.Must(template.New("tsconfigbuild").Parse(ct.TsConfigBuildTemplate))
	tsDevTmpl := template.Must(template.New("tsconfigdev").Parse(ct.TsConfigDevTemplate))

	tsTmpl.Execute(w, "")
	tsBuildTmpl.Execute(w, "")
	tsDevTmpl.Execute(w, "")

	fmt.Println("\n~~~Done Generating TsConfig~~~")
	wg.Done()
	return nil
}
