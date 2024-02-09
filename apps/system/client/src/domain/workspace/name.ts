import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class WorkspaceName extends ValueObject<string> {
  static max = 255
  static create(v: string): Result<WorkspaceName, Error> {
    const trimmedName = v.replace(/[\s　]/g, "")
    if (trimmedName.length > WorkspaceName.max) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.AccountName,
          value: v,
          message: `Workspace name must be less than ${WorkspaceName.max} characters: ${v}`
        })
      )
    }
    return Result.ok(new WorkspaceName(trimmedName))
  }
}
