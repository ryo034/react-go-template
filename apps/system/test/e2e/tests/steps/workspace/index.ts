import { Step } from "gauge-ts"
import { expect } from "playwright/test"
import { page } from "../../browser"

export default class WorkspaceInfoStep {
  @Step("選択中のワークスペース名が<workspaceName>である")
  async checkWorkspaceNameOnSidebar(workspaceName: string) {
    const target = page.getByTestId("workspaceSwitcher").getByText(workspaceName)
    await target.waitFor()
    expect(target).toBeVisible()
  }

  @Step("所属中の表示名が<displayName>である")
  async checkDisplayName(displayName: string) {
    const target = page.getByTestId("displayNameOnSidebar")
    await target.waitFor()
    const text = await target.textContent()
    expect(text).toBe(displayName)
  }

  @Step("メンバー設定のメンバー<email>の権限が<role>である")
  async checkMemberRole(email: string, role: string) {
    const target = page.getByTestId(`settingMember-${email}`).locator("button").getByText(role)
    await target.waitFor()
    expect(target).toBeVisible()
  }

  @Step("メンバー設定のメンバー<email>の権限変更ボタンをクリック")
  async clickRoleChangeButton(email: string) {
    const target = page.getByTestId(`settingMember-${email}`).locator("button")
    await target.waitFor()
    await target.click()
  }

  @Step("権限選択ポップアップの<role>をクリック")
  async clickRoleSelectButton(role: string) {
    await page.getByTestId("selectMemberRole").getByText(role).click()
  }
}
