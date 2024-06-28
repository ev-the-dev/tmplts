package templates

var TsConfigTemplate = `{
  "extends": "@tsconfig/node18/tsconfig.json",
  "compilerOptions": {
    "allowImportingTsExtensions": true,
    "allowJs": false,
    "allowUnreachableCode": false,
    "alwaysStrict": true,
    "baseUrl": ".",
    "checkJs": false,
    "esModuleInterop": true,
    "forceConsistentCasingInFileNames": true,
    "lib": ["esnext"],
    "module": "nodenext",
    "moduleResolution": "nodenext",
    "noEmit": true,
    "noErrorTruncation": true,
    "noImplicitAny": true,
    "noImplicitOverride": true,
    "noImplicitReturns": true,
    "noImplicitThis": true,
    "noPropertyAccessFromIndexSignature": true,
    "noUncheckedIndexedAccess": true,
    "noUnusedLocals": true,
    "outDir": "dist",
    "removeComments": true,
    "resolveJsonModule": true,
    "rootDir": ".",
    "strict": true,
    "target": "esnext",
    "paths": {
      "@example/*": ["some/relative/path/*"],
    }
  },
  "include": ["./**/package.json", "layers/**/*", "src/**/*", "e2e/**/*"],
  "exclude": ["node_modules", "coverage", "__snapshots__", "docs/**", "dist/**"]
}  
`

var TsConfigBuildTemplate = `{
  "extends": "./tsconfig.json",
  "compilerOptions": {
    "noEmit": false
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

var TsConfigDevTemplate = `{
  "extends": "./tsconfig.json",
  "ts-node": {
    "require": ["tsconfig-paths/register"],
    "swc": true,
    "files": true
  }
}`
