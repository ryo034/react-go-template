import { expect, test } from "@playwright/test"
import { components } from "schema/openapi/systemApi"
import { authHeaders, defaultPostHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo, getInvitationIdByToken, getInviteToken, statefulTest } from "../../scripts"
const client = genAPIClient()

test.describe("Me Invitations", () => {
  // not used user accept invitation
  statefulTest(
    "create account and add invited workspace to joinedWorkspaces and currentWorkspace set to invited workspace @stateful",
    async ({ page }) => {
      const email = "invite_test_not_expired@example.com"
      const inviteToken = await getInviteToken(email)
      const processRes = await client.POST("/api/v1/auth/invitations/process", {
        headers: defaultPostHeaders,
        body: { token: inviteToken, email }
      })
      expect(processRes.response.status).toBe(200)

      const authInfo = await getAuthInfo(email)

      const res = await client.GET("/api/v1/me", { headers: authHeaders(authInfo.token) })
      expect(res.response.status).toBe(200)
      expect(res.data?.self.email).toBe(email)
      expect(res.data?.self.userId).not.toBeNull()
      expect(res.data?.member).toBeUndefined()
      expect(res.data?.currentWorkspace).toBeUndefined()
      expect(res.data?.receivedInvitations).toBeUndefined()

      if (res.data === undefined) {
        throw new Error("res.data is undefined")
      }
      // set account name
      const data: components["schemas"]["User"] = {
        userId: res.data?.self.userId,
        email: res.data?.self.email,
        name: "Updated Name",
        phoneNumber: ""
      }
      const updateProfileRes = await client.PUT("/api/v1/me/profile", {
        headers: authHeaders(authInfo.token),
        body: { user: data }
      })
      expect(updateProfileRes.response.status).toBe(200)

      // accept invitation
      const acceptInvitationRes = await client.POST("/api/v1/members/invitations/{invitationId}/accept", {
        headers: authHeaders(authInfo.token),
        params: { path: { invitationId: await getInvitationIdByToken(inviteToken) } }
      })
      expect(acceptInvitationRes.response.status).toBe(200)
      if (acceptInvitationRes.data === undefined) {
        throw new Error("acceptInvitationRes.data is undefined")
      }
      expect(acceptInvitationRes.data.joinedWorkspaces).toEqual([
        {
          workspaceId: "c1bd2603-b9cd-4f84-8b83-3548f6ae150b",
          name: "Example",
          subdomain: "example"
        }
      ])
      expect(acceptInvitationRes.data.currentWorkspace?.workspaceId).toBe("c1bd2603-b9cd-4f84-8b83-3548f6ae150b")
    }
  )

  // already used user accept invitation
  statefulTest(
    "add invited workspace to joinedWorkspaces and currentWorkspace changed to invited workspace @stateful",
    async ({ page }) => {
      const email = "invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com"
      const inviteToken = await getInviteToken(email)
      const processRes = await client.POST("/api/v1/auth/invitations/process", {
        headers: defaultPostHeaders,
        body: { token: inviteToken, email }
      })
      expect(processRes.response.status).toBe(200)

      const authInfo = await getAuthInfo(email)

      const res = await client.GET("/api/v1/me", { headers: authHeaders(authInfo.token) })
      expect(res.response.status).toBe(200)
      expect(res.data).toStrictEqual((await import("./already_used_user_accept_invitation_me_res.json")).default)

      // accept invitation
      const acceptInvitationRes = await client.POST("/api/v1/members/invitations/{invitationId}/accept", {
        headers: authHeaders(authInfo.token),
        params: { path: { invitationId: await getInvitationIdByToken(inviteToken) } }
      })
      expect(acceptInvitationRes.response.status).toBe(200)
      if (acceptInvitationRes.data === undefined) {
        throw new Error("acceptInvitationRes.data is undefined")
      }

      expect(acceptInvitationRes.data.joinedWorkspaces).toEqual([
        {
          name: "InviteTest 2",
          subdomain: "invite-test-2",
          workspaceId: "018d9b4d-e340-74f7-914c-2476eff949bb"
        },
        {
          workspaceId: "c1bd2603-b9cd-4f84-8b83-3548f6ae150b",
          name: "Example",
          subdomain: "example"
        }
      ])
      expect(acceptInvitationRes.data.currentWorkspace?.workspaceId).toBe("c1bd2603-b9cd-4f84-8b83-3548f6ae150b")
    }
  )

  test("failed to accept already accepted invitation return ConflictError", async () => {
    const email = "invite_test_already_used@example.com"
    const inviteToken = await getInviteToken(email)
    const authInfo = await getAuthInfo(email)
    const res = await client.POST("/api/v1/members/invitations/{invitationId}/accept", {
      headers: authHeaders(authInfo.token),
      params: { path: { invitationId: await getInvitationIdByToken(inviteToken) } }
    })
    expect(res.response.status).toBe(410)
  })

  test("failed to accept already expired invitation return BadRequestError", async () => {
    const email = "invite_test_expired@example.com"
    const inviteToken = await getInviteToken(email)
    const authInfo = await getAuthInfo(email)
    const res = await client.POST("/api/v1/members/invitations/{invitationId}/accept", {
      headers: authHeaders(authInfo.token),
      params: { path: { invitationId: await getInvitationIdByToken(inviteToken) } }
    })
    expect(res.response.status).toBe(410)
  })

  test("failed to accept revoked invitation return BadRequestError", async () => {
    const email = "invite_test_revoked@example.com"
    const inviteToken = await getInviteToken(email)
    const authInfo = await getAuthInfo(email)
    const res = await client.POST("/api/v1/members/invitations/{invitationId}/accept", {
      headers: authHeaders(authInfo.token),
      params: { path: { invitationId: await getInvitationIdByToken(inviteToken) } }
    })
    expect(res.response.status).toBe(410)
  })
})
