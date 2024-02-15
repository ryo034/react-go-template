import createClient from "openapi-fetch"
import { paths } from "~/generated/schema/openapi/systemApi"
import { firebaseAuth } from "~/infrastructure/firebase"

const fetchRequestInterceptor = async (config: RequestInit | undefined) => {
  let newConfig = { ...config }

  newConfig = config === undefined ? {} : { ...config }

  if (firebaseAuth.currentUser === null) {
    return newConfig
  }
  const token = await firebaseAuth.currentUser.getIdToken()
  if (!token) {
    return config
  }
  return {
    ...config,
    headers: { ...new Headers(newConfig.headers), Authorization: `Bearer ${token}`, "Content-Type": "application/json" }
  }
}

export const openapiFetchClient = createClient<paths>({
  baseUrl: import.meta.env.VITE_API_BASE_URL,
  fetch: async (input, init) => {
    const options = await fetchRequestInterceptor(init)
    return fetch(input, { ...options })
  }
})

export type SystemAPIClient = typeof openapiFetchClient
