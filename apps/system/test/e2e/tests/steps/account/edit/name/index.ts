import { expect } from "@playwright/test"
import { Step } from "gauge-ts"
import { page } from "../../../../browser"

export default class EditAccountNameStep {
  @Step("アカウント設定画面の名前編集ボタンをクリック")
  async clickEditAccountNameButton() {
    const target = page.getByTestId("editAccountNameButton")
    await target.waitFor()
    await target.click()
  }

  @Step("アカウント設定画面の名前が<name>")
  async checkAccountName(name: string) {
    const section = page.getByTestId("editAccountSection")
    await section.waitFor()
    const target = section.getByText(name, { exact: true })
    expect(await target.textContent()).toBe(name)
  }
}
