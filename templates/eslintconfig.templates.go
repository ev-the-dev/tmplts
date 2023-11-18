package templates

var eslintConfigTemplate = `
import globals from "globals"
import js from "@eslint/js"
import stylistic from "@stylistic/eslint-plugin"
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
        project: "./tsconfig.json",
      }
    },
    plugins: {
      "@stylistic": stylistic,
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
      "default-param-last": "off", // must be "off" to not conflict with TS "default-param-last"
      "eqeqeq": "error",
      "func-names": ["error", "as-needed"],
      "max-depth": ["error", 2],
      "max-nested-callbacks": ["error", 2],
      "max-params": "off", // TSLint below
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
      "@typescript-eslint/adjacent-overload-signatures": "error",
      "@typescript-eslint/await-thenable": "error",
      "@typescript-eslint/consistent-type-exports": ["warn", {
        "fixMixedExportsWithInlineTypeSpecifier": false
      }],
      "@typescript-eslint/consistent-type-imports": ["warn", {
        "prefer": "type-imports",
        "disallowTypeAnnotations": true,
        "fixStyle": "separate-type-imports"
      }],
      "@typescript-eslint/default-param-last": "error",
      "@typescript-eslint/explicit-function-return-type": "warn",
      "@typescript-eslint/explicit-member-accessibility": "error",
      "@typescript-eslint/max-params": ["error", { "max": 4 }],
      "@typescript-eslint/member-ordering": ["error", {
        default: {
          "memberTypes": [
            // Index Signature
            "call-signature",

            // Fields
            "public-abstract-field",
            "protected-abstract-field",

            "public-static-field",
            "protected-static-field",
            "private-static-field",
            "#private-static-field",

            "public-decorated-field",
            "protected-decorated-field",
            "private-decorated-field",
            // "#private-decorated-field",

            "public-instance-field",
            "protected-instance-field",
            "private-instance-field",
            "#private-instance-field",

            "abstract-field",
            "static-field",
            "decorated-field",
            "instance-field",

            "public-field",
            "protected-field",
            "private-field",
            "#private-field",

            "field",

            // Constructors
            "public-constructor",
            "protected-constructor",
            "private-constructor",

            "constructor",

            // Methods
            "public-abstract-method",
            "protected-abstract-method",

            "public-static-method",
            "protected-static-method",
            "private-static-method",
            "#private-static-method",

            "public-decorated-method",
            "protected-decorated-method",
            "private-decorated-method",
            // "#private-decorated-method",

            "public-instance-method",
            "protected-instance-method",
            "private-instance-method",
            "#private-instance-method",

            "abstract-method",
            "static-method",
            "decorated-method",
            "instance-method",

            "public-method",
            "protected-method",
            "private-method",
            "#private-method",

            "method",
          ],
          order: "natural-case-insensitive"
        }
      }],

      /**
       * Stylistic ESLint Rules
       */
      "@stylistic/indent": "error",
    },
  },

  {
    files: ["*.js"],
    rules: {
      "@typescript-eslint/explicit-member-accessibility": "off"
    }
  }
]
`
