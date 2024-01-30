import { Result } from "true-myth"
import { StringId, ValueObject } from "~/domain/shared"

export class MemberId extends ValueObject<StringId> {
  static create(v: StringId): MemberId {
    return new MemberId(v)
  }

  static fromString(v: string): Result<MemberId, Error> {
    const id = StringId.create(v)
    if (id.isErr) {
      return Result.err(id.error)
    }
    return Result.ok(new MemberId(id.value))
  }
}
