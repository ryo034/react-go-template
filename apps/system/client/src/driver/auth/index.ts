import { Email, Password } from "~/domain"
import { PromiseResult } from "~/infrastructure/shared"

export interface AuthProviderUser {
  readonly uid: string
  readonly email: string | null
  readonly emailVerified: boolean
  readonly displayName: string | null
}

export interface UserCredential {
  readonly user: AuthProviderUser
  readonly providerId: string | null
}

export interface AuthProviderDriver {
  readonly currentUser: AuthProviderUser | null
  reload(): PromiseResult<null, Error>
  sendEmailVerification(): PromiseResult<null, Error>
  signInWithEmailAndPassword(email: Email, password: Password): PromiseResult<UserCredential, Error>
  signOut(): PromiseResult<null, Error>
}
