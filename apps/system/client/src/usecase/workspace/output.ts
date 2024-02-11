import { Me, Members } from "~/domain"

export interface WorkspaceUseCaseOutput {
  setMembers: (vs: Members) => void
  setMembersIsLoading: (v: boolean) => void
  clear: () => void
}
