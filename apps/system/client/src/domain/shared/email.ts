import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "."

export class Email extends ValueObject<string> {
  static max = 320
  static pattern = new RegExp(/^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$/)

  static create(v: string): Result<Email, Error> {
    if (!v) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Email,
          value: v,
          message: "Email is required"
        })
      )
    }

    if (v.length > Email.max) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Email,
          value: v,
          message: `Email cannot be longer than ${Email.max} characters`
        })
      )
    }

    if (!Email.pattern.test(v)) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Email,
          value: v,
          message: "Email is invalid"
        })
      )
    }

    return Result.ok(new Email(v))
  }
}
