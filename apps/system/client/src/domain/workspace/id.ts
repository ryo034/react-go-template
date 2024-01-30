import { Result } from "true-myth"
import { StringId, ValueObject } from "~/domain/shared"

export class WorkspaceId extends ValueObject<StringId> {
  static create(v: StringId): WorkspaceId {
    return new WorkspaceId(v)
  }

  static fromString(v: string): Result<WorkspaceId, Error> {
    const id = StringId.create(v)
    if (id.isErr) {
      return Result.err(id.error)
    }
    return Result.ok(new WorkspaceId(id.value))
  }
}
