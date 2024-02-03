import { test } from "@playwright/test"
import { defaultPostHeaders, headers } from "config/config"
import createClient from "openapi-fetch"
import { paths } from "../schema/openapi/systemApi"
import { firebaseConfig } from "./config"
import { MainDb } from "./database"
import { Firebase } from "./firebase"

const APIBaseURL = "http://localhost:19004"

export const statefulBeforeEach = async () => {
  const fb = new Firebase(firebaseConfig, { showConsole: false })
  const db = new MainDb()
  await Promise.all([fb.clear(), db.clear()])
  await Promise.all([fb.setup(), db.setup()])
}

export const genAPIClient = () => {
  return createClient<paths>({ baseUrl: APIBaseURL })
}

const client = createClient<paths>({ baseUrl: APIBaseURL })

type AuthInfo = {
  token: string
  currentWorkspaceId: string
}

export const getAuthInfo = async (email: string): Promise<AuthInfo> => {
  const { data, response, error } = await client.POST("/api/v1/auth/otp", {
    headers: defaultPostHeaders,
    body: { email }
  })
  if (data === undefined) {
    throw new Error("data is undefined")
  }
  const { code } = data
  const verifyRes = await client.POST("/api/v1/auth/otp/verify", {
    headers: defaultPostHeaders,
    body: { email, otp: code }
  })
  if (verifyRes.data === undefined) {
    throw new Error("verifyRes.data is undefined")
  }
  const { token } = verifyRes.data
  const fb = new Firebase(firebaseConfig, { showConsole: false })
  const tk = await fb.signInWithCustomToken(token)
  return {
    token: tk.token,
    currentWorkspaceId: (tk.claims.current_workspace_id as string) ?? ""
  }
}

export const statefulTest = test.extend({
  page: async ({ baseURL, page }, use) => {
    await statefulBeforeEach()
    await use(page)
  }
})
