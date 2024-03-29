import test, { expect } from "@playwright/test"
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
            id: "018e4922-563a-7566-bc5e-65dc2f8faefe",
            inviteeEmail: "invite_test_not_expired@example.com",
            accepted: false,
            inviter: {
              id: "018e4922-563a-7807-b01f-2e630e4d22e9",
              membershipStatus: "ACTIVE",
              profile: {
                bio: "bio",
                displayName: "Invite TestHasEvent",
                idNumber: "DEV-12345"
              },
              role: "OWNER",
              user: {
                email: "invite_test_has_event_inviter@example.com",
                name: "Invite TestHasEvent",
                userId: "018e4922-563a-7097-bbdb-ffa9f74da283",
              }
            }
          },
          inviter: {
            member: {
              id: "018e4922-563a-7807-b01f-2e630e4d22e9",
              membershipStatus: "ACTIVE",
              profile: {
                displayName: "Invite TestHasEvent",
                idNumber: "DEV-12345",
                bio: "bio"
              },
              role: "OWNER",
              user: {
                email: "invite_test_has_event_inviter@example.com",
                name: "Invite TestHasEvent",
                userId: "018e4922-563a-7097-bbdb-ffa9f74da283",
              }
            },
            workspace: {
              name: "Invite TestHasEvent",
              subdomain: "invite-test-has-event",
              workspaceId: "018e4922-563a-7731-b389-c2a9ac0d97e9",
            }
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
          name: "Invite TestHasEvent",
          subdomain: "invite-test-has-event",
          workspaceId: "018e4922-563a-7731-b389-c2a9ac0d97e9",
        }
      ])
      expect(acceptInvitationRes.data.me.currentWorkspace?.workspaceId).toBe("018e4922-563a-7731-b389-c2a9ac0d97e9")
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
          (w) => w.workspaceId === "018d96b9-c920-7434-b5c3-02e5e920ae9d"
        )
      ).toEqual({
        workspaceId: "018d96b9-c920-7434-b5c3-02e5e920ae9d",
        name: "InviteTest 1",
        subdomain: "invite-test-1"
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
      expect(acceptInvitationRes.data.me.currentWorkspace?.workspaceId).toBe("018d96b9-c920-7434-b5c3-02e5e920ae9d")
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

test.describe("receive invitation from already left workspace member", () => {
  test("check invitation if received masked inviter details invitation from already left workspace", async ({ page }) => {
    const email = "once_leave_workspace_check_receive_from_already_left_member@example.com"
    const inviteToken = await getInviteToken(email)
    const invitationRes = await client.GET("/api/v1/auth/invitations", {
      headers: defaultPostHeaders,
      params: { query: { token: inviteToken } }
    })
    expect(invitationRes.response.status).toBe(200)
    expect(invitationRes.data).toStrictEqual((await import("./success_get_masked_inviter_invitation_received_from_left_member.json")).default)
  })

  test("if received invitation from already left workspace, user accept invitation and join workspace", async ({ page }) => {
    const email = "once_leave_workspace_accept_receive_from_already_left_member@example.com"
    const inviteToken = await getInviteToken(email)
    const authInfo = await getAuthInfo(email)
    const meRes = await client.GET("/api/v1/me", { headers: authHeaders(authInfo.token) })
    expect(meRes.response.status).toBe(200)
    if (meRes.data === undefined) {
      throw new Error("meRes.data is undefined")
    }

    // set account name
    const data: components["schemas"]["User"] = {
      userId: meRes.data.me?.self.userId,
      email: meRes.data.me?.self.email,
      name: "ReceivedInvitation FromLeftMember",
      phoneNumber: ""
    }
    const updateProfileRes = await client.PUT("/api/v1/me/profile", {
      headers: authHeaders(authInfo.token),
      body: { profile: data }
    })
    expect(updateProfileRes.response.status).toBe(200)

    const res = await client.POST("/api/v1/members/invitations/{invitationId}/accept", {
      headers: authHeaders(authInfo.token),
      params: { path: { invitationId: await getInvitationIdByToken(inviteToken) } }
    })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual(
      {
        me: {
          currentWorkspace: {
            name: "Once Leave Workspace Invite",
            subdomain: "once-leave-workspace-invite",
            workspaceId: "018e3f69-4a17-7af9-bdcb-ae05aadf429c"
          },
          joinedWorkspaces: [
            {
              name: "Once Leave Workspace Invite",
              subdomain: "once-leave-workspace-invite",
              workspaceId: "018e3f69-4a17-7af9-bdcb-ae05aadf429c"
            }
          ],
          member: {
            id: res.data?.me.member?.id, // member id is created dynamically
            membershipStatus: "ACTIVE",
            profile: {
              displayName: "ReceivedInvitation FromLeftMember",
              idNumber: ""
            },
            role: "MEMBER",
            user: {
              email: "once_leave_workspace_accept_receive_from_already_left_member@example.com",
              name: "ReceivedInvitation FromLeftMember",
              userId: meRes.data.me?.self.userId
            }
          },
          providers: ["email"],
          self: {
            email: "once_leave_workspace_accept_receive_from_already_left_member@example.com",
            name: "ReceivedInvitation FromLeftMember",
            userId: meRes.data.me?.self.userId
          }
        }
      }
    )
  })
})