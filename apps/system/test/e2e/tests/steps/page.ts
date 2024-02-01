import { Step } from "gauge-ts"
import { page } from "../browser"
import { homeURL } from "../config"

export default class PagesStep {
  @Step("ホームに遷移する")
  async openOrdersPage() {
    await page.goto(homeURL)
  }
}
