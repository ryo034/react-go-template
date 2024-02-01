import { Step } from "gauge-ts"
import { page } from "../../../browser"

export default class OtpConfirmStep {
  @Step("ワンタイムパスワード確認画面で<message>をクリック")
  async clickMessageOnOtpConfirmPage(message: string) {
    await page.getByText(message).click()
  }
}
