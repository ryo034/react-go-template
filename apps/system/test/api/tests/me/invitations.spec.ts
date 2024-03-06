import { expect } from "@playwright/test"
import { components } from "schema/openapi/systemApi"
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
              id: "g57lunkvmbhurkm5dhf5nkblbu",
              profile: {
                bio: "John Doe is a passionate software engineer with 8 years of experience specializing in web development, particularly with React and Node.js. A graduate from MIT with a strong focus on clean architecture and Agile methodologies, John has successfully led multiple projects, from innovative startups to established tech giants. He's a firm believer in continuous learning, contributing regularly to open-source projects, and sharing insights through tech blogs and meetups. Outside of work, John enjoys hiking 🚶‍♂️, drone photography 📸, and playing the guitar 🎸. He's committed to using technology to drive positive social change.",
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
              id: "g57lunkvmbhurkm5dhf5nkblbu",
              profile: {
                displayName: "John Doe",
                idNumber: "DEV-12345",
                bio: "John Doe is a passionate software engineer with 8 years of experience specializing in web development, particularly with React and Node.js. A graduate from MIT with a strong focus on clean architecture and Agile methodologies, John has successfully led multiple projects, from innovative startups to established tech giants. He's a firm believer in continuous learning, contributing regularly to open-source projects, and sharing insights through tech blogs and meetups. Outside of work, John enjoys hiking 🚶‍♂️, drone photography 📸, and playing the guitar 🎸. He's committed to using technology to drive positive social change."
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
