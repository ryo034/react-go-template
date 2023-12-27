import terser from "@rollup/plugin-terser"
import react from "@vitejs/plugin-react"
import { visualizer } from "rollup-plugin-visualizer"
import { defineConfig, loadEnv } from "vite"
import viteCompression from "vite-plugin-compression"

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
        output: {
          manualChunks: {
            react: ["react", "react-dom", "react-hook-form", "react-i18next", "react-icons"],
            firebase: ["firebase/app", "react-firebase-hooks/auth"],
            router: ["react-router-dom"],
            i18n: ["i18next"],
            utility: ["clsx", "immer", "true-myth", "class-variance-authority"],
            state: ["zustand"],
            validation: ["zod"],
            tracking: ["@sentry/react", "react-ga4"]
          }
        }
      }
    },
    resolve: {
      alias: {
        "@/": `${__dirname}/src/`,
        "~/": `${__dirname}/src/`
      }
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
