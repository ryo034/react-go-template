import { Me } from "~/domain/me"
import { PromiseResult } from "~/infrastructure/shared"

export interface MeRepository {
  login(): PromiseResult<Me, Error>
  signUp(): PromiseResult<null, Error>
  signOut(): PromiseResult<null, Error>
  find(): PromiseResult<Me, Error>
}
