import { ThemeInteractor, ToggleInput } from "~/usecase"

export class ThemeController {
  constructor(private readonly useCase: ThemeInteractor) {}

  toggle(isDark: boolean): null {
    const input: ToggleInput = { isDark }
    return this.useCase.toggle(input)
  }

  init(): void {
    this.useCase.get()
  }
}
