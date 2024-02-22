import { create } from "zustand"
import { immer } from "zustand/middleware/immer"
import { ReceivedInvitation } from "~/domain"

type State = {
  invitation: ReceivedInvitation | null
  invitationIsLoading: boolean
  isInvitationProcessing: boolean
}

type Actions = {
  setReceivedInvitation: (v: ReceivedInvitation) => void
  setReceivedInvitationIsLoading: (v: boolean) => void
  setIsInvitationProcessing: (v: boolean) => void
}

export const receivedInvitationStore = create(
  immer<State & Actions>((set) => ({
    invitation: null,
    invitationIsLoading: false,
    isInvitationProcessing: false,
    setReceivedInvitation: (v: ReceivedInvitation) => set({ invitation: v }),
    setReceivedInvitationIsLoading: (v: boolean) => set({ invitationIsLoading: v }),
    setIsInvitationProcessing: (v: boolean) => set({ isInvitationProcessing: v })
  }))
)

export type ReceivedInvitationStoreType = typeof receivedInvitationStore
