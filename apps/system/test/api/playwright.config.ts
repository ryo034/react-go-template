import { defineConfig, devices } from "@playwright/test"
import { config } from "dotenv"

config()

const globalSetupFilePath = "./scripts/global.setup"
const globalTeardownFilePath = "./scripts/global.teardown"

/**
 * See https://playwright.dev/docs/test-configuration.
 */
export default defineConfig({
  testDir: "tests",
  forbidOnly: !!process.env.CI,
  fullyParallel: !!process.env.PARALLEL,
  retries: process.env.CI ? 2 : 0,
  workers: process.env.CI ? 1 : undefined,
  globalSetup: globalSetupFilePath,
  globalTeardown: globalTeardownFilePath,
  timeout: 15000,
  reporter: process.env.CI ? "html" : "line",
  use: {
    trace: process.env.CI ? "on-first-retry" : "on",
    video: process.env.CI ? "on-first-retry" : "on"
  },
  projects: [
    {
      name: "chromium",
      use: { ...devices["Desktop Chrome"] }
    }
    // {
    //   name: 'firefox',
    //   use: { ...devices['Desktop Firefox'] },
    // },
    // {
    //   name: 'webkit',
    //   use: { ...devices['Desktop Safari'] },
    // },
  ]
})
