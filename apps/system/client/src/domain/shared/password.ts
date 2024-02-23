import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class Password extends ValueObject<string> {
  static min = 6
  static max = 128

  // 半角英数字大文字をそれぞれ1種類以上含む6文字以上128文字以下
  static pattern = new RegExp(/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[a-zA-Z0-9!@#$%^&*()-_+=~`{}\[\]|\\:;"'<>,.?/]{6,128}$/)

  static create(v: string): Result<Password, Error> {
    if (!v) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Password,
          value: v,
          message: "Password is required"
        })
      )
    }

    if (v.length < Password.min) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Password,
          value: v,
          message: `Password must contain at least ${Password.min} characters`
        })
      )
    }

    if (v.length > Password.max) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Password,
          value: v,
          message: `Password cannot be longer than ${Password.max} characters`
        })
      )
    }

    if (!Password.pattern.test(v)) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Password,
          value: v,
          message: "Password must contain at least one uppercase letter, one lowercase letter, and one digit."
        })
      )
    }

    return Result.ok(new Password(v))
  }
}
