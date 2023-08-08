package templates

var swcConfigTemplate = `{
  "jsc": {
    "parser": {
      "syntax": "typescript",
      "jsx": false,
      "dynamicImport": false,
      "privateMethod": false,
      "functionBind": false,
      "exportDefaultFrom": false,
      "exportNamespaceFrom": false,
      "decorators": false,
      "decoratorsBeforeExport": false,
      "topLevelAwait": false,
      "importMeta": true
    },
    "transform": null,
    "target": "es2021",
    "loose": false,
    "externalHelpers": false,
    "keepClassNames": false
  }
}
`
