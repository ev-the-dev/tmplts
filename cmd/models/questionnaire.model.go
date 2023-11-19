package models

type UserAnswers struct {
	AppName    string
	EsBuild    bool
	EsLint     bool
	Jest       bool
	Swc        bool
	Typescript bool
}

/*
*
* JEST
*
 */

func (u *UserAnswers) ListJestDevDependencies() []Dependency {
	deps := []Dependency{}

	if u.Typescript {
		deps = append(deps,
			Dependency{
				Key:   "@types/jest",
				Value: "^29.4.0",
			},
			Dependency{
				Key:   "ts-node", // Jest's TS  runner is tightly coupled to ts-node atm
				Value: "^10.9.1",
			},
		)
	}

	deps = append(deps, Dependency{
		Key:   "jest",
		Value: "^27.2.4",
	})

	return deps
}

func (u *UserAnswers) ListJestScripts() []Script {
	scripts := []Script{
		{
			Key:   "test:unit",
			Value: "npx jest --selectProjects unit --passWithNoTests",
		},
		{
			Key:   "test:integration",
			Value: "npx jest --selectProjects integration",
		},
		{
			Key:   "test:e2e",
			Value: "npx jest --selectProjects e2e",
		},
	}

	return scripts
}

/*
*
* TYPESCRIPT
*
 */

func (u *UserAnswers) ListTypescriptDevDependencies() []Dependency {
	deps := []Dependency{
		{
			Key:   "@types/node",
			Value: "^18.15.11",
		},
		{
			Key:   "typescript",
			Value: "^5.0.4",
		},
	}

	return deps
}

func (u *UserAnswers) ListTypescriptScripts() []Script {
	scripts := []Script{
		{
			Key:   "check",
			Value: "npx tsc --noEmit",
		},
	}

	if !u.Swc && !u.EsBuild {
		scripts = append(scripts, Script{
			Key:   "compile",
			Value: "npx tsc -p ./tsconfig.build.json",
		})
	}
	return scripts
}

/*
*
* ESBUILD
*
 */

func (u *UserAnswers) ListEsBuildDevDependencies() []Dependency {
	deps := []Dependency{
		{
			Key:   "esbuild",
			Value: "^0.19.6",
		},
	}

	return deps
}

func (u *UserAnswers) ListEsBuildScripts() []Script {
	scripts := []Script{
		{
			Key:   "compile",
			Value: "npx esbuild ./index.ts --bundle --sourcemap --minify --platform=node --format=esm --outdir=dist",
		},
	}
	return scripts
}
