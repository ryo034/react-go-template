import type { Invitee } from "~/domain"
import { Entities } from "~/domain/shared"

export class Invitees extends Entities<Invitee> {
  static create(vs: Array<Invitee>): Invitees {
    return new Invitees(vs)
  }
  static empty(): Invitees {
    return new Invitees([])
  }

  add(invitee: Invitee): Invitees {
    return Invitees.create([...this.props, invitee])
  }
}
