import { expect, test } from "@playwright/test"
import { authHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo } from "../../scripts"
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

test.describe("Workspace", () => {
  test("Owner can not leave workspace", async () => {
    const authInfo = await getAuthInfo("me_owner_leave_workspace@example.com")
    const res = await client.POST('/api/v1/me/workspace/leave', {
      headers: authHeaders(authInfo.token),
    })
    expect(res.response.status).toBe(403)
  })

  test("Admin can leave workspace", async () => {
    const authInfo = await getAuthInfo("me_admin_leave_workspace@example.com")
    const meRes = await client.GET('/api/v1/me', { headers: authHeaders(authInfo.token) })
    expect(meRes.data?.me.currentWorkspace?.workspaceId).toBe('018e3642-86c6-7e74-b874-9c2835b2ce50')
    expect(meRes.data?.me.member?.id).toBe('018e3642-86c6-76da-a51f-05bb129508b6')

    const res = await client.POST('/api/v1/me/workspace/leave', { headers: authHeaders(authInfo.token) })
    expect(res.response.status).toBe(204)

    const meRes2 = await client.GET('/api/v1/me', { headers: authHeaders(authInfo.token) })
    expect(meRes2.data?.me.currentWorkspace).toBeUndefined()
    expect(meRes2.data?.me.member).toBeUndefined()
  })

  test("Member can leave workspace", async () => {
    const authInfo = await getAuthInfo("me_member_leave_workspace@example.com")
    const meRes = await client.GET('/api/v1/me', { headers: authHeaders(authInfo.token) })
    expect(meRes.data?.me.currentWorkspace?.workspaceId).toBe('018e3642-86c6-7e74-b874-9c2835b2ce50')
    expect(meRes.data?.me.member?.id).toBe('018e3642-86c6-7cfc-b321-62d665d62c8b')

    const res = await client.POST('/api/v1/me/workspace/leave', { headers: authHeaders(authInfo.token) })
    expect(res.response.status).toBe(204)

    const meRes2 = await client.GET('/api/v1/me', { headers: authHeaders(authInfo.token) })
    expect(meRes2.data?.me.currentWorkspace).toBeUndefined()
    expect(meRes2.data?.me.member).toBeUndefined()
  })

  test("Guest can leave workspace", async () => {
    const authInfo = await getAuthInfo("me_guest_leave_workspace@example.com")
    const meRes = await client.GET('/api/v1/me', { headers: authHeaders(authInfo.token) })
    expect(meRes.data?.me.currentWorkspace?.workspaceId).toBe('018e3642-86c6-7e74-b874-9c2835b2ce50')
    expect(meRes.data?.me.member?.id).toBe('018e3642-86c6-7483-8c30-f017d34eed3a')

    const res = await client.POST('/api/v1/me/workspace/leave', { headers: authHeaders(authInfo.token) })
    expect(res.response.status).toBe(204)

    const meRes2 = await client.GET('/api/v1/me', { headers: authHeaders(authInfo.token) })
    expect(meRes2.data?.me.currentWorkspace).toBeUndefined()
    expect(meRes2.data?.me.member).toBeUndefined()
  })
})
