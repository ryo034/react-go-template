import { ThemeType } from "~/store"

export interface ThemeUseCaseOutput {
  set: (v: ThemeType) => void
}
