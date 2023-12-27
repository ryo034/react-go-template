import { Entity } from "~/domain/shared"

interface Props {
  limit: number
  page: number
  total: number
}

export class OffsetPagination extends Entity<Props> {
  static defaultLimit = 3

  static create(v: Props): OffsetPagination {
    return new OffsetPagination(v)
  }

  static default(): OffsetPagination {
    return new OffsetPagination({
      limit: OffsetPagination.defaultLimit,
      page: 1,
      total: 0
    })
  }

  get isDefault(): boolean {
    return this.value.limit === OffsetPagination.defaultLimit && this.value.page === 1 && this.value.total === 0
  }

  get limit(): number {
    return this.value.limit
  }

  get page(): number {
    return this.value.page
  }

  get total(): number {
    return this.value.total
  }
}
