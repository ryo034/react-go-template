import { Avatar, AvatarFallback, AvatarImage } from "shared-ui"
import { Member } from "~/domain/workspace/member"

interface Props {
  member: Member
}

export const MemberCard = ({ member }: Props) => {
  return (
    <div className="flex flex-wrap ju text-left">
      <Avatar className="mr-auto w-24 h-24 rounded-[36px]">
        <AvatarImage
          src={
            "https://img.freepik.com/free-psd/3d-illustration-of-person-with-sunglasses_23-2149436188.jpg?w=826&t=st=1707604373~exp=1707604973~hmac=86a0d39e4a6cfe5fac7e6c036015ce5a216cec8360cce331ce803d62b3541e3b"
          }
          alt={"aa"}
        />
        <AvatarFallback>SC</AvatarFallback>
      </Avatar>
      <div className="mt-4 w-full">
        <p className="font-bold">{member.profile.displayName?.value}</p>
        <span className="text-sm text-muted-foreground">Co-Founder, CEO</span>
      </div>
    </div>
  )
}
