import { Result } from "true-myth"
import { StringId, ValueObject } from "~/domain/shared"

export class InvitationId extends ValueObject<StringId> {
  static create(v: StringId): InvitationId {
    return new InvitationId(v)
  }

  static fromString(v: string): Result<InvitationId, Error> {
    const id = StringId.create(v)
    if (id.isErr) {
      return Result.err(id.error)
    }
    return Result.ok(new InvitationId(id.value))
  }
}
