package templates

var eslintConfigTemplate = `
import globals from "globals"
import js from "@eslint/js"
// import tslint from "@typescript-eslint"
import tslintPlugin from "@typescript-eslint/eslint-plugin"
import tslintParser from "@typescript-eslint/parser"

export default [
  js.configs.recommended,
  // tslint.eslint-plugin.configs.recommended, // throws error on new config file type due to internal "extends"
  {
    files: ["**/*.ts"],
    languageOptions: {
      globals: {
        ...globals.node,
        ...globals.jest,
      },
      ecmaVersion: "latest",
      sourceType: "module",
      parser: tslintParser,
      parserOptions: {
        ecmaFeatures: {
          modules: true,
          project: "./tsconfig.json"
        }
      }
    },
    plugins: {
      "@typescript-eslint": tslintPlugin,
    },
    rules: {
      "no-await-in-loop": "error",
      "no-constructor-return": "error",
      "no-duplicate-imports": "error",
      "camelcase": "error",
    }
  }
]
`
