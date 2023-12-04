package templates

var tsConfigTemplate = `{
  "extends": "@tsconfig/node18/tsconfig.json",
  "compilerOptions": {
    "baseUrl": ".",
    "rootDir": ".",
    "outDir": "dist",
    "moduleResolution": "nodenext",
    "module": "nodenext",
    "target": "ES6",
    "paths": {
      "@example/*": ["some/relative/path/*"],
    },
    "noEmit": true,
    "resolveJsonModule": true,
    "allowJs": false,
    "checkJs": false,
    "noImplicitAny": true
  },
  "include": ["./**/package.json", "layers/**/*", "src/**/*", "e2e/**/*"],
  "exclude": ["node_modules", "coverage", "__snapshots__", "docs/**", "dist/**"]
}  
`

var tsConfigBuildTemplate = `{
  "extends": "./tsconfig.json",
  "compilerOptions": {
    "outDir": "dist",
    "noEmit": false,
    "target": "ES6",
    "module": "nodenext",
    "esModuleInterop": true
  },
  "exclude": [
    "node_modules",
    "coverage",
    "__snapshots__",
    "docs/**",
    "dist/**",
    "e2e/**",
    "./**/*.test.ts",
    "./**/*.spec.ts",
    "./**/*.test.ts",
    "./**/*.spec.ts"
  ]
}`

var tsConfigDevTemplate = `{
  "extends": "./tsconfig.json",
  "ts-node": {
    "require": ["tsconfig-paths/register"],
    "swc": true,
    "files": true,
    "compilerOptions": {
      "module": "CommonJS",
      "target": "ES2017"
    }
  }
}`
