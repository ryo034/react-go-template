import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class WorkspaceSubdomain extends ValueObject<string> {
  static max = 63
  static min = 1
  static pattern = /^[a-z0-9]+(-[a-z0-9]+)*$/

  private static isValidSubdomain(subdomain: string): boolean {
    return (
      WorkspaceSubdomain.pattern.test(subdomain) &&
      subdomain.length >= WorkspaceSubdomain.min &&
      subdomain.length <= WorkspaceSubdomain.max
    )
  }

  static create(v: string): Result<WorkspaceSubdomain, Error> {
    if (!WorkspaceSubdomain.isValidSubdomain(v)) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.WorkspaceSubdomain,
          value: v,
          message: `Invalid workspace domain: ${v}`
        })
      )
    }
    return Result.ok(new WorkspaceSubdomain(v))
  }

  toString(): string {
    return this.value
  }
}
