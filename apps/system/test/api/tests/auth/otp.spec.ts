import { expect, test } from "@playwright/test"
import { defaultPostHeaders } from "../../config/config"
import { genAPIClient, statefulTest } from "../../scripts"

const client = genAPIClient()

test.describe("Otp API", () => {
  statefulTest("Create Account And Verify By OTP @stateful", async () => {
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
})
