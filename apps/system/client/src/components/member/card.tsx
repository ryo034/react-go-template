import { Avatar, AvatarFallback, AvatarImage } from "shared-ui"
import { Member } from "~/domain/workspace/member"
import { AccountAvatar } from "../account/avatar"

interface Props {
  member: Member
}

export const MemberCard = ({ member }: Props) => {
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
        <span className="text-sm text-muted-foreground">Co-Founder, CEO</span>
      </div>
    </div>
  )
}
