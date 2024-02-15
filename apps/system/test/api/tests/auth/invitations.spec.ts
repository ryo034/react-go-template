import { expect, test } from "@playwright/test"
import { defaultPostHeaders } from "../../config/config"
import { checkVerifyInvitation, genAPIClient, getInviteToken, statefulTest } from "../../scripts"

const client = genAPIClient()

test.describe("Invalidations", () => {
  test("send invalid token, return BadRequestError", async () => {
    const email = "invite_test_not_expired@example.com"
    const processRes = await client.POST("/api/v1/auth/invitations/process", {
      headers: defaultPostHeaders,
      body: { token: "018d96b7-587c-7614-b234-e086b1944e79", email }
    })
    expect(processRes.response.status).toBe(400)
    expect(processRes.error?.code).toBe("400-000")
    expect(processRes.error?.title).toBe("不正な招待トークンです")
  })

  statefulTest("send valid token and email, return 200 and invite verified @stateful", async ({ page }) => {
    const email = "invite_test_not_expired@example.com"
    const token = await getInviteToken(email)
    const processRes = await client.POST("/api/v1/auth/invitations/process", {
      headers: defaultPostHeaders,
      body: { token, email }
    })
    expect(processRes.response.status).toBe(200)
    expect(processRes.error).toBeUndefined()
    expect(await checkVerifyInvitation(email, token)).toBeTruthy()
  })
})
