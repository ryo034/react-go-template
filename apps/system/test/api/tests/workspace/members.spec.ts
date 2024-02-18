import { expect, test } from "@playwright/test"
import { authHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo } from "../../scripts"

const client = genAPIClient()

test.describe("Workspace members", () => {
  test("get workspace members", async () => {
    const authInfo = await getAuthInfo("system_account@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.GET("/api/v1/members", {
      headers: hs
    })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./success_get_members.json")).default)
  })
})
