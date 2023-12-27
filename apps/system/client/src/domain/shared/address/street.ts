import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class AddressStreet extends ValueObject<string> {
  static create(v: string): Result<AddressStreet, Error> {
    if (v === "") {
      return Result.err(new DomainError({ domainKey: domainKeys.AddressStreet, value: v }))
    }
    return Result.ok(new AddressStreet(v))
  }
}
