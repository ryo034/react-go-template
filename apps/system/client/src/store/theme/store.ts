import { create } from "zustand"
import { immer } from "zustand/middleware/immer"

export type ThemeType = "light" | "dark"

type State = {
  theme: ThemeType
}

type Actions = {
  set: (v: ThemeType) => void
}

export const themeStore = create(
  immer<State & Actions>((set) => ({
    theme: "light" as ThemeType,
    set: (v: ThemeType) => set({ theme: v })
  }))
)

export type ThemeStoreType = typeof themeStore
