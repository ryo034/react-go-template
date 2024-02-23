import { ThemeType } from "~/store"
import { ThemeInteractor, ToggleInput } from "~/usecase"

export class ThemeController {
  constructor(private readonly useCase: ThemeInteractor) {}

  toggle(theme: ThemeType): void {
    this.useCase.toggle(theme)
  }

  init(): void {
    this.useCase.init()
  }
}
