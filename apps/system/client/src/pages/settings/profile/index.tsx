import { useContext, useState } from "react"
import { Separator, useToast } from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"
import { SettingsProfileForm, SettingsProfileFormValues } from "./form"

export const settingsProfilePageRoute = "/settings/profile"

export const SettingsProfilePage = () => {
  const { store, controller } = useContext(ContainerContext)
  const { toast } = useToast()

  const [isUpdating, setIsUpdating] = useState(false)

  const me = store.me((state) => state.me)

  if (!me) return <></>

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

  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-lg font-medium">Profile</h3>
        <p className="text-sm text-muted-foreground">This is how others will see you on the site.</p>
      </div>
      <Separator />
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
