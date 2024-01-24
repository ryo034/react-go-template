import { APIRequestContext, APIResponse } from "@playwright/test"

export const systemApiHost = process.env.SYSTEM_API_HOST || "http://localhost:19004"

export const grpcRestAuthenticatedRequest = async (
  request: APIRequestContext,
  path: string,
  authToken: string,
  data: Object = {}
): Promise<APIResponse> => {
  return await request.post(`${systemApiHost}${path}`, {
    headers: { "content-type": "application/json", Authorization: `Bearer ${authToken}` },
    data
  })
}
