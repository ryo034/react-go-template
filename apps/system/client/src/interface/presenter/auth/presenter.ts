import { Email } from "~/domain"
import { AuthStoreType } from "~/store"
import { AuthUseCaseOutput } from "~/usecase/auth"

export class AuthPresenter implements AuthUseCaseOutput {
  constructor(private readonly store: AuthStoreType) {}

  setEmail(v: Email) {
    this.store.getState().setEmail(v)
  }

  clear() {
    this.store.getState().set(null)
  }

  setIsLoading(v: boolean) {
    this.store.getState().setIsLoading(v)
  }
}
