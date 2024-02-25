import { ThemeStoreType, ThemeType } from "~/store/theme/store"
import { ThemeUseCaseOutput } from "~/usecase"

export class ThemePresenter implements ThemeUseCaseOutput {
  constructor(private readonly themeStore: ThemeStoreType) {}

  set(v: ThemeType) {
    this.themeStore.getState().set(v)
    document.documentElement.classList.remove("light", "dark")
    document.documentElement.classList.add(v)
  }
}
