import { MouseEventHandler, useEffect } from "react"
import { SubmitHandler, useForm } from "react-hook-form"
import { Button, FormResultErrorMessage, Separator } from "shared-ui"
import { useLoginPageFormMessage } from "~/components/auth/login/message"
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

const loginFormId = "loginForm"

export const LoginPageForm = ({ onSubmit, onClickGoogleLoginButton, errorMessage }: Props) => {
  const {
    register,
    handleSubmit,
    setFocus,
    formState: { errors }
  } = useForm<LoginFormValues>()

  const message = useLoginPageFormMessage()

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
      <form className="space-y-6" id={loginFormId} onSubmit={handleSubmit(onSubmit)}>
        <div>
          <Button fullWidth type="button" onClick={onClickGoogleLoginButton} data-testid="googleLoginButton">
            Googleログイン
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
        <Button fullWidth type="submit" form={loginFormId} data-testid="loginButton">
          {message.action.sendOneTimeCode}
        </Button>
      </form>
    </>
  )
}
