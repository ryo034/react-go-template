import type { Me } from "~/domain"
import type { MeStoreType } from "~/store/me/store"
import type { MeUseCaseOutput } from "~/usecase"

export class MePresenter implements MeUseCaseOutput {
  constructor(private readonly store: MeStoreType) {}

  set(v: Me) {
    this.store.getState().set(v)
  }

  clear() {
    this.store.getState().set(null)
  }

  setIsLoading(v: boolean) {
    this.store.getState().setIsLoading(v)
  }
}
