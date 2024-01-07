import createClient from "openapi-fetch"
import { paths } from "~/generated/schema/openapi/systemApi"

export const openapiFetchClient = createClient<paths>({ baseUrl: "http://localhost:19004/api/v1" })

export type SystemAPIClient = typeof openapiFetchClient
