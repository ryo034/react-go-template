import { useContext, useState } from "react"
import { Separator, useToast } from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"
import { SettingsProfileForm, SettingsProfileFormValues } from "./form"
import { SettingsProfileUploadPhotoForm, SettingsProfileUploadPhotoFormValues } from "./uploadPhotoForm"

export const settingsProfilePageRoute = "/settings/profile"

export const SettingsProfilePage = () => {
  const { store, controller } = useContext(ContainerContext)
  const { toast } = useToast()

  const [isUpdating, setIsUpdating] = useState(false)

  const me = store.me((state) => state.me)

  if (!me || !me.member) return <></>

  const onSubmit = async (d: SettingsProfileFormValues) => {
    if (isUpdating) return
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

  const onSubmitUploadPhoto = async (d: SettingsProfileUploadPhotoFormValues) => {
    if (isUpdating) return
    setIsUpdating(true)
    const err = await controller.me.updatePhoto(d.photo)
    setIsUpdating(false)
    if (err) {
      toast({ title: "Failed to update profile" })
      return
    }
    toast({ title: "Profile updated" })
  }

  const onClickRemoveProfilePhotoButton = () => {
    if (isUpdating) return
    console.log("remove profile photo button clicked")
  }

  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-lg font-medium">Profile</h3>
        <p className="text-sm text-muted-foreground">This is how others will see you on the site.</p>
      </div>
      <Separator />
      <SettingsProfileUploadPhotoForm
        me={me}
        isUpdating={isUpdating}
        onSubmit={onSubmitUploadPhoto}
        onClickRemoveProfilePhotoButton={onClickRemoveProfilePhotoButton}
      />

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
