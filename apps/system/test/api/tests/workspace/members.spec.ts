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
    const authInfo = await getAuthInfo("once_leave_workspace_invite_owner@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.GET("/api/v1/members", { headers: hs })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./get_workspace_members_exclude_left_member.json")).default)
  })
})

systemTest.describe("Member", () => {
  systemTest("Can not update to same role", { tag: ["@stateful"] }, async () => {
    const authInfo = await getAuthInfo("account@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.PUT("/api/v1/members/{memberId}/role", {
      headers: hs,
      body: { role: "admin" },
      params: { path: { memberId: "018e1398-3d80-79dc-9459-c7a3f1609124" } }
    })
    expect(res.response.status).toBe(400)
  })
  systemTest("Update Admin role to Member Role", { tag: ["@stateful"] }, async () => {
    const authInfo = await getAuthInfo("account@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.PUT("/api/v1/members/{memberId}/role", {
      headers: hs,
      body: { role: "member" },
      params: { path: { memberId: "018e1398-3d80-79dc-9459-c7a3f1609124" } }
    })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./success_update_member_to_member_role.json")).default)
  })
  systemTest("Update Admin role to Guest Role", { tag: ["@stateful"] }, async () => {
    const authInfo = await getAuthInfo("account@example.com")
    const hs = authHeaders(authInfo.token)
    const res = await client.PUT("/api/v1/members/{memberId}/role", {
      headers: hs,
      body: { role: "guest" },
      params: { path: { memberId: "018e1398-3d80-79dc-9459-c7a3f1609124" } }
    })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./success_update_member_to_guest_role.json")).default)
  })
})
