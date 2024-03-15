import test, { expect } from "@playwright/test"
import { authHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo, systemTest } from "../../scripts"

const client = genAPIClient()

systemTest.describe("Create Workspace", () => {
  systemTest("Workspace can not be created with already exists subdomain", async () => {
    const authInfo = await getAuthInfo("account@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.POST("/api/v1/workspaces", {
      headers: hs,
      body: { subdomain: "example" }
    })
    expect(res.response.status).toBe(409)
  })

  systemTest("Workspace can be created with already exists name", { tag: ["@stateful"] }, async ({ stateful }) => {
    const authInfo = await getAuthInfo("account@example.com")
    const res = await client.POST("/api/v1/workspaces", {
      headers: authHeaders(authInfo.token),
      body: { subdomain: "test-example-subdomain" }
    })
    expect(res.response.status).toBe(201)
    expect(res.error).toBeUndefined()
  })
})

test.describe("Update workspace", () => {
  test("Admin role can update workspace detail", async () => {
    const authInfo = await getAuthInfo("update_workspace_detail@example.com")
    const res = await client.PUT("/api/v1/workspaces/{workspaceId}", {
      headers: authHeaders(authInfo.token),
      params: { path: { workspaceId: "018e201b-67d4-7265-a022-1b29793b2a91" } },
      body: { name: "Update TestUpdated", subdomain: "update-test-updated" }
    })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./success_update_workspace.json")).default)
  })

  test("Member role can not update workspace detail", async () => {
    const authInfo = await getAuthInfo("update_workspace_detail_member_role@example.com")
    const res = await client.PUT("/api/v1/workspaces/{workspaceId}", {
      headers: authHeaders(authInfo.token),
      params: { path: { workspaceId: "018e201b-67d4-7265-a022-1b29793b2a91" } },
      body: { name: "Update TestUpdated", subdomain: "update-test-updated" }
    })
    expect(res.response.status).toBe(403)
  })

  test("Guest role can not update workspace detail", async () => {
    const authInfo = await getAuthInfo("update_workspace_detail_guest_role@example.com")
    const res = await client.PUT("/api/v1/workspaces/{workspaceId}", {
      headers: authHeaders(authInfo.token),
      params: { path: { workspaceId: "018e201b-67d4-7265-a022-1b29793b2a91" } },
      body: { name: "Update TestUpdated", subdomain: "update-test-updated" }
    })
    expect(res.response.status).toBe(403)
  })
})

test.describe("Leave workspace", () => {
  test("get workspace members exclude left member", async () => {
    const authInfo = await getAuthInfo("once_leave_workspace_invite_owner@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.GET("/api/v1/members", { headers: hs })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./get_workspace_members_exclude_left_member.json")).default)
  })
})
