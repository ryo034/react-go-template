import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class AddressBuilding extends ValueObject<string> {
  static create(v: string): Result<AddressBuilding, Error> {
    if (!v) {
      return Result.err(new DomainError({ domainKey: domainKeys.AddressBuilding, value: v }))
    }
    return Result.ok(new AddressBuilding(v))
  }
}
