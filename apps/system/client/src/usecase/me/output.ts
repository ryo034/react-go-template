import type { Me } from "~/domain"

export interface MeUseCaseOutput {
  set: (v: Me) => void
  setIsLoading: (v: boolean) => void
  clear: () => void
}
