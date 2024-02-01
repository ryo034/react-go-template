import { expect } from "@playwright/test"
import { Step } from "gauge-ts"
import { page } from "../../../../browser"

export default class EditAccountEmailStep {
  @Step("アカウント設定画面のメールアドレス編集ボタンをクリック")
  async clickEditAccountEmailButton() {
    const target = page.getByTestId("editAccountEmailButton")
    await target.waitFor()
    await target.click()
  }

  @Step("アカウント設定画面のメールアドレスが<email>")
  async checkAccountEmail(email: string) {
    const section = page.getByTestId("editAccountSection")
    await section.waitFor()
    const target = section.getByText(email, { exact: true })
    expect(await target.textContent()).toBe(email)
  }
}
