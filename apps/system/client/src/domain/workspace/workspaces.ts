import { Workspace } from "~/domain"
import { Entities } from "~/domain/shared"

export class Workspaces extends Entities<Workspace> {
  static create(vs: Array<Workspace>): Workspaces {
    return new Workspaces(vs)
  }
  static empty(): Workspaces {
    return new Workspaces([])
  }
}
