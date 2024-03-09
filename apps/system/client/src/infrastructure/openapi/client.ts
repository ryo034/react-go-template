import createClient, { Middleware } from "openapi-fetch"
import { paths } from "~/generated/schema/openapi/systemApi"
import { firebaseAuth } from "~/infrastructure/firebase"

const fetchRequestInterceptor: Middleware = {
  async onRequest(req, options) {
    if (firebaseAuth.currentUser === null) {
      return req
    }
    const token = await firebaseAuth.currentUser.getIdToken()
    if (!token) {
      return req
    }
    req.headers.set("Authorization", `Bearer ${token}`)
    return req
  }
}

const fetchResponseInterceptor: Middleware = {
  async onResponse(res, options) {
    if (res.status === 401) {
      await firebaseAuth.signOut()
    }
    return res
  }
}

const debugMiddleware: Middleware = {
  async onRequest(req, options) {
    console.log("onRequest", req, options)
    return req
  },
  async onResponse(res, options) {
    console.log("onResponse", res, options)
    return res
  }
}

export const openapiFetchClient = createClient<paths>({
  baseUrl: import.meta.env.VITE_API_BASE_URL
})

openapiFetchClient.use(fetchRequestInterceptor)
openapiFetchClient.use(fetchResponseInterceptor)

if (import.meta.env.MODE === "development") {
  openapiFetchClient.use(debugMiddleware)
}

export type SystemAPIClient = typeof openapiFetchClient
