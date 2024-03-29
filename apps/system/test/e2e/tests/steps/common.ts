import { expect } from "@playwright/test"
import { Step } from "gauge-ts"
import { page } from "../browser"
import { homeURL } from "../config"

export default class PagesStep {
  @Step("<second>秒待機")
  async waitSeconds(seconds: number) {
    await page.waitForTimeout(seconds * 1000)
  }

  @Step("<path>を開く")
  async openPage(path: string) {
    await page.goto(`${homeURL}${path}`)
  }

  @Step("画面を更新する")
  async reloadPage() {
    await page.reload()
  }

  @Step("ダッシュボードに遷移し左側のメニューから<pageName>を選択")
  async moveAndSelectPage(pageName: string) {
    await page.goto(`${homeURL}/dashboard/references`)
    const target = page.getByTestId("sidebarMenuList").getByText(pageName)
    await target.waitFor()
    await target.click()
  }

  @Step("サイドバーメニューから<pageName>を選択")
  async selectFromSidebarMenu(pageName: string) {
    const target = page.getByTestId("sidebarMenuList").getByText(pageName)
    await target.click()
  }

  @Step("ログインボタンをクリック")
  async clickLoginButton() {
    const target = page.locator("button").getByText("ログイン")
    await target.click()
  }

  @Step("作成ボタンをクリック")
  async clickCreateButton() {
    const target = page.locator("button").getByText("作成")
    await target.click()
  }

  @Step("<buttonName>ボタンをクリック")
  async clickButtonByName(buttonName: string) {
    await page.locator("button").getByText(buttonName).click()
  }

  @Step("ダイアログ上の<buttonName>ボタンをクリック")
  async clickButtonByNameOnAlertDialog(buttonName: string) {
    await page.getByTestId('dialog').locator("button").getByText(buttonName).click()
  }

  @Step("入力フォームの<buttonName>ボタンをクリック")
  async clickButtonByNameOnForm(buttonName: string) {
    await page.locator("form").locator("button").getByText(buttonName).click()
  }

  @Step("<text>をクリック")
  async clickTextByName(text: string) {
    const target = page.getByText(text)
    await target.waitFor()
    await target.click()
  }

  @Step("モーダルのキャンセルボタンをクリック")
  async clickCancelButtonModal() {
    const target = page.locator(".modal").locator("button").getByText("キャンセル")
    await target.waitFor()
    await target.click()
  }

  @Step("モーダルの作成ボタンをクリック")
  async clickCreateButtonModal() {
    const target = page.locator(".modal").locator("button").getByText("作成")
    await target.waitFor()
    await target.click()
  }

  @Step("モーダルの閉じるアイコンボタンをクリック")
  async clickCancelIconButtonModal() {
    const target = page.getByTestId("closeIconButtonOnModal")
    await target.waitFor()
    await target.click()
  }

  @Step("エラーメッセージ<errorMessage>が表示されている")
  async isVisibleErrorMessage(errorMessage: string) {
    const target = page.getByTestId("resultError").getByText(errorMessage)
    await target.waitFor()
    expect(await target.isVisible()).toBeTruthy()
  }

  private testURL(url: URL, pattern: string): boolean {
    const regexPattern = pattern.replace(/:[a-zA-Z0-9_]+/, "\\w+")
    const regex = new RegExp(regexPattern)
    return regex.test(url.pathname)
  }

  private createPatternFromURL(url: string): RegExp {
    const patternString = url.replace(/:[^/]*/g, "[^/]*")
    const pattern = new RegExp(patternString)
    return pattern
  }

  @Step("現在のURLのパスのパターンが<url>")
  async assetUrlRegex(url: string) {
    const pattern = this.createPatternFromURL(url)
    await page.waitForURL(pattern)
    const currentURL = new URL(await page.url())
    const isMatch = this.testURL(currentURL, url)
    expect(isMatch).toBeTruthy()
  }

  @Step("サイドバーのユーザーアイコンをクリック")
  async clickSidebarUserIcon() {
    const target = page.getByTestId("userNavigationOnSidebar").getByTestId("avatarOnSidebar")
    await target.click()
  }

  @Step("サイドバーの<target>をクリック")
  async clickSidebarTarget(target: string) {
    const targetsOnSidebar = page.getByTestId("pagesOnSidebar")
    const targetElement = targetsOnSidebar.getByText(target)
    await targetElement.click()
  }

  @Step("トーストメッセージ<message>が表示されている")
  async isVisibleToastMessage(message: string) {
    const target = page.getByTestId("toastTitle")
    await target.waitFor()
    expect(await target.textContent()).toBe(message)
  }

  @Step("ページタイトルが<title>である")
  async assetPageTitle(title: string) {
    const pageTitle = await page.getByTestId("pageTitle").textContent()
    expect(pageTitle).toBe(title)
  }

  @Step("ワンタイムパスワード入力画面が表示されている")
  async isVisibleOtpPage() {
    const target = page.getByTestId("verifyOtpPage")
    await target.waitFor()
    expect(target).toBeVisible()
  }

  @Step("オンボーディングの名前入力画面が表示されている")
  async isVisibleOnboardingNamePage() {
    const target = page.getByTestId("onboardingSettingNamePage")
    await target.waitFor()
    expect(target).toBeVisible()
  }

  @Step("オンボーディングのワークスペース作成画面が表示されている")
  async isVisibleOnboardingWorkspacePage() {
    const target = page.getByTestId("onboardingCreateWorkspacePage")
    await target.waitFor()
    expect(target).toBeVisible()
  }

  @Step("招待受諾画面が表示されている")
  async isVisibleReceivedInvitationsPage() {
    const target = page.getByTestId("receivedInvitationsPage")
    await target.waitFor()
    expect(target).toBeVisible()
  }

  @Step("ログイン画面が表示されている")
  async isVisibleLoginPage() {
    const target = page.getByTestId("authPage")
    await target.waitFor()
    expect(target).toBeVisible()
  }

  @Step("ホーム画面が表示されている")
  async isVisibleHomePage() {
    const target = page.getByTestId("homePage")
    await target.waitFor()
    expect(target).toBeVisible()
  }
}
