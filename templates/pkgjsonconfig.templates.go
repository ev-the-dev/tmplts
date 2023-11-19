package templates

var pkgJsonTemplate = `{
  "name": {{.Name}},
  "scripts": {
    "check": "npx tsc --noEmit",
    "compile": "npx swc src/ -d dist/",
    "test:unit": "npx jest --selectProjects unitTest --passWithNoTests"
  },
  "dependencies": {{.Dependencies}},
  "devDependencies": {{.DevDependencies}}
}`
