import { ThemeDriver } from "~/driver"
import { ThemeType } from "~/store"
import { ThemeUseCaseOutput, ToggleInput } from "~/usecase"

export interface ThemeUseCase {
  toggle(theme: ThemeType): void
  init(): void
}

export class ThemeInteractor implements ThemeUseCase {
  constructor(private readonly driver: ThemeDriver, private readonly presenter: ThemeUseCaseOutput) {}

  toggle(theme: ThemeType): void {
    this.driver.set(theme)
    this.presenter.set(theme)
  }

  init(): void {
    const theme = this.driver.get()
    if (theme != null) {
      this.presenter.set(theme)
    }
  }
}
