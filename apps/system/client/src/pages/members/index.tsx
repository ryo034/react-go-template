import { useContext, useLayoutEffect, useRef } from "react"
import { Separator } from "shared-ui"
import { MemberCard } from "~/components/member/card"
import { MemberCardListLoading } from "~/components/member/cardListLoading"
import { MembersSettingDialog } from "~/components/member/membersSettingDialog"
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
    <>
      <header>
        <div className="flex space-x-4 items-center">
          <h2 className="text-2xl font-bold tracking-tight" data-testid="pageTitle">
            Members
          </h2>
          <MembersSettingDialog />
        </div>
        {/* <InviteMembersDialog /> */}
        <Separator className="my-6" />
      </header>
      <div className="grid grid-cols-4 gap-8">
        {membersIsLoading && <MemberCardListLoading count={10} />}
        {!membersIsLoading && membersRef.current.values.map((m) => <MemberCard key={m.id.value.asString} member={m} />)}
      </div>
    </>
  )
}
