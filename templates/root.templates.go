package templates

import (
	"fmt"
	"os"
	"text/template"

	"github.com/ehutchllew/template.ts/cmd/models"
)

func GenerateRoot(userAnswers *models.UserAnswers, cwd string) {

	// Init Package.json File
	w, err := os.Create(fmt.Sprintf("%s/package.json", cwd))
	if err != nil {
		fmt.Printf("\nUnable to create package.json file: (%v)", err)
	}
	defer w.Close()

	// Init Config to be written later
	pkgJsonConfig := models.PackageJsonConfig{
		Name:            userAnswers.AppName,
		Dependencies:    make(map[string]string),
		DevDependencies: make(map[string]string),
		Scripts:         make(map[string]string),
	}

	// TS CONFIG
	if userAnswers.Typescript {
		w, err := os.Create(fmt.Sprintf("%s/tsconfig.json", cwd))
		if err != nil {
			fmt.Printf("\nUnable to create tsconfig.json file: (%v)", err)
		}

		tsTmpl := template.Must(template.New("tsconfig").Parse(tsConfigTemplate))
		tsBuildTmpl := template.Must(template.New("tsconfigbuild").Parse(tsConfigBuildTemplate))
		tsDevTmpl := template.Must(template.New("tsconfigdev").Parse(tsConfigDevTemplate))

		tsTmpl.Execute(w, "")
		tsBuildTmpl.Execute(w, "")
		tsDevTmpl.Execute(w, "")
		w.Close()

		pkgJsonConfig.AddScripts(userAnswers.ListTypescriptScripts())
		pkgJsonConfig.AddDevDependencies(userAnswers.ListTypescriptDevDependencies())
	}

	// JEST CONFIG
	if userAnswers.Jest {
		// TODO: If ESBuild is chosen over SWC, need to change test runner in jestconfig template
		w, err := os.Create(fmt.Sprintf("%s/jest.config.ts", cwd))
		if err != nil {
			fmt.Printf("\nUnable to create jest.config.ts file: (%v)", err)
		}

		jestTmpl := template.Must(template.New("jestconfig").Parse(jestConfigTemplate))

		jestTmpl.Execute(w, "")
		w.Close()

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
		w, err := os.Create(fmt.Sprintf("%s/.swcrc", cwd))
		if err != nil {
			fmt.Printf("\nUnable to create .swcrc file: (%v)", err)
		}

		swcTmpl := template.Must(template.New("swcconfig").Parse(swcConfigTemplate))

		swcTmpl.Execute(w, "")
		w.Close()

		pkgJsonConfig.AddScripts(userAnswers.ListSWCScripts())
		pkgJsonConfig.AddDevDependencies(userAnswers.ListSWCDevDependencies())
	}

	// ES LINT
	if userAnswers.EsLint {
		w, err := os.Create(fmt.Sprintf("%s/eslint.config.js", cwd))
		if err != nil {
			fmt.Printf("\nUnable to create eslint.config.ts file: (%v)", err)
		}

		eslintTmpl := template.Must(template.New("eslintconfig").Parse(eslintConfigTemplate))

		eslintTmpl.Execute(w, "")
		w.Close()

		pkgJsonConfig.AddScripts(userAnswers.ListEsLintScripts())
		pkgJsonConfig.AddDevDependencies(userAnswers.ListEsLintDevDependencies())
	}

	// Populate Package.json
	pkgJsonTmpl := template.Must(template.New("pkgjson").Parse(pkgJsonTemplate))

	pkgJsonMarshalled, err := pkgJsonConfig.MarshallData()
	if err != nil {
		fmt.Printf("Unable to marshal package json file: (%v)", err)
	}
	pkgJsonTmpl.Execute(w, pkgJsonMarshalled)
}
