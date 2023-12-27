import { create } from "zustand"
import { immer } from "zustand/middleware/immer"

type State = {
  isDark: boolean
}

type Actions = {
  set: (v: boolean) => void
}

export const themeStore = create(
  immer<State & Actions>((set) => ({
    isDark: false,
    set: (v: boolean) => set({ isDark: v })
  }))
)

export type ThemeStoreType = typeof themeStore
