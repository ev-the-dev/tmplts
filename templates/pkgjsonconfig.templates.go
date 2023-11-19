package templates

var pkgJsonTemplate = `{
  "name": {{.Name}},
  "scripts": {{.Scripts}},
  "dependencies": {{.Dependencies}},
  "devDependencies": {{.DevDependencies}}
}`
