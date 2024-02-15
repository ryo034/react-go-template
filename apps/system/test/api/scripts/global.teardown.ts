import { FullConfig } from "@playwright/test"
import { statefulAfterEach } from "./common"

export default async function globalTeardown(config: FullConfig) {
  await statefulAfterEach()
}
