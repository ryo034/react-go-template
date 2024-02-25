import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class AccountFullName extends ValueObject<string> {
  // 漢字/ひらがな/カタカナ/半角文字
  static pattern = /^[ぁ-んァ-ン一-龥a-zA-Z\u3000 ]+$/
  static max = 50
  static create(v: string): Result<AccountFullName, Error> {
    const formatted = AccountFullName.format(v)
    if (formatted.length > AccountFullName.max) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.AccountFullName,
          value: v,
          message: `Account name must be less than ${AccountFullName.max} characters: ${v}`
        })
      )
    }
    if (!AccountFullName.pattern.test(formatted)) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.AccountFullName,
          value: v,
          message: `Account name must be in Japanese or English: ${v}`
        })
      )
    }
    return Result.ok(new AccountFullName(formatted))
  }

  static format(v: string): string {
    return v.trim().replace(/　/g, " ")
  }
}
