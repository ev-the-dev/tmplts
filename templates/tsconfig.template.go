package templates

var tsConfigTemplate = `
 {
  "extends": "@tsconfig/node18/tsconfig.json",
  "compilerOptions": {
    "baseUrl": ".",
    "rootDir": ".",
    "outDir": "dist",
    "moduleResolution": "node",
    "module": "ESNext",
    "target": "ES6",
    "paths": {
      "/opt/nodejs/*": ["layers/base-layer/nodejs/*"],
      "*": ["layers/base-layer/nodejs/node_modules/*"]
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
