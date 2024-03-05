import { Avatar, AvatarFallback, AvatarImage } from "shared-ui"
import { Member } from "~/domain/workspace/member"

interface Props {
  member: Member
}

export const MemberCard = ({ member }: Props) => {
  return (
    <div className="flex flex-wrap ju text-left">
      <Avatar className="mr-auto w-24 h-24 rounded-[36px]">
        <AvatarImage src={member.user.photo?.photoURL} alt={"aa"} />
        <AvatarFallback>{member.profile.displayName?.firstTwoCharacters}</AvatarFallback>
      </Avatar>
      <div className="mt-4 w-full">
        <p className="font-bold">{member.profile.displayName?.value}</p>
        <span className="text-sm text-muted-foreground">Co-Founder, CEO</span>
      </div>
    </div>
  )
}
