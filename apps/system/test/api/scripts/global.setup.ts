import { FullConfig } from "@playwright/test"
import { statefulBeforeEach } from "./common"

export default async function globalSetup(config: FullConfig) {
  console.log("globalSetup")
  await statefulBeforeEach()
}
