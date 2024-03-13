import { expect } from "@playwright/test"
import type { components } from "schema/openapi/systemApi"
import { authHeaders, defaultPostHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo, getInvitationIdByToken, getInviteToken, systemTest } from "../../scripts"
const client = genAPIClient()

systemTest.describe("Me Invitations", () => {
  // not used user accept invitation
  systemTest(
    "create account and add invited workspace to joinedWorkspaces and currentWorkspace set to invited workspace",
    { tag: ["@stateful"] },
    async ({ stateful }) => {
      const email = "invite_test_not_expired@example.com"
      const inviteToken = await getInviteToken(email)
      const processRes = await client.POST("/api/v1/auth/invitations/process/email", {
        headers: defaultPostHeaders,
        body: { token: inviteToken, email }
      })
      expect(processRes.response.status).toBe(200)
      const authInfo = await getAuthInfo(email)
      const res = await client.GET("/api/v1/me", { headers: authHeaders(authInfo.token) })
      expect(res.response.status).toBe(200)
      expect(res.data?.me.self.email).toBe(email)
      expect(res.data?.me.self.userId).not.toBeNull()
      expect(res.data?.me.member).toBeUndefined()
      expect(res.data?.me.currentWorkspace).toBeUndefined()
      expect(res.data?.me.receivedInvitations).toEqual([
        {
          invitation: {
            displayName: "",
            expiredAt: "2200-01-10T21:00:00+09:00",
            id: "018d96b8-2211-7862-bcbe-e9f4d002a8fc",
            inviteeEmail: "invite_test_not_expired@example.com",
            accepted: false,
            inviter: {
              id: "377eba35-5560-4f48-a99d-19cbd6a82b0d",
              membershipStatus: "ACTIVE",
              profile: {
                bio: "bio",
                displayName: "John Doe",
                idNumber: "DEV-12345"
              },
              role: "OWNER",
              user: {
                email: "account@example.com",
                name: "John Doe",
                userId: "394e67b6-2850-4ddf-a4c9-c2a619d5bf70"
              }
            }
          },
          inviter: {
            member: {
              id: "377eba35-5560-4f48-a99d-19cbd6a82b0d",
              membershipStatus: "ACTIVE",
              profile: {
                displayName: "John Doe",
                idNumber: "DEV-12345",
                bio: "bio"
              },
              role: "OWNER",
              user: {
                email: "account@example.com",
                name: "John Doe",
                userId: "394e67b6-2850-4ddf-a4c9-c2a619d5bf70"
              }
            },
            workspace: { name: "Example", subdomain: "example", workspaceId: "c1bd2603-b9cd-4f84-8b83-3548f6ae150b" }
          }
        }
      ])
      if (res.data === undefined) {
        throw new Error("res.data is undefined")
      }
      // set account name
      const data: components["schemas"]["User"] = {
        userId: res.data.me?.self.userId,
        email: res.data.me?.self.email,
        name: "Updated Name",
        phoneNumber: ""
      }
      const updateProfileRes = await client.PUT("/api/v1/me/profile", {
        headers: authHeaders(authInfo.token),
        body: { profile: data }
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
      expect(acceptInvitationRes.data.me.joinedWorkspaces).toEqual([
        {
          workspaceId: "c1bd2603-b9cd-4f84-8b83-3548f6ae150b",
          name: "Example",
          subdomain: "example"
        }
      ])
      expect(acceptInvitationRes.data.me.currentWorkspace?.workspaceId).toBe("c1bd2603-b9cd-4f84-8b83-3548f6ae150b")
    }
  )
  // already used user accept invitation
  systemTest(
    "add invited workspace to joinedWorkspaces and currentWorkspace changed to invited workspace",
    { tag: ["@stateful"] },
    async ({ stateful }) => {
      const email = "invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com"
      const inviteToken = await getInviteToken(email)
      const processRes = await client.POST("/api/v1/auth/invitations/process/email", {
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
      expect(
        acceptInvitationRes.data.me.joinedWorkspaces.find(
          (w) => w.workspaceId === "c1bd2603-b9cd-4f84-8b83-3548f6ae150b"
        )
      ).toEqual({
        workspaceId: "c1bd2603-b9cd-4f84-8b83-3548f6ae150b",
        name: "Example",
        subdomain: "example"
      })
      expect(
        acceptInvitationRes.data.me.joinedWorkspaces.find(
          (w) => w.workspaceId === "018d9b4d-e340-74f7-914c-2476eff949bb"
        )
      ).toEqual({
        name: "InviteTest 2",
        subdomain: "invite-test-2",
        workspaceId: "018d9b4d-e340-74f7-914c-2476eff949bb"
      })
      expect(acceptInvitationRes.data.me.currentWorkspace?.workspaceId).toBe("c1bd2603-b9cd-4f84-8b83-3548f6ae150b")
      expect(acceptInvitationRes.data.me.receivedInvitations).toBeUndefined()
    }
  )
  systemTest("failed to accept already accepted invitation return GoneError", async () => {
    const email = "invite_test_already_accepted@example.com"
    const inviteToken = await getInviteToken(email)
    const authInfo = await getAuthInfo(email)
    const res = await client.POST("/api/v1/members/invitations/{invitationId}/accept", {
      headers: authHeaders(authInfo.token),
      params: { path: { invitationId: await getInvitationIdByToken(inviteToken) } }
    })
    expect(res.response.status).toBe(410)
  })
  systemTest("failed to accept already expired invitation return GoneError", async () => {
    const email = "invite_test_expired@example.com"
    const inviteToken = await getInviteToken(email)
    const authInfo = await getAuthInfo(email)
    const res = await client.POST("/api/v1/members/invitations/{invitationId}/accept", {
      headers: authHeaders(authInfo.token),
      params: { path: { invitationId: await getInvitationIdByToken(inviteToken) } }
    })
    expect(res.response.status).toBe(410)
  })
  systemTest("failed to accept revoked invitation return GoneError", async () => {
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
