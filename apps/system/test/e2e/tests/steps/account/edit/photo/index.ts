import { log } from "console"
import { expect } from "@playwright/test"
import { Step } from "gauge-ts"
import { page } from "../../../../browser"

export default class EditAccountPhotoStep {
  @Step("サイドバーのアカウント情報にプロフィール画像<imgPath>が設定されている")
  async checkProfilePhotoIsSetWithPath(imgPath: string) {
    const profilePhoto = page.getByTestId("avatarOnSidebar").locator("img")
    const src = await profilePhoto.getAttribute("src")
    expect(src).toContain(imgPath)
  }

  @Step("サイドバーのアカウント情報にプロフィール画像が表示されている")
  async checkProfilePhotoIsSet() {
    const profilePhoto = page.getByTestId("avatarOnSidebar").locator("img")
    const src = await profilePhoto.getAttribute("src")
    expect(src).toContain("http")
  }

  @Step("サイドバーのアカウント情報にプロフィール画像が表示されていない")
  async checkProfilePhotoIsNotSet() {
    const profilePhotoFallback = page.getByTestId("avatarOnSidebar").getByTestId("avatarFallback")
    expect(await profilePhotoFallback.textContent()).toHaveLength(2)
  }

  @Step("プロフィール設定画面のプロフィール画像をクリック")
  async clickProfilePhoto() {
    const profilePhoto = page.getByTestId("avatarOnUpdateProfileForm")
    await profilePhoto.click()
  }

  @Step("画像<imagePath>をアップロードする")
  async uploadImage(imagePath: string) {
    await page.setInputFiles('input[type="file"]', imagePath)
  }

  @Step("プロフィール設定画面のプロフィール削除ボタンをクリック")
  async clickDeleteProfilePhotoButton() {
    const deleteButton = await page.getByTestId("removeProfilePhotoIconButton")
    await deleteButton.click()
  }
}
