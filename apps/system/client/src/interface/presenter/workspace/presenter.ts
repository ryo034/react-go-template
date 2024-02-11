import { Members } from "~/domain"
import { WorkspaceStoreType } from "~/store/workspace/store"
import { WorkspaceUseCaseOutput } from "~/usecase"

export class WorkspacePresenter implements WorkspaceUseCaseOutput {
  constructor(private readonly store: WorkspaceStoreType) {}

  setMembers(vs: Members) {
    this.store.getState().setMembers(vs)
  }

  setMembersIsLoading(v: boolean) {
    this.store.getState().setMembersIsLoading(v)
  }

  clear() {
    this.store.getState().setMembers(Members.empty())
  }
}
