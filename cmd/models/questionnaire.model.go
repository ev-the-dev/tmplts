package models

type UserAnswers struct {
	AppName    string
	EsLint     bool
	Jest       bool
	Swc        bool
	Typescript bool
}

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
