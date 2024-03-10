import type { ThemeDriver } from "~/driver"
import type { ThemeType } from "~/store"
import { type ThemeUseCaseOutput, ToggleInput } from "~/usecase"

export interface ThemeUseCase {
  toggle(theme: ThemeType): void
  init(): void
}

export class ThemeInteractor implements ThemeUseCase {
  constructor(
    private readonly driver: ThemeDriver,
    private readonly presenter: ThemeUseCaseOutput
  ) {}

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
