import type { ReceivedInvitation } from "~/domain"
import { Entities } from "~/domain/shared"

export class ReceivedInvitations extends Entities<ReceivedInvitation> {
  static create(vs: Array<ReceivedInvitation>): ReceivedInvitations {
    return new ReceivedInvitations(vs)
  }
  static empty(): ReceivedInvitations {
    return new ReceivedInvitations([])
  }
}
