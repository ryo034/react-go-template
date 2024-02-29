import { Step } from "gauge-ts"
import { page } from "../../../browser"

export default class OAuthStep {
  @Step("Googleアカウント選択画面でメールアドレス<email>を選択する")
  async selectEmailOnGoogleAccountSelectPage(email: string) {
    await page.locator("#accounts-list").waitFor()
    const target = await page.locator("#reuse-email").getByText(email)
    await target.click()
  }
}
