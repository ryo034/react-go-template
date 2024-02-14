import { expect, test } from "@playwright/test"
import { defaultPostHeaders } from "../../config/config"
import { genAPIClient, getInviteToken } from "../../scripts"

const client = genAPIClient()

test.describe("Invite members", () => {
  test("user is already member of any workspace and display name is not set when invited", async () => {
    const email = "invite_test_already_joined_any_workspace@example.com"
    const token = await getInviteToken(email)
    const res = await client.GET("/api/v1/members/invitations/verify", {
      headers: defaultPostHeaders,
      params: { query: { token } }
    })
    expect(res.response.status).toBe(200)
    expect(res.error).toBeUndefined()
    expect(res.data).toStrictEqual((await import("./success_invitations_verify.json")).default)
  })

  test("failed to get invite token by invalid token return BadRequestError", async () => {
    const res = await client.GET("/api/v1/members/invitations/verify", {
      headers: defaultPostHeaders,
      params: { query: { token: "invalid-token" } }
    })
    expect(res.response.status).toBe(400)
  })

  test("failed to get invite token by expired token return BadRequestError with custom code", async () => {
    const res = await client.GET("/api/v1/members/invitations/verify", {
      headers: defaultPostHeaders,
      params: { query: { token: "018d96b7-587c-7614-b234-e086b1944e74" } }
    })
    expect(res.response.status).toBe(400)
    expect(res.error?.code).toBe("400-001")
  })

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
