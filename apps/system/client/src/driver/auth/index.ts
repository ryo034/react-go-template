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
  signUp(): PromiseResult<null, Error>
  signOut(): PromiseResult<null, Error>
}
