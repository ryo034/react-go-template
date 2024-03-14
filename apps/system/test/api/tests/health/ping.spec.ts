import { expect, test } from "@playwright/test"
import { defaultPostHeaders } from "../../config/config"
import { genAPIClient } from "../../scripts"

const client = genAPIClient()

test.describe("Otp API", () => {
  test("Create Account OK", async () => {
    const res = await client.GET("/ping", {
      headers: defaultPostHeaders
    })
    expect(res.response.status).toBe(200)
    expect(res.error).toBeUndefined()
    expect(res.data).toEqual({ message: "pong" })
  })
})
