import type { Member } from "~/domain"
import { Entities } from "~/domain/shared"

export class Members extends Entities<Member> {
  static create(vs: Array<Member>): Members {
    return new Members(vs)
  }
  static empty(): Members {
    return new Members([])
  }
}
