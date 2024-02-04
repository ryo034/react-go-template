import { create } from "zustand"
import { immer } from "zustand/middleware/immer"
import { Email } from "~/domain"
import { Otp } from "~/domain/auth"

type State = {
  otp: Otp | null
  email: Email | null
  isLoading: boolean
}

type Actions = {
  set: (v: Otp | null) => void
  setEmail: (v: Email | null) => void
  setIsLoading: (v: boolean) => void
}

export const authStore = create(
  immer<State & Actions>((set) => ({
    otp: null,
    email: null,
    isLoading: false,
    set: (v: Otp | null) => set({ otp: v }),
    setEmail: (v: Email | null) => set({ email: v }),
    setIsLoading: (v: boolean) => set({ isLoading: v })
  }))
)

export type AuthStoreType = typeof authStore
