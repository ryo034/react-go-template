import { Members, Workspace, WorkspaceSubdomain } from "~/domain"
import { PromiseResult } from "~/infrastructure/shared"

export interface WorkspaceCreateInput {
  subdomain: WorkspaceSubdomain
}

export interface WorkspaceRepository {
  create(i: WorkspaceCreateInput): PromiseResult<Workspace, Error>
  findAllMembers(): PromiseResult<Members, Error>
}
