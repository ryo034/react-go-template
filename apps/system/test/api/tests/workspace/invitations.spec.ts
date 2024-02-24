import { expect, test } from "@playwright/test"
import { authHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo, statefulTest } from "../../scripts"

const client = genAPIClient()

test.describe("invite members", () => {
  statefulTest("success to invite members @stateful", async ({ page }) => {
    const email = "system_account@example.com"
    const authInfo = await getAuthInfo(email)
    const res = await client.POST("/api/v1/members/invitations/bulk", {
      headers: authHeaders(authInfo.token),
      body: {
        invitees: [
          { email: "invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com", name: "test" }
        ]
      }
    })
    expect(res.response.status).toBe(200)
    expect(res.data?.total).toBe(1)
    expect(res.data?.successfulInvitations.length).toBe(1)
    expect(res.data?.failedInvitations.length).toBe(0)
    expect(
      res.data?.successfulInvitations.filter(
        (i) => i.inviteeEmail === "invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com"
      ).length
    ).toBe(1)
  })
  statefulTest("Already joined any workspace @stateful", async ({ page }) => {
    const email = "invite_test_already_joined_any_workspace@example.com"
    const authInfo = await getAuthInfo(email)
    const res = await client.POST("/api/v1/members/invitations/bulk", {
      headers: authHeaders(authInfo.token),
      body: {
        invitees: [
          { email: "invite_test_not_exist@example.com", name: "Not Exist" },
          { email: "invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com", name: "test" }
        ]
      }
    })
    expect(res.response.status).toBe(200)
    expect(res.data?.total).toBe(2)
    expect(res.data?.successfulInvitations.length).toBe(1)
    expect(res.data?.failedInvitations.length).toBe(0)
    expect(
      res.data?.registeredInvitations.filter(
        (i) => i.inviteeEmail === "invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com"
      ).length
    ).toBe(1)
  })
  test("Contains invalid email", async () => {
    const email = "invite_test_already_joined_any_workspace@example.com"
    const authInfo = await getAuthInfo(email)
    const res = await client.POST("/api/v1/members/invitations/bulk", {
      headers: authHeaders(authInfo.token),
      body: {
        invitees: [{ email: "invalid_example_test", name: "Invalid Email" }]
      }
    })
    expect(res.response.status).toBe(400)
  })
})

test.describe("get invitations", () => {
  test("success to get invitations without revoked and already registered", async () => {
    const email = "system_account@example.com"
    const authInfo = await getAuthInfo(email)
    const res = await client.GET("/api/v1/invitations", { headers: authHeaders(authInfo.token) })
    expect(res.response.status).toBe(200)
    expect(res.error).toBeUndefined()
    expect(JSON.stringify(res.data)).toEqual(JSON.stringify((await import("./success_get_invitations.json")).default))
  })
  statefulTest(
    "success to get verified invitations without revoked and already registered @stateful",
    async ({ page }) => {
      const email = "system_account@example.com"
      const authInfo = await getAuthInfo(email)
      const res = await client.GET("/api/v1/invitations", {
        headers: authHeaders(authInfo.token),
        params: { query: { status: "accepted" } }
      })
      expect(res.response.status).toBe(200)
      expect(res.error).toBeUndefined()
      expect(res.data).toStrictEqual((await import("./success_get_accepted_invitations.json")).default)
    }
  )
})
