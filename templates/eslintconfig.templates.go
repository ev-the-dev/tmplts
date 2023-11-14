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
      /**
       * OG ESLint Rules
       */
      "no-await-in-loop": "error",
      "no-constructor-return": "error",
      "no-duplicate-imports": "error",
      "camelcase": "error",
      "default-case-last": "error",
      "default-param-last": "error",
      "eqeqeq": "error",
      "func-names": ["error", "as-needed"],
      "max-depth": ["error", 2],
      "max-nested-callbacks": ["error", 2],
      "multiline-comment-style": "warn",
      "no-empty-function": "warn",
      "no-lone-blocks": "error",
      "no-loop-func": "warn",
      "no-nested-ternary": "error",
      "no-param-reassign": "error",
      "no-return-assign": "error",
      "no-unneeded-ternary": "error",
      "no-unused-expressions": "warn",
      "no-useless-computed-key": "warn",
      "no-useless-concat": "warn",
      "no-useless-rename": "error",
      "prefer-const": ["error", {
        "destructuring": "any",
        "ignoreReadBeforeAssign": false
      }],
      "sort-imports": "error",
      "sort-keys": ["error", "asc", {
        "natural": true
      }],

      /**
       * TypeScript ESLint Rules
       */
    }
  }
]
`
