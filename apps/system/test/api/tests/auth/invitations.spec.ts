import { expect, test } from "@playwright/test"
import { defaultPostHeaders } from "../../config/config"
import { checkVerifyInvitation, genAPIClient, getInviteToken, systemTest } from "../../scripts"

const client = genAPIClient()

systemTest.describe("Invalidations", () => {
  systemTest("send invalid token, return BadRequestError", async () => {
    const email = "invite_test_not_expired@example.com"
    const processRes = await client.POST("/api/v1/auth/invitations/process/email", {
      headers: defaultPostHeaders,
      body: { token: "018d96b7-587c-7614-b234-e086b1944e79", email }
    })
    expect(processRes.response.status).toBe(400)
    expect(processRes.error?.code).toBe("400-000")
    expect(processRes.error?.title).toBe("不正な招待トークンです")
  })

  systemTest(
    "send valid token and email, return 200 and invite verified",
    { tag: ["@stateful"] },
    async ({ stateful: page }) => {
      const email = "invite_test_not_expired@example.com"
      const token = await getInviteToken(email)
      const processRes = await client.POST("/api/v1/auth/invitations/process/email", {
        headers: defaultPostHeaders,
        body: { token, email }
      })
      expect(processRes.response.status).toBe(200)
      expect(processRes.error).toBeUndefined()
      expect(await checkVerifyInvitation(email, token)).toBeTruthy()
    }
  )
})

test.describe("auth invitations", () => {
  test("user is already member of any workspace and display name is not set when invited", async () => {
    const email = "invite_test_already_joined_any_workspace@example.com"
    const token = await getInviteToken(email)
    const res = await client.GET("/api/v1/auth/invitations", {
      headers: defaultPostHeaders,
      params: { query: { token } }
    })
    expect(res.response.status).toBe(200)
    expect(res.error).toBeUndefined()
    expect(res.data).toStrictEqual((await import("./success_get_auth_invitation.json")).default)
  })

  test("failed to get invite token by invalid token return BadRequestError", async () => {
    const res = await client.GET("/api/v1/auth/invitations", {
      headers: defaultPostHeaders,
      params: { query: { token: "invalid-token" } }
    })
    expect(res.response.status).toBe(400)
  })

  test("failed to get invite token by already expired token return GoneError with custom code", async () => {
    const res = await client.GET("/api/v1/auth/invitations", {
      headers: defaultPostHeaders,
      params: { query: { token: "018d96b7-587c-7614-b234-e086b1944e74" } }
    })
    expect(res.response.status).toBe(410)
    expect(res.error?.code).toBe("410-001")
  })
  test("failed to get invite token by already revoked token return GoneError with custom code", async () => {
    const res = await client.GET("/api/v1/auth/invitations", {
      headers: defaultPostHeaders,
      params: { query: { token: "018dbe28-a7c6-7b51-885a-7c4647e4aff4" } }
    })
    expect(res.response.status).toBe(410)
    expect(res.error?.code).toBe("410-002")
  })
  test("success to get invite token by already verified token", async () => {
    const res = await client.GET("/api/v1/auth/invitations", {
      headers: defaultPostHeaders,
      params: { query: { token: "018dcee9-4ec8-7f93-9a9c-f9ad7ae3d592" } }
    })
    expect(res.response.status).toBe(200)
    expect(res.error).toBeUndefined()
    expect(res.data).toStrictEqual((await import("./success_get_auth_invitation_by_verified_token.json")).default)
  })
  test("failed to get invite token by already accepted token return GoneError with custom code", async () => {
    const res = await client.GET("/api/v1/auth/invitations", {
      headers: defaultPostHeaders,
      params: { query: { token: "018d96bb-975d-769c-aa3d-dfe09fc9f207" } }
    })
    expect(res.response.status).toBe(410)
    expect(res.error?.code).toBe("410-003")
  })
})
