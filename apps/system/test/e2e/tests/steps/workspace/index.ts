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
}
