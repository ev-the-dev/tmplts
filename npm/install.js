const fs = require("fs")
const os = require("os")
const path = require("path")

const supportedUnixLikePackages = {
  "darwin arm64 LE": "@tmplts/darwin-arm64",
  "darwin x64 LE": "@tmplts/darwin-x64",
  "linux x64 LE": "@tmplts/linux-x64"
}

function getPkgAndPathForPlatform() {
  const platformKey = `${process.platform} ${os.arch()} ${os.endianness()}`
  const pkgPath = supportedUnixLikePackages[platformKey]

  if(!pkgPath) {
    throw new Error(`Unsupported OS: ${platformKey}`)
  }

  return pkgPath
}

function main() {
  const pkgPath = getPkgAndPathForPlatform()
  const binSubPath = "bin/esbuild"

  const binFullPath = require.resolve(`${pkgPath}/${binSubPath}`)
  const tempPath = path.join(__dirname, "bin-esbuild")
  try {
    /*
     * First: linking binary with temp file.
     */
    fs.linkSync(binFullPath, tempPath)

    /*
     * Second: rename to atomically replace target file with temp file
     */
    fs.renameSync(tempPath, path.join(__dirname, "bin", "tmplts"))

    /*
     * Third: Remove temp file
     */
    fs.unlinkSync(tempPath)
  } catch (e) {
    console.error("Failed to swap out default binary with platform compatible one:\nERROR:", e)
  }
}

