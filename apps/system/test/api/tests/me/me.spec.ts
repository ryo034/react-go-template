import { expect, test } from "@playwright/test"
import { authHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo, systemTest } from "../../scripts"
const client = genAPIClient()

test.describe("Me success", () => {
  test("if onboarding is completed, response include workspace info", async () => {
    const authInfo = await getAuthInfo("account@example.com")
    const res = await client.GET("/api/v1/me", { headers: authHeaders(authInfo.token) })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./success.json")).default)
  })

  test("if onboarding is not completed, response not include workspace info", async () => {
    const authInfo = await getAuthInfo("unfinished_onboarding@example.com")
    expect(authInfo.currentWorkspaceId).toBe("")
    const res = await client.GET("/api/v1/me", { headers: authHeaders(authInfo.token) })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./unfinished_onboarding.json")).default)
  })
})

systemTest("Workspace", () => {
  systemTest("Success Leave Workspace", async ({ stateful }) => {
    const authInfo = await getAuthInfo("unfinished_onboarding@example.com")
  })
})
