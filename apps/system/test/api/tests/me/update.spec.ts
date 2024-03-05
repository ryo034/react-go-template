import { expect } from "@playwright/test"
import { components } from "schema/openapi/systemApi"
import { authHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo, systemTest } from "../../scripts"
const client = genAPIClient()

systemTest.describe("Update me success", () => {
  systemTest("update name", { tag: ["@stateful"] }, async ({ stateful }) => {
    const authInfo = await getAuthInfo("account@example.com")
    const data: components["schemas"]["User"] = {
      userId: "394e67b6-2850-4ddf-a4c9-c2a619d5bf70",
      email: "account@example.com",
      name: "Updated Name",
      phoneNumber: ""
    }
    const res = await client.PUT("/api/v1/me/profile", {
      headers: authHeaders(authInfo.token),
      body: { profile: data }
    })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./update_me.json")).default)
  })
})

systemTest.describe("Update me member profile success", () => {
  systemTest("update me member profile", { tag: ["@stateful"] }, async ({ stateful }) => {
    const authInfo = await getAuthInfo("account@example.com")
    const meRes = await client.GET("/api/v1/me", { headers: authHeaders(authInfo.token) })
    expect(meRes.data).toStrictEqual((await import("./update_me_member_get_me.json")).default)

    const data: components["schemas"]["MemberProfile"] = {
      displayName: "Updated Display Name",
      bio: "Updated Bio",
      idNumber: "1234567890"
    }
    const res = await client.PUT("/api/v1/me/member/profile", {
      headers: authHeaders(authInfo.token),
      body: { memberProfile: data }
    })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./update_me_member_success.json")).default)
  })

  systemTest("success if request has empty fields", { tag: ["@stateful"] }, async ({ stateful }) => {
    const authInfo = await getAuthInfo("account@example.com")
    const meRes = await client.GET("/api/v1/me", { headers: authHeaders(authInfo.token) })
    expect(meRes.data).toStrictEqual((await import("./update_me_member_get_me.json")).default)
    const data: components["schemas"]["MemberProfile"] = {
      displayName: "Updated Display Name"
    }
    const res = await client.PUT("/api/v1/me/member/profile", {
      headers: authHeaders(authInfo.token),
      body: { memberProfile: data }
    })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./update_me_member_success_has_empty_bio_idnumber.json")).default)

    // empty display name
    const data2: components["schemas"]["MemberProfile"] = {
      displayName: ""
    }
    const emptyDisplayNameRes = await client.PUT("/api/v1/me/member/profile", {
      headers: authHeaders(authInfo.token),
      body: { memberProfile: data2 }
    })
    expect(emptyDisplayNameRes.response.status).toBe(200)
    expect(emptyDisplayNameRes.data).toStrictEqual(
      (await import("./update_me_member_success_has_empty_display_name.json")).default
    )
  })
})
