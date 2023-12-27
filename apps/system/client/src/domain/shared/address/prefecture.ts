import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class AddressPrefecture extends ValueObject<string> {
  static create(v: string): Result<AddressPrefecture, Error> {
    if (v === "") {
      return Result.err(new DomainError({ domainKey: domainKeys.AddressPrefecture, value: v }))
    }
    return Result.ok(new AddressPrefecture(v))
  }
}
