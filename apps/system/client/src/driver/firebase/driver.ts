import { Auth, reload, sendEmailVerification, signInWithEmailAndPassword, signOut } from "firebase/auth"
import { ApiErrorHandler } from "shared-network"
import { Result } from "true-myth"
import { Email, Password } from "~/domain"
import { AuthProviderDriver, AuthProviderUser, UserCredential } from "~/driver/auth"
import { AuthProviderCurrentUserNotFoundError } from "~/infrastructure/error"
import { PromiseResult } from "~/infrastructure/shared"

export class FirebaseDriver implements AuthProviderDriver {
  constructor(private readonly client: Auth, private readonly errorHandler: ApiErrorHandler) {}

  get currentUser(): AuthProviderUser | null {
    this.client.onAuthStateChanged
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

  async signInWithEmailAndPassword(email: Email, password: Password): PromiseResult<UserCredential, Error> {
    try {
      const res = await signInWithEmailAndPassword(this.client, email.value, password.value)
      const credential: UserCredential = {
        user: this.adaptAuthProviderUser(res.user),
        providerId: res.providerId
      }
      return Result.ok(credential)
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async sendEmailVerification(): PromiseResult<null, Error> {
    if (this.client.currentUser === null) {
      return Result.err(new AuthProviderCurrentUserNotFoundError("currentUser is null"))
    }
    try {
      await sendEmailVerification(this.client.currentUser)
      return Result.ok(null)
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
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

  async signUp(): PromiseResult<null, Error> {
    throw new Error("Method not implemented.")
  }

  async signOut(): PromiseResult<null, Error> {
    try {
      await signOut(this.client)
      return Result.ok(null)
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }
}
