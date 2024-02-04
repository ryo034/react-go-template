import { Me } from "~/domain/me"
import { PromiseResult } from "~/infrastructure/shared"

export interface MeRepository {
  signOut(): PromiseResult<null, Error>
  find(): PromiseResult<Me, Error>
}
