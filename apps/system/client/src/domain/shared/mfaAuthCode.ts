import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export const regexErrMsg = "Password must contain at least one uppercase letter, one lowercase letter, and one digit."

export class MFAAuthCode extends ValueObject<string> {
  static pattern = new RegExp(/[0-9]{6}/)

  static create(v: string): Result<MFAAuthCode, Error> {
    if (!v) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.MFAAuthCode,
          value: v,
          message: "MFAAuthCode is required"
        })
      )
    }
    if (!MFAAuthCode.pattern.test(v.toString())) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.MFAAuthCode,
          value: v,
          message: regexErrMsg
        })
      )
    }
    return Result.ok(new MFAAuthCode(v))
  }
}
