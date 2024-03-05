import { expect } from "@playwright/test"
import { Step } from "gauge-ts"
import { page } from "../../../../browser"

export default class EditAccountPhotoStep {
  @Step("サイドバーのアカウント情報にプロフィール画像<imgPath>が設定されている")
  async checkProfilePhotoIsSet(imgPath: string) {
    const profilePhoto = await page.getByTestId("avatarOnSidebar")
    const src = await profilePhoto.getAttribute("src")
    expect(src).toContain(imgPath)
  }
}
