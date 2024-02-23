import { Separator } from "shared-ui"
import { AccountForm } from "./form"

export const settingsAccountPageRoute = "/settings/account"

export const SettingsAccountPage = () => {
  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-lg font-medium">Account</h3>
        <p className="text-sm text-muted-foreground">
          Update your account settings. Set your preferred language and timezone.
        </p>
      </div>
      <Separator />
      <AccountForm />
    </div>
  )
}
