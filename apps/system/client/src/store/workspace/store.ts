import { create } from "zustand"
import { immer } from "zustand/middleware/immer"
import { Members } from "~/domain"

type State = {
  members: Members
  membersIsLoading: boolean
}

type Actions = {
  setMembers: (v: Members) => void
  setMembersIsLoading: (v: boolean) => void
}

export const workspaceStore = create(
  immer<State & Actions>((set) => ({
    members: Members.empty(),
    membersIsLoading: false,
    setMembers: (v: Members) => set({ members: v }),
    setMembersIsLoading: (v: boolean) => set({ membersIsLoading: v })
  }))
)

export type WorkspaceStoreType = typeof workspaceStore
