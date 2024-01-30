import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class MemberIdNumber extends ValueObject<string> {
  static max = 50
  static create(v: string): Result<MemberIdNumber, Error> {
    const trimmed = v.replace(/[\s　]/g, "")
    if (trimmed.length > MemberIdNumber.max) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.MemberIdNumber,
          value: v,
          message: `MemberIdNumber must be less than ${MemberIdNumber.max} characters: ${v}`
        })
      )
    }
    return Result.ok(new MemberIdNumber(trimmed))
  }
}
