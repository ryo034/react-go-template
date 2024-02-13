import { expect, test } from "@playwright/test"
import { authHeaders, defaultPostHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo, getInviteToken, statefulTest } from "../../scripts"

const client = genAPIClient()

test.describe("Invite members", () => {
  // 招待時に名前が設定されておらず、まだアカウントが作成されていない場合
  test("user is not a member of any workspace and display name is not set when invited", async () => {
    const email = "invite_test_not_expired@example.com"
    const token = await getInviteToken(email)
    const res = await client.GET("/api/v1/members/invitations/verify", {
      headers: defaultPostHeaders,
      params: { query: { token } }
    })
    expect(res.response.status).toBe(200)
    expect(res.error).toBeUndefined()
    expect(res.data).toStrictEqual((await import("./no_workspace_no_display_name.json")).default)
  })

  // 招待時に名前が設定されており、まだアカウントが作成されていない場合
  test("user is not a member of any workspace and display name is set when invited", async () => {
    const email = "invite_test_not_expired_with_display_name@example.com"
    const token = await getInviteToken(email)
    const res = await client.GET("/api/v1/members/invitations/verify", {
      headers: defaultPostHeaders,
      params: { query: { token } }
    })
    expect(res.response.status).toBe(200)
    expect(res.error).toBeUndefined()
    expect(res.data).toStrictEqual((await import("./no_workspace_has_display_name.json")).default)
  })

  // 招待時に名前が設定されており、すでにアカウントが作成されている場合
  test("user is already a member of any workspace and display name is set when invited", async () => {
    const email = "invite_test_already_joined_any_workspace_with_display_name_when_invite@example.com"
    const token = await getInviteToken(email)
    const res = await client.GET("/api/v1/members/invitations/verify", {
      headers: defaultPostHeaders,
      params: { query: { token } }
    })
    expect(res.response.status).toBe(200)
    expect(res.error).toBeUndefined()
    expect(res.data).toStrictEqual((await import("./has_workspace_has_display_name.json")).default)
  })

  // 招待時に名前が設定されておらず、すでにアカウントが作成されている場合
  test("user is already member of any workspace and display name is not set when invited", async () => {
    const email = "invite_test_already_joined_any_workspace@example.com"
    const token = await getInviteToken(email)
    const res = await client.GET("/api/v1/members/invitations/verify", {
      headers: defaultPostHeaders,
      params: { query: { token } }
    })
    expect(res.response.status).toBe(200)
    expect(res.error).toBeUndefined()
    expect(res.data).toStrictEqual((await import("./has_workspace_no_display_name.json")).default)
  })

  test("failed to get invite token by invalid token", async () => {
    const res = await client.GET("/api/v1/members/invitations/verify", {
      headers: defaultPostHeaders,
      params: { query: { token: "invalid-token" } }
    })
    expect(res.response.status).toBe(400)
  })

  test("failed to get invite token by expired token", async () => {
    const res = await client.GET("/api/v1/members/invitations/verify", {
      headers: defaultPostHeaders,
      params: { query: { token: "018d96b7-587c-7614-b234-e086b1944e74" } }
    })
    expect(res.response.status).toBe(400)
    expect(res.error?.code).toBe("400-001")
  })

  // verifyOTPをする際に、招待情報があり(true)、そのワークスペースに所属していない場合、
  // そのユーザーのワークスペースに追加しメンバーにする
  // jwtにcurrentWorkspaceIdが入るが、名前がないのでオンボーディングの流れになる。
  // すでにワークスペースには所属しているので、ワークスペースの入力情報はスキップされる
  statefulTest(
    "if user is not a member of any workspace, user is added to the invited workspace @stateful",
    async () => {
      const email = "invite_test_not_expired@example.com"
      const token = await getInviteToken(email)
      const res = await client.GET("/api/v1/members/invitations/verify", {
        headers: defaultPostHeaders,
        params: { query: { token } }
      })
      expect(res.response.status).toBe(200)
      expect(res.error).toBeUndefined()
      const authInfo = await getAuthInfo(email)
      const meRes = await client.GET("/api/v1/me", { headers: authHeaders(authInfo.token) })
      expect(meRes.response.status).toBe(200)
      expect(meRes.error).toBeUndefined()
      expect(meRes.data).toStrictEqual((await import("./invite_not_a_member_of_any_workspace.json")).default)
    }
  )

  // // verifyOTPをする際に、招待情報があり(true)、そのワークスペースに所属していない場合、
  // // そのユーザーのワークスペースに追加しメンバーにする
  // // jwtのcurrentWorkspaceIdが更新され、そのままワークスペースに参加する
  // statefulTest(
  //   "if user is a member of another workspace, user is added to the invited workspace @stateful",
  //   async () => {
  //     const email = "invite_test_already_joined_any_workspace@example.com"
  //     const token = await getInviteToken(email)

  //     const res = await client.GET("/api/v1/members/invitations/verify", {
  //       headers: defaultPostHeaders,
  //       params: { query: { token } }
  //     })
  //     expect(res.response.status).toBe(200)
  //     expect(res.error).toBeUndefined()

  //     const authInfo = await getAuthInfo(email)
  //     const hs = authHeaders(authInfo.token)
  //     const meRes = await client.GET("/api/v1/me", {
  //       headers: hs
  //     })
  //     expect(meRes.response.status).toBe(200)
  //     expect(meRes.error).toBeUndefined()
  //     expect(meRes.data).toStrictEqual((await import("./invite_already_a_member_of_another_workspace.json")).default)
  //   }
  // )
})
