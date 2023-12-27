import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class AddressZipCode extends ValueObject<string> {
  static pattern = new RegExp(/^[0-9]{3}-?[0-9]{4}$/)
  static strictPattern = new RegExp(/^\d{7}$/)

  static create(v: string): Result<AddressZipCode, Error> {
    if (v === "") {
      return Result.err(new DomainError({ domainKey: domainKeys.AddressZipCode, value: v }))
    }
    if (AddressZipCode.strictPattern.test(v)) {
      return Result.ok(new AddressZipCode(v))
    }
    if (AddressZipCode.pattern.test(v)) {
      return Result.ok(new AddressZipCode(v.replace("-", "")))
    }
    return Result.err(new DomainError({ domainKey: domainKeys.AddressZipCode, value: v }))
  }
}
