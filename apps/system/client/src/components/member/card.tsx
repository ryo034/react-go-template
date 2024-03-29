import type { Member } from "~/domain/workspace/member"
import { useRole } from "~/infrastructure/hooks/role"
import { AccountAvatar } from "../account/avatar"

interface Props {
  member: Member
}

export const MemberCard = ({ member }: Props) => {
  const { translateRole } = useRole()
  return (
    <div className="flex flex-wrap ju text-left">
      <AccountAvatar
        size="xl"
        url={member.user.photo?.photoURL || ""}
        alt={"aa"}
        fallbackString={member.profile.displayName?.firstTwoCharacters || ""}
      />
      <div className="mt-4 w-full">
        <p className="font-bold">{member.profile.displayName?.value}</p>
        <span className="text-sm text-muted-foreground">{translateRole(member.role)}</span>
      </div>
    </div>
  )
}
