import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class MemberDisplayName extends ValueObject<string> {
  static max = 50
  static create(v: string): Result<MemberDisplayName, Error> {
    const trimmed = v.replace(/[\s　]/g, "")
    if (trimmed.length > MemberDisplayName.max) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.MemberDisplayName,
          value: v,
          message: `MemberDisplayName must be less than ${MemberDisplayName.max} characters: ${v}`
        })
      )
    }
    return Result.ok(new MemberDisplayName(trimmed))
  }

  get firstTwoCharacters(): string {
    return this.value.slice(0, 2).toUpperCase() ?? ""
  }
}
