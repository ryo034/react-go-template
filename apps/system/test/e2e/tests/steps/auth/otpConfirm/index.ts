import { Step } from "gauge-ts"
import { page } from "../../../browser"
import { dataStore, keys } from "../../../dataStore"
import { getOtpCodeFromRedis } from "../../../redis"

export default class OtpConfirmStep {
  @Step("ワンタイムパスワード確認画面で<message>をクリック")
  async clickMessageOnOtpConfirmPage(message: string) {
    await page.getByText(message).click()
  }

  @Step("メールアドレス<email>に送信されたワンタイムパスワードを取得")
  async getOtpCodeFromRedis(email: string) {
    const code = await getOtpCodeFromRedis(email)
    if (code === null) {
      throw new Error("code is null")
    }
    dataStore.put(keys.otp, code)
  }

  @Step("ワンタイムパスワード確認画面にワンタイムパスワードを入力する")
  async inputOtpCodeOnOtpConfirmPage() {
    const code = dataStore.get<string, string>(keys.otp)
    const [code1, code2, code3, code4, code5, code6] = [...code]
    await page.getByTestId("otpInput1").fill(code1)
    await page.getByTestId("otpInput2").fill(code2)
    await page.getByTestId("otpInput3").fill(code3)
    await page.getByTestId("otpInput4").fill(code4)
    await page.getByTestId("otpInput5").fill(code5)
    await page.getByTestId("otpInput6").fill(code6)
  }
}
