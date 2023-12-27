import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class AddressCountry extends ValueObject<string> {
  static create(v: string): Result<AddressCountry, Error> {
    if (v === "") {
      return Result.err(new DomainError({ domainKey: domainKeys.AddressCountry, value: v }))
    }
    return Result.ok(new AddressCountry(v))
  }

  static default(): AddressCountry {
    return new AddressCountry("日本")
  }
}
