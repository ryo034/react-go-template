import { expect, test } from "@playwright/test"
import { headers } from "../../config/config"
import { genAPIClient, getAuthInfo } from "../../scripts"
const client = genAPIClient()

test.describe("Me success", () => {
  test("success", async () => {
    const authInfo = await getAuthInfo("system_account@example.com")
    const hs = headers(authInfo.token)
    const res = await client.GET("/api/v1/me", { headers: hs })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./success.json")).default)
  })
})
