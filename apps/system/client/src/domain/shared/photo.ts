import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class Photo extends ValueObject<URL> {
  static fromString(v: string): Result<Photo, Error> {
    if (v === "") {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Photo,
          value: v,
          message: "Photo is required"
        })
      )
    }

    try {
      return Result.ok(new Photo(new URL(v)))
    } catch (e) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Photo,
          value: v,
          message: "Photo is invalid"
        })
      )
    }
  }
}
