import createClient from "openapi-fetch"
import { paths } from "~/generated/schema/openapi/systemApi"
import { firebaseAuth } from "../firebase"

const fetchRequestInterceptor = async (config: RequestInit | undefined) => {
  if (!config) {
    return config
  }
  if (firebaseAuth.currentUser === null) {
    return config
  }
  const headers = new Headers(config.headers)
  const token = await firebaseAuth.currentUser.getIdToken()
  if (!token) {
    return config
  }
  return { ...config, headers: { ...headers, Authorization: `Bearer ${token}` } }
}

export const openapiFetchClient = createClient<paths>({
  baseUrl: import.meta.env.VITE_API_BASE_URL,
  fetch: async (input, init) => {
    const options = await fetchRequestInterceptor(init)
    return fetch(input, options)
  }
})

export type SystemAPIClient = typeof openapiFetchClient
