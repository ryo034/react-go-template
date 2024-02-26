import { Auth, reload, signInWithCustomToken, signOut } from "firebase/auth"
import { ApiErrorHandler } from "shared-network"
import { Result } from "true-myth"
import { CustomToken } from "~/domain/auth"
import { AuthProviderCurrentUserNotFoundError } from "~/infrastructure/error"
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
  signOut(): PromiseResult<null, Error>
  signInWithCustomToken(customToken: CustomToken): PromiseResult<UserCredential, Error>
}

export class FirebaseDriver implements AuthProviderDriver {
  constructor(private readonly client: Auth, private readonly errorHandler: ApiErrorHandler) {}

  get currentUser(): AuthProviderUser | null {
    if (this.client.currentUser === null) {
      return null
    }
    return this.adaptAuthProviderUser(this.client.currentUser)
  }

  private adaptAuthProviderUser(user: AuthProviderUser): AuthProviderUser {
    return {
      uid: user.uid,
      email: user.email,
      displayName: user.displayName,
      emailVerified: user.emailVerified
    }
  }

  async reload(): PromiseResult<null, Error> {
    try {
      if (this.client.currentUser === null) {
        return Result.err(new AuthProviderCurrentUserNotFoundError("currentUser is null"))
      }
      await reload(this.client.currentUser)
      return Result.ok(null)
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async signOut(): PromiseResult<null, Error> {
    try {
      await signOut(this.client)
      return Result.ok(null)
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async signInWithCustomToken(customToken: CustomToken): PromiseResult<UserCredential, Error> {
    try {
      const res = await signInWithCustomToken(this.client, customToken.value)
      const credential: UserCredential = {
        user: this.adaptAuthProviderUser(res.user),
        providerId: res.providerId
      }
      return Result.ok(credential)
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }
}
