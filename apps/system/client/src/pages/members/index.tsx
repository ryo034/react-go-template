import { useContext, useLayoutEffect, useRef } from "react"
import { MemberCard } from "~/components/member/card"
import { MemberCardListLoading } from "~/components/member/cardListLoading"
import { ContainerContext } from "~/infrastructure/injector/context"

export const membersPageRoute = "/members"

export const MembersPage = () => {
  const { controller, store } = useContext(ContainerContext)

  const me = store.me((s) => s.me)
  const members = store.workspace((s) => s.members)
  const membersIsLoading = store.workspace((s) => s.membersIsLoading)
  const membersRef = useRef(members)

  useLayoutEffect(() => {
    store.workspace.subscribe((v) => {
      membersRef.current = v.members
    })

    const fetchMembers = async () => {
      if (me === null || me.workspace === undefined) {
        return null
      }
      await controller.workspace.findAllMembers()
    }
    fetchMembers()
  }, [])

  if (me === null) {
    return <></>
  }

  return (
    <div className="px-8">
      <header className="py-8">
        <h1 className="text-2xl font-bold tracking-tight">Members</h1>
      </header>
      <div className="grid grid-cols-4 gap-8">
        {membersIsLoading && <MemberCardListLoading count={10} />}
        {!membersIsLoading &&
          members.values.map((m) => {
            return <MemberCard key={m.profile.id.value.asString} member={m} />
          })}
      </div>
    </div>
  )
}
