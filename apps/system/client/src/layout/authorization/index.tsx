import { ReactNode, useContext } from "react"
import { Outlet } from "react-router-dom"
import { MemberRole } from "~/domain"
import { ContainerContext } from "~/infrastructure/injector/context"

interface Props {
  roles: MemberRole[]
  fallback: ReactNode
}

export const AuthorizationLayout = ({ roles, fallback }: Props) => {
  const { store } = useContext(ContainerContext)
  const me = store.me((state) => state.me)
  const meIsLoading = store.me((state) => state.isLoading)

  if (me === null || meIsLoading) return <></>

  return roles.some((role) => role === me.member?.role) ? <Outlet /> : fallback
}
