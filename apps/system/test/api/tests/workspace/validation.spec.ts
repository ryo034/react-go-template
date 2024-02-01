import { expect, test } from "@playwright/test"
import { headers } from "../../config/config"
import { genAPIClient, getToken, statefulTest } from "../../scripts"

const client = genAPIClient()

test.describe("Workspace Validation", () => {
  statefulTest("Workspace can be created with already exists name @stateful", async ({ page }) => {
    const token = await getToken("system_account@example.com")
    const res = await client.POST("/api/v1/workspaces", {
      headers: headers(token),
      body: {
        name: "Example",
        subdomain: "test-example-subdomain"
      }
    })
    expect(res.response.status).toBe(201)
    expect(res.error).toBeUndefined()
  })

  test("Workspace can not be created with already exists subdomain", async () => {
    const token = await getToken("system_account@example.com")
    const hs = headers(token)
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
