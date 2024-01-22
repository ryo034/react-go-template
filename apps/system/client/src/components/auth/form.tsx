import { MouseEventHandler, useEffect } from "react"
import { SubmitHandler, useForm } from "react-hook-form"
import { Button, FormResultErrorMessage, Separator } from "shared-ui"
import { useAuthPageFormMessage } from "~/components/auth/message"
import { FormInputSection } from "~/components/common/form/inputSection"
import { Email } from "~/domain"

export type LoginFormValues = {
  email: string
}

interface Props {
  onSubmit: SubmitHandler<LoginFormValues>
  onClickGoogleLoginButton: MouseEventHandler<HTMLButtonElement>
  errorMessage: string
}

const authFormId = "authForm"

export const AuthPageForm = ({ onSubmit, onClickGoogleLoginButton, errorMessage }: Props) => {
  const {
    register,
    handleSubmit,
    setFocus,
    formState: { errors }
  } = useForm<LoginFormValues>()

  const message = useAuthPageFormMessage()

  const emailInputField = register("email", {
    required: message.form.validation.email.required,
    maxLength: {
      value: Email.max,
      message: message.form.validation.email.max
    },
    pattern: {
      value: Email.pattern,
      message: message.form.validation.email.regex
    }
  })

  useEffect(() => {
    setFocus("email")
  }, [])

  return (
    <>
      <form className="space-y-6" id={authFormId} onSubmit={handleSubmit(onSubmit)}>
        <div>
          <Button fullWidth type="button" onClick={onClickGoogleLoginButton} data-testid="googleLoginButton">
            Googleでログイン
          </Button>
        </div>

        <Separator />

        <FormInputSection
          fullWidth
          title={message.word.email}
          id="email"
          type="email"
          placeholder="name@company.com"
          autoComplete="email"
          reactHookForm={emailInputField}
          errorMessage={errors.email?.message ?? ""}
        />
        <FormResultErrorMessage message={errorMessage} />
        <Button fullWidth type="submit" form={authFormId} data-testid="startButton">
          {message.action.startWithEmail}
        </Button>
      </form>
    </>
  )
}