import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class AccountFullName extends ValueObject<string> {
  // 漢字/ひらがな/カタカナ/半角文字
  static pattern = /^[ぁ-んァ-ン一-龥a-zA-Z ]+$/
  static max = 50
  static create(v: string): Result<AccountFullName, Error> {
    const trimmed = v.trim()
    if (trimmed.length > AccountFullName.max) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.AccountFullName,
          value: v,
          message: `Account name must be less than ${AccountFullName.max} characters: ${v}`
        })
      )
    }
    if (!AccountFullName.pattern.test(trimmed)) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.AccountFullName,
          value: v,
          message: `Account name must be in Japanese or English: ${v}`
        })
      )
    }
    return Result.ok(new AccountFullName(trimmed))
  }
}
