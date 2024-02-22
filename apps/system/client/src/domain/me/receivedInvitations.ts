import { Entities, ReceivedInvitation } from "~/domain"

export class ReceivedInvitations extends Entities<ReceivedInvitation> {
  static create(vs: Array<ReceivedInvitation>): ReceivedInvitations {
    return new ReceivedInvitations(vs)
  }
  static empty(): ReceivedInvitations {
    return new ReceivedInvitations([])
  }
}
