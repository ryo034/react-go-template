import { useEffect } from "react"
import { SubmitHandler, useForm } from "react-hook-form"
import { Button, FormResultErrorMessage, LoadingButton } from "shared-ui"
import { FormInputSection } from "~/components/common/form/inputSection"
import { WorkspaceSubdomain } from "~/domain"
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
      value: WorkspaceSubdomain.max,
      message: message.form.validation.subdomain.max
    },
    pattern: {
      value: WorkspaceSubdomain.pattern,
      message: message.form.validation.subdomain.regex
    }
  })

  useEffect(() => {
    setFocus("subdomain")
  }, [setFocus])

  return (
    <form
      className="space-y-12 m-auto flex flex-wrap justify-center"
      id={onboardingSettingWorkspacePageFormId}
      onSubmit={handleSubmit(onSubmit)}
    >
      <div className="flex items-center justify-center">
        <span className="text-gray-700 mr-3">example.com/</span>
        <FormInputSection
          fullWidth
          required
          title={message.word.subdomain}
          id="subdomain"
          showLabel={false}
          placeholder={message.form.placeholder.name}
          reactHookForm={subdomainInputField}
          errorMessage={""} // hide error
        />
      </div>
      <FormResultErrorMessage message={errorMessage} />
      {isLoading ? (
        <LoadingButton className="max-w-[320px] w-full" />
      ) : (
        <Button
          className="max-w-[320px] w-full"
          type="submit"
          form={onboardingSettingWorkspacePageFormId}
          data-testid="nextButton"
        >
          {message.action.submit}
        </Button>
      )}
    </form>
  )
}
