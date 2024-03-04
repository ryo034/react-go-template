import { useContext, useState } from "react"
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
  Button,
  Separator,
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
  useToast
} from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"
import { SettingsProfileForm, SettingsProfileFormValues } from "./form"

export const settingsProfilePageRoute = "/settings/profile"

export const SettingsProfilePage = () => {
  const { store, controller } = useContext(ContainerContext)
  const { toast } = useToast()

  const [isUpdating, setIsUpdating] = useState(false)

  const me = store.me((state) => state.me)

  if (!me || !me.member) return <></>

  const onSubmit = async (d: SettingsProfileFormValues) => {
    setIsUpdating(true)
    const err = await controller.me.updateMemberProfile({
      displayName: d.displayName,
      bio: d.bio,
      idNumber: ""
    })
    setIsUpdating(false)
    if (err) {
      toast({ title: "Failed to update profile" })
      return
    }
    toast({ title: "Profile updated" })
  }

  const onAvatarClick = () => {
    console.log("avatar clicked")
  }

  const onClickUpdateProfilePhotoButton = () => {
    console.log("update profile photo button clicked")
  }

  const onClickRemoveProfilePhotoButton = () => {
    console.log("remove profile photo button clicked")
  }

  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-lg font-medium">Profile</h3>
        <p className="text-sm text-muted-foreground">This is how others will see you on the site.</p>
      </div>
      <Separator />
      <Avatar className="mr-auto w-24 h-24 rounded-[36px] cursor-pointer" onClick={onAvatarClick}>
        <AvatarImage
          // src={
          //   "https://img.freepik.com/free-psd/3d-illustration-of-person-with-sunglasses_23-2149436188.jpg?w=826&t=st=1707604373~exp=1707604973~hmac=86a0d39e4a6cfe5fac7e6c036015ce5a216cec8360cce331ce803d62b3541e3b"
          // }
          src={me.member.profile.photo?.url || ""}
          alt={"aa"}
        />
        <AvatarFallback className="rounded-[36px]">{me.member.profile.displayName?.firstTwoCharacters}</AvatarFallback>
      </Avatar>
      <div className="flex space-x-4">
        <Button type="button" onClick={onClickUpdateProfilePhotoButton} data-testid="updateProfilePhoto">
          {me.member.profile.hasPhoto ? "画像を変更" : "画像を追加"}
        </Button>
        <Button
          type="button"
          variant="ghost"
          onClick={onClickRemoveProfilePhotoButton}
          data-testid="removeProfilePhoto"
        >
          <span className="text-red-600">画像を削除</span>
        </Button>
      </div>

      <SettingsProfileForm
        isUpdating={isUpdating}
        onSubmit={onSubmit}
        defaultValues={{
          displayName: me?.member?.profile.displayName?.value || "",
          bio: me?.member?.profile.bio.value || ""
        }}
      />
    </div>
  )
}
