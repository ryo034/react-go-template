import { Entities, Invitee } from "~/domain"

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
