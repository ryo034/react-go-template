export const origin = process.env.ORIGIN ?? "http://localhost"

export const defaultPostHeaders = {
  "content-type": "application/json",
  origin
}

export const headers = (token: string) => {
  if (!token) return defaultPostHeaders
  return {
    ...defaultPostHeaders,
    Authorization: `Bearer ${token}`
  }
}
