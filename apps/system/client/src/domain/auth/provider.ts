import { Entities } from "~/domain/shared"

export const authProviderList = {
  email: "email",
  google: "google"
} as const

export type AuthProvider = (typeof authProviderList)[keyof typeof authProviderList]

export class AuthProviders extends Entities<AuthProvider> {
  static create(vs: Array<AuthProvider>): AuthProviders {
    return new AuthProviders(vs)
  }

  get hasEmail(): boolean {
    return this.values.includes(authProviderList.email)
  }

  get hasGoogle(): boolean {
    return this.values.includes(authProviderList.google)
  }
}
