import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class MemberBio extends ValueObject<string> {
  static max = 2000
  static create(v: string): Result<MemberBio, Error> {
    if (v.length > MemberBio.max) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.MemberBio,
          value: v,
          message: `MemberBio must be less than ${MemberBio.max} characters: ${v}`
        })
      )
    }
    return Result.ok(new MemberBio(v))
  }
}
