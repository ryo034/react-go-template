import { expect, test } from "@playwright/test"
import { components } from "schema/openapi/systemApi"
import { authHeaders } from "../../config/config"
import { genAPIClient, getAuthInfo, statefulTest } from "../../scripts"
const client = genAPIClient()

test.describe("Update me success", () => {
  statefulTest("update me @stateful", async () => {
    const authInfo = await getAuthInfo("system_account@example.com")
    const data: components["schemas"]["User"] = {
      userId: "394e67b6-2850-4ddf-a4c9-c2a619d5bf70",
      email: "system_account@example.com",
      name: "Updated Name",
      phoneNumber: ""
    }
    const res = await client.PUT("/api/v1/me/profile", { headers: authHeaders(authInfo.token), body: { user: data } })
    expect(res.response.status).toBe(200)
    expect(res.data).toStrictEqual((await import("./update_me.json")).default)
  })
})
