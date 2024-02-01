import { expect } from "@playwright/test"
import { Step } from "gauge-ts"
import { page } from "../../../../browser"

export default class EditAccountPhoneNumberStep {
  @Step("アカウント設定画面の電話番号編集ボタンをクリック")
  async clickEditAccountPhoneNumberButton() {
    const target = page.getByTestId("editAccountPhoneNumberButton")
    await target.waitFor()
    await target.click()
  }

  @Step("アカウント設定画面の電話番号が<phoneNumber>")
  async checkAccountPhoneNumber(phoneNumber: string) {
    const section = page.getByTestId("editAccountSection")
    await section.waitFor()
    const target = section.getByText(phoneNumber, { exact: true })
    expect(await target.textContent()).toBe(phoneNumber)
  }
}
