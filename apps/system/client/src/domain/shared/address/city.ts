import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class AddressCity extends ValueObject<string> {
  static create(v: string): Result<AddressCity, Error> {
    if (v === "") {
      return Result.err(new DomainError({ domainKey: domainKeys.AddressCity, value: v }))
    }
    return Result.ok(new AddressCity(v))
  }
}
