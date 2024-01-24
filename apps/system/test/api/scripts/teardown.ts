import { FullConfig } from "@playwright/test"
import { statefulBeforeEach } from "./common"

export default async function globalTeardown(config: FullConfig) {
  await statefulBeforeEach()
}
