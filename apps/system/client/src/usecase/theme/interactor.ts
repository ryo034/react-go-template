import { ThemeDriver } from "~/driver"
import { ThemeUseCaseOutput, ToggleInput } from "~/usecase"

export interface ThemeUseCase {
  toggle(i: ToggleInput): null
  get(): void
}

export class ThemeInteractor implements ThemeUseCase {
  constructor(private readonly driver: ThemeDriver, private readonly presenter: ThemeUseCaseOutput) {}

  toggle(i: ToggleInput): null {
    this.driver.set(i.isDark)
    this.presenter.set(i.isDark)
    return null
  }

  get(): void {
    const isDark = this.driver.get()
    this.presenter.set(isDark)
  }
}
