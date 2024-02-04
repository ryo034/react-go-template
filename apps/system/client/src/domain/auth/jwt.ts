import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class Jwt extends ValueObject<string> {
  static create(v: string): Result<Jwt, Error> {
    if (v === "") {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Jwt,
          value: v,
          message: "Jwt must not be empty"
        })
      )
    }
    return Result.ok(new Jwt(v))
  }
}
