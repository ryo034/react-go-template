import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class PhotoPath extends ValueObject<string> {
  static create(v: string): Result<PhotoPath, Error> {
    if (v === "") {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.PhotoPath,
          value: v,
          message: `Photo path must not be empty: ${v}`
        })
      )
    }
    // /からのパスは許可しない
    if (v.startsWith("/")) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.PhotoPath,
          value: v,
          message: `Photo path must not start with /: ${v}`
        })
      )
    }
    return Result.ok(new PhotoPath(v))
  }
}
