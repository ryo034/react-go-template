import { Result } from "true-myth"
import { AuthProvider, AuthProviders, CustomToken } from "~/domain"
import { components } from "~/generated/schema/openapi/systemApi"

export class AuthGatewayAdapter {
  adaptJwt(customToken: components["schemas"]["JwtToken"]): Result<CustomToken, Error> {
    return Result.ok(new CustomToken(customToken.token))
  }

  adaptAuthProvider(authProvider: components["schemas"]["AuthProvider"]): Result<AuthProvider, Error> {
    switch (authProvider) {
      case "email":
        return Result.ok("email")
      case "google":
        return Result.ok("google")
      default:
        return Result.err(new Error(`Unknown auth provider: ${authProvider}`))
    }
  }

  adaptAllAuthProvider(vs: Array<components["schemas"]["AuthProvider"]> | undefined): Result<AuthProviders, Error> {
    if (vs === undefined) {
      return Result.ok(AuthProviders.create([]))
    }
    const result: Array<AuthProvider> = []
    for (const v of vs) {
      const r = this.adaptAuthProvider(v)
      if (r.isErr) {
        return Result.err(r.error)
      }
      result.push(r.value)
    }
    return Result.ok(AuthProviders.create(result))
  }
}
