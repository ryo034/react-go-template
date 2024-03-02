import { expect } from "@playwright/test"
import { authHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo, systemTest } from "../../scripts"

const client = genAPIClient()

systemTest.describe("Create Workspace", () => {
  systemTest("Workspace can not be created with already exists subdomain", async () => {
    const authInfo = await getAuthInfo("system_account@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.POST("/api/v1/workspaces", {
      headers: hs,
      body: { subdomain: "example" }
    })
    expect(res.response.status).toBe(409)
  })

  systemTest("Workspace can be created with already exists name", { tag: ["@stateful"] }, async ({ stateful }) => {
    const authInfo = await getAuthInfo("system_account@example.com")
    const res = await client.POST("/api/v1/workspaces", {
      headers: authHeaders(authInfo.token),
      body: { subdomain: "test-example-subdomain" }
    })
    expect(res.response.status).toBe(201)
    expect(res.error).toBeUndefined()
  })
})
