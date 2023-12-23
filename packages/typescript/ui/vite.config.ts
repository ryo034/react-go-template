import path from "path"
import terser from "@rollup/plugin-terser"
import react from "@vitejs/plugin-react"
import { visualizer } from "rollup-plugin-visualizer"
import { defineConfig, loadEnv } from "vite"
import viteCompression from "vite-plugin-compression"

const packageName = "shared-ui"

type Mode = "development" | "production" | "analyze"

export default ({ mode }: { mode: Mode }) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) }

  return defineConfig({
    build: {
      minify: mode === "production" ? true : false,
      rollupOptions: {
        plugins: [
          ...(mode === "analyze"
            ? [
                visualizer({
                  open: true,
                  filename: "dist/stats.html",
                  gzipSize: true,
                  brotliSize: true
                })
              ]
            : [])
        ],
        external: ["react", "react-dom", "react-hook-form"],
        output: {
          globals: {
            react: "React",
            "react-dom": "ReactDOM",
            "react-hook-form": "ReactHookForm"
          }
        }
      },
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
    css: {
      transformer: "lightningcss"
    },
    plugins: [
      viteCompression(),
      react(),
      terser({
        compress: {
          drop_console: true
        }
      })
    ]
  })
}
