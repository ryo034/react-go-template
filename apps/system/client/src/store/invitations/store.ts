import { create } from "zustand"
import { immer } from "zustand/middleware/immer"
import { Invitations } from "~/domain"

type State = {
  invitations: Invitations
  invitationsIsLoading: boolean
}

type Actions = {
  setInvitations: (vs: Invitations) => void
  setInvitationsIsLoading: (v: boolean) => void
}

export const invitationStore = create(
  immer<State & Actions>((set) => ({
    invitations: Invitations.empty(),
    invitationsIsLoading: false,
    setInvitations: (vs: Invitations) => set({ invitations: vs }),
    setInvitationsIsLoading: (v: boolean) => set({ invitationsIsLoading: v })
  }))
)

export type InvitationStoreType = typeof invitationStore
