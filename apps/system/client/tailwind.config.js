/** @type {import('tailwindcss').Config} */
const baseConfig = require("../../../packages/typescript/config/tailwind/tailwind.config.js")
module.exports = {
  ...baseConfig,
  mode: "jit",
  darkMode: ["class"],
  content: [
    "./index.html",
    "./src/**/*.{ts,tsx}",
    "../../../packages/typescript/ui/src/components/**/*.{js,jsx,ts,tsx}"
  ],
  plugins: [require("tailwindcss-animate")]
}
