package templates

var PkgJsonTemplate = `{
  "name": {{.Name}},
  "scripts": {{.Scripts}},
  "dependencies": {{.Dependencies}},
  "devDependencies": {{.DevDependencies}}
}`
