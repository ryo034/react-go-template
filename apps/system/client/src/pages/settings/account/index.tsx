import { useContext, useState } from "react"
import { Separator, useToast } from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"
import { SettingsAccountForm, SettingsAccountFormValues } from "./form"

export const settingsAccountPageRoute = "/settings/account"

export const SettingsAccountPage = () => {
  const { store, controller } = useContext(ContainerContext)
  const { toast } = useToast()

  const [isUpdating, setIsUpdating] = useState(false)

  const me = store.me((state) => state.me)

  if (!me) return <></>

  const onSubmit = async (d: SettingsAccountFormValues) => {
    setIsUpdating(true)
    const err = await controller.me.updateProfile({ name: d.name })
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
        <h3 className="text-lg font-medium">Account</h3>
        <p className="text-sm text-muted-foreground">
          Update your account settings. Set your preferred language and timezone.
        </p>
      </div>
      <Separator />
      <SettingsAccountForm
        defaultValues={{ name: me.self.name?.value ?? "" }}
        onSubmit={onSubmit}
        isUpdating={isUpdating}
      />
    </div>
  )
}
