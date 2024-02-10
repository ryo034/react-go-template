import { expect } from "@playwright/test"
import { Step } from "gauge-ts"
import { page } from "../browser"
import { homeURL } from "../config"

export default class PagesStep {
  @Step("<second>秒待機")
  async waitSeconds(seconds: number) {
    await page.waitForTimeout(seconds * 1000)
  }

  @Step("ホームに遷移する")
  async openOrdersPage() {
    await page.goto(homeURL)
  }

  @Step("ダッシュボードに遷移")
  async goToDashboardPage() {
    await page.goto(`${homeURL}/dashboard/references`)
  }

  @Step("<path>を開く")
  async openPage(path: string) {
    await page.goto(`${homeURL}${path}`)
  }

  @Step("ログイン画面を開く")
  async goToLoginPage() {
    await page.goto(`${homeURL}`)
  }

  @Step("認証作成画面を開く")
  async goToAuthPage() {
    await page.goto(`${homeURL}/auth`)
  }

  @Step("アカウント設定画面を開く")
  async goToAccountSettings() {
    await page.goto(`${homeURL}/account/settings`)
  }

  @Step("画面を更新する")
  async reloadPage() {
    await page.reload()
  }

  @Step("ダッシュボードに遷移し左側のメニューから<pageName>を選択")
  async moveAndSelectPage(pageName: string) {
    await this.goToDashboardPage()
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
    const target = page.locator("button").getByText(buttonName)
    await target.click()
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

  @Step("現在のURLのパスが<urlPath>")
  async assetUrl(urlPath: string) {
    const currentURL = new URL(await page.url())
    const targetURL = new URL(`${currentURL.origin}${urlPath}`)
    await expect(page).toHaveURL(targetURL.href)
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
    const target = page.getByTestId("userNavigationOnSidebar")
    await target.click()
  }
}
