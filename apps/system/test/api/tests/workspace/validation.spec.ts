import { expect, test } from "@playwright/test"
import { authHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo, statefulTest } from "../../scripts"

const client = genAPIClient()

test.describe("Workspace Validation", () => {
  statefulTest("Workspace can be created with already exists name @stateful", async ({ page }) => {
    const authInfo = await getAuthInfo("system_account@example.com")
    const res = await client.POST("/api/v1/workspaces", {
      headers: authHeaders(authInfo.token),
      body: {
        name: "Example",
        subdomain: "test-example-subdomain"
      }
    })
    expect(res.response.status).toBe(201)
    expect(res.error).toBeUndefined()
  })

  test("Workspace can not be created with already exists subdomain", async () => {
    const authInfo = await getAuthInfo("system_account@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.POST("/api/v1/workspaces", {
      headers: hs,
      body: {
        name: "test",
        subdomain: "example"
      }
    })
    expect(res.response.status).toBe(409)
  })
})