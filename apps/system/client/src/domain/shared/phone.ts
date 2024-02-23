import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class Phone extends ValueObject<string> {
  static pattern = new RegExp(/^0[789]0\d{8}$/)

  static create(v: string): Result<Phone, Error> {
    if (!Phone.pattern.test(v)) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Phone,
          value: v,
          message: "invalid phone number"
        })
      )
    }
    return Result.ok(new Phone(v))
  }

  get toInternationalNumberString(): string {
    return `+81${this.value.slice(1)}`
  }
}
