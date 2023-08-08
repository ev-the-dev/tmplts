package templates

var pkgJsonTemplate = `{
  "name": {{.AppName}},
  "scripts": {
    "check": "npx tsc --noEmit",
    "compile": "npx swc src/ -d dist/",
    "test:unit": "npx jest --selectProjects unitTest --passWithNoTests"
  },
  "dependencies": {},
  "devDependencies": {}
}`
