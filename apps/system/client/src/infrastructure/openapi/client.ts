import createClient from "openapi-fetch"
import { paths } from "~/generated/schema/openapi/systemApi"

export const openapiFetchClient = createClient<paths>({ baseUrl: import.meta.env.VITE_API_BASE_URL })

export type SystemAPIClient = typeof openapiFetchClient
