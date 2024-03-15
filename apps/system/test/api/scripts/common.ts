import { test } from "@playwright/test"
import { defaultPostHeaders } from "config/config"
import createClient from "openapi-fetch"
import type { paths } from "../schema/openapi/systemApi"
import { firebaseConfig } from "./config"
import { MainDb } from "./database"
import { Firebase } from "./firebase"
import { RedisClient } from "./redis"

const APIBaseURL = "http://localhost:19004"

export const statefulBeforeEach = async () => {
  const fb = new Firebase(firebaseConfig, { showConsole: false })
  const db = new MainDb()
  await Promise.all([fb.clear(), db.clear()])
  await Promise.all([fb.setup(), db.setup()])
}

export const statefulAfterEach = async () => {
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

export const getOtpCodeFromRedis = async (email: string): Promise<string> => {
  const redis = await RedisClient.create()
  const code = await redis.get(`otp:${email}`)
  if (code === null) {
    throw new Error("code is null")
  }
  return code
}

export const getAuthInfo = async (email: string): Promise<AuthInfo> => {
  const { response, error } = await client.POST("/api/v1/auth/otp", {
    headers: defaultPostHeaders,
    body: { email }
  })
  if (error !== undefined || response.status !== 200) {
    throw new Error("data is undefined")
  }
  const code = await getOtpCodeFromRedis(email)
  const verifyRes = await client.POST("/api/v1/auth/otp/verify", {
    headers: defaultPostHeaders,
    body: { email, otp: code }
  })
  if (verifyRes.error !== undefined || verifyRes.response.status !== 200) {
    console.debug(verifyRes)
    throw new Error("verifyRes.data is undefined error")
  }
  const { token } = verifyRes.data
  const fb = new Firebase(firebaseConfig, { showConsole: false })
  const tk = await fb.signInWithCustomToken(token)
  return {
    token: tk.token,
    currentWorkspaceId: (tk.claims.current_workspace_id as string) ?? ""
  }
}

export const systemTest = test.extend({
  stateful: async ({ baseURL, page }, use) => {
    await statefulBeforeEach()
    await use(page)
  }
})

export const getInviteToken = async (email: string) => {
  const db = new MainDb()
  const conn = await db.getConnection()
  const inviteesRes = await conn.query(`SELECT invitation_id FROM invitees WHERE email = '${email}'`)
  if (inviteesRes.rows.length === 0) {
    throw new Error("inviteesRes is empty")
  }
  const res = await conn.query(
    `SELECT token FROM invitation_tokens WHERE invitation_id = '${inviteesRes.rows[0].invitation_id}'`
  )
  if (res.rows.length === 0) {
    throw new Error("res is empty")
  }
  return res.rows[0].token as string
}

export const getInvitationIdByToken = async (token: string) => {
  const db = new MainDb()
  const conn = await db.getConnection()
  const res = await conn.query(`SELECT invitation_id FROM invitation_tokens WHERE token = '${token}'`)
  if (res.rows.length === 0) {
    throw new Error("tokenRes is empty")
  }
  if (res.rows.length === 0) {
    throw new Error("res is empty")
  }
  return res.rows[0].invitation_id as string
}

export const checkVerifyInvitation = async (email: string, token: string) => {
  const db = new MainDb()
  const conn = await db.getConnection()
  const res = await conn.query(`SELECT * FROM invitation_tokens WHERE token = '${token}'`)
  if (res.rows.length === 0) {
    throw new Error("res is empty")
  }
  // invitation_events
  const eventsRes = await conn.query(
    `SELECT * FROM invitation_events WHERE invitation_id = '${res.rows[0].invitation_id}' ORDER BY created_at DESC LIMIT 1`
  )
  if (eventsRes.rows.length === 0) {
    throw new Error("eventsRes is empty")
  }
  return (eventsRes.rows[0].event_type as string) === "verified"
}
