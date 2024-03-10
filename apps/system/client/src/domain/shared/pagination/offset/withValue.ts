import { Entity } from "~/domain/shared"
import type { OffsetPagination } from "~/domain/shared/pagination"

interface Props<T> {
  values: T
  pageInfo: OffsetPagination
}

export class WithSearchParam<T> extends Entity<Props<T>> {
  static create<T>(v: Props<T>): WithSearchParam<T> {
    return new WithSearchParam(v)
  }

  get values(): T {
    return this.value.values
  }

  get pageInfo(): OffsetPagination {
    return this.value.pageInfo
  }
}
