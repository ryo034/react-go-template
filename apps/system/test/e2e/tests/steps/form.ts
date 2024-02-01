import { expect } from "@playwright/test"
import { Step } from "gauge-ts"
import { page } from "../browser"

export default class FormStep {
  @Step("入力欄<label>に<value>と入力する")
  async enterFieldValue(label: string, value: string) {
    const target = page.getByLabel(label, { exact: true })
    await target.waitFor()
    await target.fill(value)
  }

  @Step("入力欄<label>に<value>と入力されている")
  async checkFieldValue(label: string, value: string) {
    const target = page.getByLabel(label, { exact: true })
    await target.waitFor()
    await expect(target).toHaveValue(value)
  }

  @Step("入力欄<label>のエラーメッセージ<message>が表示されている")
  async isVisibleEmailErrorMessage(label: string, message: string) {
    const target = page.getByLabel(label, { exact: true })
    await target.waitFor()
    const id = await target.getAttribute("id")
    const errorMessageTarget = page.getByTestId(`${id}-errorMessage`)
    await errorMessageTarget.waitFor()
    expect(await errorMessageTarget.isVisible()).toBeTruthy()
    expect(await errorMessageTarget.innerText()).toEqual(message)
  }

  @Step("入力欄<label>のエラーメッセージが表示されていない")
  async isHiddenFieldErrorMessage(label: string) {
    const target = page.getByLabel(label, { exact: true })
    await target.waitFor()
    const id = await target.getAttribute("id")
    const errorMessageTarget = await page.getByTestId(`${id}-errorMessage`)
    expect(await errorMessageTarget.isHidden()).toBeTruthy()
  }
}
