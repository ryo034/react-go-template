import { Result } from "true-myth"
import { CustomToken } from "~/domain"
import { components } from "~/generated/schema/openapi/systemApi"

export class AuthGatewayAdapter {
  adaptJwt(customToken: components["schemas"]["JwtToken"]): Result<CustomToken, Error> {
    return Result.ok(new CustomToken(customToken.token))
  }
}
