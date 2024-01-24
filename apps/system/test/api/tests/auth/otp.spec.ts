import { expect, test } from "@playwright/test"
import { defaultPostHeaders } from "../../config/config"
import { genAPIClient } from "../../scripts"

const client = genAPIClient()

test.describe("Otp API", () => {
  test("Create Account And Verify By OTP @stateful", async () => {
    const email = "test@example.com"
    const { data, response, error } = await client.POST("/api/v1/otp/auth", {
      headers: defaultPostHeaders,
      body: { email }
    })
    expect(response.status).toBe(200)
    expect(error).toBeUndefined()
    const { code } = data
    expect(code).toMatch(/^\d{6}$/)

    const verifyRes = await client.POST("/api/v1/otp/verify", {
      headers: defaultPostHeaders,
      body: { email, otp: code }
    })
    expect(verifyRes.response.status).toBe(200)
    expect(verifyRes.error).toBeUndefined()
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