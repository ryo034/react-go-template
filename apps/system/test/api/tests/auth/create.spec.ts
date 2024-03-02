import { expect, test } from "@playwright/test"
import { defaultPostHeaders } from "../../config/config"
import { genAPIClient, getOtpCodeFromRedis, systemTest } from "../../scripts"

const client = genAPIClient()

systemTest.describe("Create Account", () => {
  systemTest("Create Account And Verify By OTP", { tag: ["@stateful"] }, async ({ stateful }) => {
    const email = "test+999@example.com"
    const { response, error } = await client.POST("/api/v1/auth/otp", {
      headers: defaultPostHeaders,
      body: { email }
    })
    expect(response.status).toBe(200)
    expect(error).toBeUndefined()

    const code = await getOtpCodeFromRedis(email)

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
