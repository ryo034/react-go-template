import { useEffect } from "react"
import { SubmitHandler, useForm } from "react-hook-form"
import { Button, FormResultErrorMessage, LoadingButton } from "shared-ui"
import { FormInputSection } from "~/components/common/form/inputSection"
import { AccountFullName } from "~/domain"
import { useOnboardingSettingNamePageFormMessage } from "./message"

export type OnboardingSettingNamePageFormValues = {
  name: string
}

interface Props {
  onSubmit: SubmitHandler<OnboardingSettingNamePageFormValues>
  errorMessage: string
  isLoading: boolean
}

const onboardingSettingNamePageFormId = "onboardingSettingNamePageForm"

export const OnboardingSettingNamePageForm = ({ onSubmit, errorMessage, isLoading = false }: Props) => {
  const {
    register,
    handleSubmit,
    setFocus,
    formState: { errors }
  } = useForm<OnboardingSettingNamePageFormValues>()

  const message = useOnboardingSettingNamePageFormMessage()

  const nameInputField = register("name", {
    required: message.form.validation.name.required,
    maxLength: {
      value: AccountFullName.max,
      message: message.form.validation.name.max
    },
    pattern: {
      value: AccountFullName.pattern,
      message: message.form.validation.name.regex
    }
  })

  useEffect(() => {
    setFocus("name")
  }, [setFocus])

  return (
    <form
      className="space-y-6 max-w-[320px] m-auto"
      id={onboardingSettingNamePageFormId}
      onSubmit={handleSubmit(onSubmit)}
    >
      <div className="space-y-2">
        <FormInputSection
          fullWidth
          required
          title={message.word.accountName}
          id="name"
          autoComplete="name"
          showLabel={false}
          placeholder={message.form.placeholder.name}
          reactHookForm={nameInputField}
          errorMessage={errors.name?.message ?? ""}
        />
      </div>
      <FormResultErrorMessage message={errorMessage} />
      {isLoading ? (
        <LoadingButton fullWidth />
      ) : (
        <Button fullWidth type="submit" form={onboardingSettingNamePageFormId} data-testid="nextButton">
          {message.action.submit}
        </Button>
      )}
    </form>
  )
}
