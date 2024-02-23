import { Step } from "gauge-ts"
import { expect } from "playwright/test"
import { page } from "../../browser"
import { homeURL } from "../../config"

export default class InvitationStep {
  @Step("トークン<token>の招待画面を開く")
  async openInvitationPageWithToken(token: string) {
    await page.goto(`${homeURL}/invitation?token=${token}`)
  }

  @Step("招待受諾画面でワークスペース<workspaceName>招待者<inviterName>からの招待が表示されている")
  async checkInvitation(workspaceName: string, inviterName: string) {
    const invitations = page.getByTestId("receivedInvitations").getByTestId("invitation")
    await invitations.waitFor()
    const target = (await invitations.all()).find(async (invitation) => {
      const workspace = await invitation.getByTestId("workspaceName").textContent()
      const inviter = await invitation.getByTestId("inviterName").textContent()
      return workspace === workspaceName && inviter === `招待者: ${inviterName}`
    })
    if (!target) {
      throw new Error(`招待が見つかりませんでした: workspace=${workspaceName}, inviter=${inviterName}`)
    }
    expect(target).toBeVisible()
  }

  @Step("招待受諾画面でワークスペース<workspaceName>招待者<inviterName>の参加ボタンをクリック")
  async clickJoinButton(workspaceName: string, inviterName: string) {
    const invitations = page.getByTestId("receivedInvitations").getByTestId("invitation")
    await invitations.waitFor()
    const target = (await invitations.all()).find(async (invitation) => {
      const workspace = await invitation.getByTestId("workspaceName").textContent()
      const inviter = await invitation.getByTestId("inviterName").textContent()
      return workspace === workspaceName && inviter === `招待者: ${inviterName}`
    })
    if (!target) {
      throw new Error(`招待が見つかりませんでした: workspace=${workspaceName}, inviter=${inviterName}`)
    }
    await target.getByTestId("joinButton").click()
  }

  @Step("招待受諾画面で<email>としてログインしていることがわかる")
  async checkLoginUserEmailOnReceivedInvitationsPage(email: string) {
    const target = page.getByTestId("loggedInBy")
    const text = await target.textContent()
    expect(text).toBe(email)
  }

  @Step("招待画面で<email>としてログインしていることがわかる")
  async checkLoginUserEmailOnReceivedInvitationPage(email: string) {
    const target = page.getByTestId("loggedInBy")
    const text = await target.textContent()
    expect(text).toBe(email)
  }

  @Step("招待中画面にログアウトする")
  async logoutOnReceivedInvitationsPage() {
    const target = page.getByTestId("loggedInBy")
    await target.click()
    const logoutButton = page.getByTestId("logoutButtonOnInvitationHeader")
    await logoutButton.click()
  }
}
