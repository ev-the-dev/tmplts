package templates

var SwcConfigTemplate = `{
  "$schema": "https://json.schemastore.org/swcrc",
  "module": {
      "type": "commonjs",
      "strict": false,
      "strictMode": true,
      "lazy": false,
      "noInterop": false
  },
  "jsc": {
    "parser": {
      "syntax": "typescript",
      "jsx": false,
      "dynamicimport": false,
      "privatemethod": false,
      "functionbind": false,
      "exportdefaultfrom": false,
      "exportnamespacefrom": false,
      "decorators": false,
      "decoratorsbeforeexport": false,
      "toplevelawait": false,
      "importmeta": true
    },
    "transform": null,
    "target": "es2021",
    "loose": false,
    "externalhelpers": false,
    "keepclassnames": false,
    "isModule": true,
    /* "minify": true, */
    /* "sourceMaps": true */
  }
}
`
