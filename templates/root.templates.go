package templates

import (
	"fmt"
	"os"
	"text/template"

	"github.com/ehutchllew/template.ts/cmd/models"
)

func GenerateRoot(userAnswers *models.UserAnswers, cwd string) {
	// Package.json
	w, err := os.Create(fmt.Sprintf("%s/package.json", cwd))
	if err != nil {
		fmt.Printf("\nUnable to create package.json file: (%v)", err)
	}

	pkgJsonTmpl := template.Must(template.New("pksJson").Parse(pkgJsonTemplate))

	pkgJsonTmpl.Execute(w, userAnswers)
	w.Close()

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
	}

	// JEST CONFIG
	if userAnswers.Jest {
		w, err := os.Create(fmt.Sprintf("%s/jest.config.ts", cwd))
		if err != nil {
			fmt.Printf("\nUnable to create jest.config.ts file: (%v)", err)
		}

		jestTmpl := template.Must(template.New("jestconfig").Parse(jestConfigTemplate))

		jestTmpl.Execute(w, "")
		w.Close()
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
	}

	// ES LINT
	if userAnswers.EsLint {
		w, err := os.Create(fmt.Sprintf("%s/eslint.config.ts", cwd))
		if err != nil {
			fmt.Printf("\nUnable to create eslint.config.ts file: (%v)", err)
		}

		eslintTmpl := template.Must(template.New("eslintconfig").Parse(eslintConfigTemplate))

		eslintTmpl.Execute(w, "")
		w.Close()
	}
}
