// import { DI } from "~/infrastructure/injector";
import type { ReactNode } from "react"
import { ContainerContext } from "~/infrastructure/injector/context"

interface Props {
  children: ReactNode
  value: any
}

export const MockContainerContextProvider = ({ children, value }: Props) => (
  <ContainerContext.Provider value={value}>{children}</ContainerContext.Provider>
)
