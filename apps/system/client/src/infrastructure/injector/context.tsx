import { createContext } from "react"
import { di } from "~/infrastructure/injector"
import type { DI } from "~/infrastructure/injector"

export const ContainerContext = createContext<DI>(di)
