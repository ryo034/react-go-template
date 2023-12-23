/** @type {import('tailwindcss').Config} */
const baseConfig = require("../../packages/config/tailwind/tailwind.config.js")
module.exports = {
  ...baseConfig,
  content: ["./src/**/*.{ts,tsx}"]
}
