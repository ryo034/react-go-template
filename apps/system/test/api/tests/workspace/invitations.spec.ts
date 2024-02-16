import { expect, test } from "@playwright/test"
import { authHeaders, defaultPostHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo, getInviteToken, statefulTest } from "../../scripts"

const client = genAPIClient()

test.describe("verify invitations", () => {
  test("user is already member of any workspace and display name is not set when invited", async () => {
    const email = "invite_test_already_joined_any_workspace@example.com"
    const token = await getInviteToken(email)
    const res = await client.GET("/api/v1/members/invitations/verify", {
      headers: defaultPostHeaders,
      params: { query: { token } }
    })
    expect(res.response.status).toBe(200)
    expect(res.error).toBeUndefined()
    expect(res.data).toStrictEqual((await import("./success_invitations_verify.json")).default)
  })

  test("failed to get invite token by invalid token return BadRequestError", async () => {
    const res = await client.GET("/api/v1/members/invitations/verify", {
      headers: defaultPostHeaders,
      params: { query: { token: "invalid-token" } }
    })
    expect(res.response.status).toBe(400)
  })

  test("failed to get invite token by expired token return BadRequestError with custom code", async () => {
    const res = await client.GET("/api/v1/members/invitations/verify", {
      headers: defaultPostHeaders,
      params: { query: { token: "018d96b7-587c-7614-b234-e086b1944e74" } }
    })
    expect(res.response.status).toBe(400)
    expect(res.error?.code).toBe("400-001")
  })
})

test.describe("invite members", () => {
  statefulTest("success to invite members @stateful", async () => {
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
  statefulTest("Already joined any workspace @stateful", async () => {
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
