import { expect, test } from "@playwright/test"
import { defaultPostHeaders } from "../../config/config"
import { genAPIClient } from "../../scripts"

const client = genAPIClient()

test.describe("Otp API", () => {
  test.beforeEach(async ({ testIdAttribute }, testInfo) => {
    // console.log("beforeEach", )
    // console.log("beforeEach", testInfo)
  })
  test("Create Account And Verify By OTP @stateful", async () => {
    const email = "test+999@example.com"
    const { data, response, error } = await client.POST("/api/v1/auth/otp", {
      headers: defaultPostHeaders,
      body: { email }
    })
    expect(response.status).toBe(200)
    expect(error).toBeUndefined()
    if (data === undefined) {
      throw new Error("data is undefined")
    }
    const { code } = data
    expect(code).toMatch(/^\d{6}$/)

    const verifyRes = await client.POST("/api/v1/auth/otp/verify", {
      headers: defaultPostHeaders,
      body: { email, otp: code }
    })
    expect(verifyRes.response.status).toBe(200)
    expect(verifyRes.error).toBeUndefined()
    if (verifyRes.data === undefined) {
      throw new Error("verifyRes.data is undefined")
    }
    const { token } = verifyRes.data
    expect(token.length).toBeGreaterThan(0)
  })

  // test("If the user already exists @stateful", async () => {
  //   const email = "test+1@example.com"
  //   const authRes = await client.POST("/api/v1/otp/auth", {
  //     headers: { "content-type": "application/json" },
  //     body: { email }
  //   })
  //   expect(authRes.response.status).toBe(200)
  //   expect(authRes.error).toBeUndefined()
  //   const { code } = authRes.data
  //   expect(code).toMatch(/^\d{6}$/)

  //   const verifyRes = await client.POST("/api/v1/otp/verify", {
  //     headers: { "content-type": "application/json" },
  //     body: { email, otp: code }
  //   })
  //   expect(verifyRes.response.status).toBe(200)
  //   expect(verifyRes.error).toBeUndefined()
  //   expect(verifyRes.data.token.length).toBeGreaterThan(0)
  // })
})
