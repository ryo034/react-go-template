import type { ReactNode } from "react"
import { di } from "~/infrastructure/injector"
import { ContainerContext } from "~/infrastructure/injector/context"

export const ContainerProvider = ({ children }: { children: ReactNode }) => {
  return <ContainerContext.Provider value={di}>{children}</ContainerContext.Provider>
}
