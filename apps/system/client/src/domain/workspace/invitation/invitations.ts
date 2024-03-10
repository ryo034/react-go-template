import type { Invitation } from "~/domain"
import { Entities } from "~/domain/shared"

export class Invitations extends Entities<Invitation> {
  static create(vs: Array<Invitation>): Invitations {
    return new Invitations(vs)
  }
  static empty(): Invitations {
    return new Invitations([])
  }
}
