export const origin = process.env.ORIGIN ?? "http://localhost"

export const defaultPostHeaders = {
  "content-type": "application/json",
  origin
}
