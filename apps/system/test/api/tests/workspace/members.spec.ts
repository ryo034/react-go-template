import { expect, test } from "@playwright/test"
import { authHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo, systemTest } from "../../scripts"

const client = genAPIClient()

test.describe("Workspace members", () => {
  test("get workspace members", async () => {
    const authInfo = await getAuthInfo("account@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.GET("/api/v1/members", { headers: hs })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./success_get_members.json")).default)
  })

  test("get workspace members exclude left member", async () => {
    const authInfo = await getAuthInfo("once_leave_check_left_workspace_owner@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.GET("/api/v1/members", { headers: hs })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./get_workspace_members_exclude_left_member.json")).default)
  })
})

test.describe("Member", () => {
  test("Can not update to same role", async () => {
    const authInfo = await getAuthInfo("update_role_owner@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.PUT("/api/v1/members/{memberId}/role", {
      headers: hs,
      body: { role: "admin" },
      params: { path: { memberId: "018e18ba-dc87-72e2-bb4b-c43252f51492" } }
    })
    expect(res.response.status).toBe(403)
  })
  test("Update Admin role to Member Role", async () => {
    const authInfo = await getAuthInfo("update_role_admin@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.PUT("/api/v1/members/{memberId}/role", {
      headers: hs,
      body: { role: "member" },
      params: { path: { memberId: "018e18ba-dc87-740c-9aeb-ba7f8f7d490e" } }
    })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./success_update_member_to_member_role.json")).default)
  })
  test("Update Admin role to Guest Role", async () => {
    const authInfo = await getAuthInfo("update_role_admin_2@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.PUT("/api/v1/members/{memberId}/role", {
      headers: hs,
      body: { role: "guest" },
      params: { path: { memberId: "018e1952-009b-7138-aea6-24b2f9596ad7" } }
    })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./success_update_member_to_guest_role.json")).default)
  })
})
