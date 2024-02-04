import { Result } from "true-myth"
import { Jwt, Otp } from "~/domain/auth"
import { components } from "~/generated/schema/openapi/systemApi"

export class AuthGatewayAdapter {
  adaptOtp(otp: components["schemas"]["Otp"]): Result<Otp, Error> {
    return Result.ok(new Otp(otp.code))
  }

  adaptJwt(jwt: components["schemas"]["JwtToken"]): Result<Jwt, Error> {
    return Result.ok(new Jwt(jwt.token))
  }
}
