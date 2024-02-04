import { firebaseAuth } from "../firebase"

const { fetch: originalFetch } = window

window.fetch = async (...args) => {
  const [resource, config] = args

  const options = await fetchRequestInterceptor(config)
  const response = await originalFetch(resource, options)
  return response
}

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
