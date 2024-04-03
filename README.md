# TmplTS
Package to bootstrap TypeScript projects with ESLint, SWC or ESBuild, and Jest configurations.
Additionally it will include appropriate scripts and dependencies in the `package.json` file for each generated configuration file.

Written in Go

## Current Functionality
### Auto-Generate All Files
This will include configuration files automatically for: eslint, jest, swc, tsconfig.
It will also default the app name to the name of the current working directory.
```sh
$ tmplts -a
```

### Selectively Choose Files
This will allow you to choose your own app name and choose which configurations you'd like to include/exclude.
```sh
$ tmplts
```

## Roadmap
* Add Prettier and its config file.
* Add more CLI options, like --version/-v.
* Implement checks for existing files in cwd. The tool should not overwrite files, but append to them if they already exist.
* Simple GitHub deploy workflow adjustment to copy over README and LICENSE into the npm dir prior to publishing.
* More robust configurations: able to selectively add rulesets to each config, instead of having to opt into the entire file as is.
* Automated scripts: go ahead and include any necessary, or quality of life, scripts to get an app up-and-running with minimal effort.
* Add functionality to include popular libraries:
    * Express, Fastify, Vanilla Node
    * Prisma, Drizzle, TypeORM, Knex, Etc
    * Misc libs like Winston, Passport, JWT, etc.
* Styling updates: it's kind of ugly right now.
