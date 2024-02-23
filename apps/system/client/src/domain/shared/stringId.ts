import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class StringId extends ValueObject<string> {
  static create(v: string): Result<StringId, Error> {
    if (!v) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.StringId,
          value: v,
          message: "StringId is required"
        })
      )
    }
    return Result.ok(new StringId(v))
  }

  static gen(): StringId {
    return new StringId(Math.floor(Math.random() * 100).toString())
  }

  get asString(): string {
    return this.value
  }
}
