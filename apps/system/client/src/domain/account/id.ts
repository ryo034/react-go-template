import { Result } from "true-myth"
import { StringId, ValueObject } from "~/domain/shared"

export class AccountId extends ValueObject<StringId> {
  static create(v: StringId): AccountId {
    return new AccountId(v)
  }

  static gen(): AccountId {
    return new AccountId(StringId.gen())
  }

  static fromString(v: string): Result<AccountId, Error> {
    const id = StringId.create(v)
    if (id.isErr) {
      return Result.err(id.error)
    }
    return Result.ok(new AccountId(id.value))
  }
}
