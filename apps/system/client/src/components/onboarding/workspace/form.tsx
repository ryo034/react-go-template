import { useEffect } from "react"
import { SubmitHandler, useForm } from "react-hook-form"
import { Button, FormResultErrorMessage, LoadingButton } from "shared-ui"
import { FormInputSection } from "~/components/common/form/inputSection"
import { AccountName } from "~/domain"
import { useOnboardingSettingWorkspacePageFormMessage } from "./message"

export type OnboardingSettingWorkspacePageFormValues = {
  subdomain: string
}

interface Props {
  onSubmit: SubmitHandler<OnboardingSettingWorkspacePageFormValues>
  errorMessage: string
  isLoading: boolean
}

const onboardingSettingWorkspacePageFormId = "onboardingSettingWorkspacePageForm"

export const OnboardingSettingWorkspacePageForm = ({ onSubmit, errorMessage, isLoading = false }: Props) => {
  const {
    register,
    handleSubmit,
    setFocus,
    formState: { errors }
  } = useForm<OnboardingSettingWorkspacePageFormValues>()

  const message = useOnboardingSettingWorkspacePageFormMessage()

  const subdomainInputField = register("subdomain", {
    required: message.form.validation.subdomain.required,
    maxLength: {
      value: AccountName.max,
      message: message.form.validation.subdomain.max
    },
    pattern: {
      value: AccountName.pattern,
      message: message.form.validation.subdomain.regex
    }
  })

  useEffect(() => {
    setFocus("subdomain")
  }, [setFocus])

  return (
    <form
      className="space-y-6 max-w-[320px] m-auto"
      id={onboardingSettingWorkspacePageFormId}
      onSubmit={handleSubmit(onSubmit)}
    >
      <div className="space-y-2 flex items-center">
        <span className="text-gray-700 mr-3">example.com/</span>
        <FormInputSection
          fullWidth
          required
          title={message.word.accountName}
          id="subdomain"
          showLabel={false}
          placeholder={message.form.placeholder.name}
          reactHookForm={subdomainInputField}
          errorMessage={errors.subdomain?.message ?? ""}
        />
      </div>
      <FormResultErrorMessage message={errorMessage} />
      {isLoading ? (
        <LoadingButton fullWidth data-testid="loadingButton" />
      ) : (
        <Button fullWidth type="submit" form={onboardingSettingWorkspacePageFormId} data-testid="nextButton">
          {message.action.submit}
        </Button>
      )}
    </form>
  )
}
