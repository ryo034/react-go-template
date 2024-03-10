import { useContext } from "react"
import { Separator } from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"
import { type AppearanceFormValues, SettingsAppearanceForm } from "./form"

export const settingsAppearancePageRoute = "/settings/appearance"

export const SettingsAppearancePage = () => {
  const { controller } = useContext(ContainerContext)

  const onSubmit = ({ theme }: AppearanceFormValues) => {
    controller.theme.toggle(theme)
  }

  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-lg font-medium">Appearance</h3>
        <p className="text-sm text-muted-foreground">
          Customize the appearance of the app. Automatically switch between day and night themes.
        </p>
      </div>
      <Separator />
      <SettingsAppearanceForm onSubmit={onSubmit} />
    </div>
  )
}
