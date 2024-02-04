import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class Otp extends ValueObject<string> {
  static pattern = /^[0-9]{6}$/
  static create(v: string): Result<Otp, Error> {
    const trimmed = v.replace(/[\s　]/g, "")
    if (!Otp.pattern.test(trimmed)) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Otp,
          value: v,
          message: `Otp must be 6 digits: ${v}`
        })
      )
    }
    return Result.ok(new Otp(trimmed))
  }
}
