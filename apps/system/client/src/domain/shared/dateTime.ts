import { Result } from "true-myth"
import { DomainError, ValueObject, domainKeys } from "~/domain/shared"

export class AppDateTime extends ValueObject<Date> {
  static create(v: Date): AppDateTime {
    return new AppDateTime(v)
  }

  static now(): AppDateTime {
    return new AppDateTime(new Date(Date.now()))
  }

  static fromString(v: string): Result<AppDateTime, Error> {
    if (v === "") {
      return Result.err(new DomainError({ domainKey: domainKeys.AppDateTime, value: v }))
    }
    const date = new Date(v)
    if (Number.isNaN(date.getTime())) {
      return Result.err(new DomainError({ domainKey: domainKeys.AppDateTime, value: v }))
    }
    return Result.ok(new AppDateTime(date))
  }

  get localeDateTimeString(): string {
    return this.value.toLocaleString("ja-JP")
  }

  get hyphenString(): string | undefined {
    return this.value.toISOString().split("T")[0]
  }
}
