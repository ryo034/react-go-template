import type { Config } from "tailwindcss"
import ta from "tailwindcss-animate"
const baseConfig: Config = require("../../../packages/typescript/config/tailwind/tailwind.config.js")

export default {
  ...baseConfig,
  mode: "jit",
  darkMode: ["class"],
  content: [
    "./index.html",
    "./src/**/*.{ts,tsx}",
    "../../../packages/typescript/ui/src/components/**/*.{js,jsx,ts,tsx}"
  ],
  plugins: [ta]
} satisfies Config
