import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class CustomToken extends ValueObject<string> {
  static create(v: string): Result<CustomToken, Error> {
    if (v === "") {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Jwt,
          value: v,
          message: "Jwt must not be empty"
        })
      )
    }
    return Result.ok(new CustomToken(v))
  }
}
