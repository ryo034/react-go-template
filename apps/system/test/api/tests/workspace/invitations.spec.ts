import { expect, test } from "@playwright/test"
import { defaultPostHeaders } from "../../config/config"
import { genAPIClient, getInviteToken, statefulTest } from "../../scripts"

const client = genAPIClient()

test.describe("Invite members", () => {
  test("success get invite token", async () => {
    const email = "invite_test_not_expired@example.com"
    const token = await getInviteToken(email)
    const res = await client.GET("/api/v1/members/invitations/verify", {
      headers: defaultPostHeaders,
      params: { query: { token } }
    })
    expect(res.response.status).toBe(200)
    expect(res.error).toBeUndefined()
    expect(res.data).toStrictEqual((await import("./verify_invite_token.json")).default)
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

      // const authInfo = await getAuthInfo(email)
      // const hs = authHeaders(authInfo.token)
      // const meRes = await client.GET("/api/v1/me", {
      //   headers: hs
      // })
      // expect(meRes.response.status).toBe(200)
      // expect(meRes.error).toBeUndefined()
      // expect(meRes.data).toStrictEqual((await import("./invite_not_a_member_of_any_workspace.json")).default)
    }
  )

  // // verifyOTPをする際に、招待情報があり(true)、そのワークスペースに所属していない場合、
  // // そのユーザーのワークスペースに追加しメンバーにする
  // // jwtのcurrentWorkspaceIdが更新され、そのままワークスペースに参加する
  // statefulTest(
  //   "if user is a member of another workspace, user is added to the invited workspace @stateful",
  //   async () => {
  //     const email = "invite_test_2@example.com"
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
