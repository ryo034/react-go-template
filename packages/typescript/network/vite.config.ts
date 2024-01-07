import path from "path"
import { defineConfig, loadEnv } from "vite"
import viteCompression from "vite-plugin-compression"

const packageName = "shared-network"

type Mode = "development" | "production"

export default ({ mode }: { mode: Mode }) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) }

  return defineConfig({
    build: {
      minify: mode === "production" ? true : false,
      emptyOutDir: false,
      sourcemap: mode !== "production" ? true : false,
      lib: {
        entry: path.resolve(__dirname, "src/index.ts"),
        name: packageName,
        fileName: (format) => `${packageName}.${format}.js`
      }
    },
    resolve: {
      alias: {
        "@/": `${__dirname}/src/`,
        "~/": `${__dirname}/src/`
      }
    },
    plugins: [viteCompression()]
  })
}
