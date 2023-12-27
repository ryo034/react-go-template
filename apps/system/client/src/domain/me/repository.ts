import { Me } from "~/domain/me"
import { Email, Password } from "~/domain/shared"
import { PromiseResult } from "~/infrastructure/shared"

export interface MeRepository {
  login(): PromiseResult<Me, Error>
  sendEmailVerification(): PromiseResult<null, Error>
  signInWithEmailAndPassword(email: Email, password: Password): PromiseResult<null, Error>
  signOut(): PromiseResult<null, Error>
  verifyEmail(): PromiseResult<null, Error>
  reloadAuth(): PromiseResult<Me, Error>
  find(): PromiseResult<Me, Error>
}
