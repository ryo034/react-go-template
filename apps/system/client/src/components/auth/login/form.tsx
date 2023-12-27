import { MouseEventHandler, useEffect } from "react"
import { SubmitHandler, useForm } from "react-hook-form"
import { Button, FormResultErrorMessage, Text } from "shared-ui"
import { useLoginPageFormMessage } from "~/components/auth/login/message"
import { FormPasswordInputSection } from "~/components/common/form/inputPassword"
import { FormInputSection } from "~/components/common/form/inputSection"
import { Email, Password } from "~/domain"

export type LoginFormValues = {
  email: string
  password: string
}

interface Props {
  onSubmit: SubmitHandler<LoginFormValues>
  onClickGoToSignUpPage: MouseEventHandler<HTMLParagraphElement>
  onClickForgotPassword: MouseEventHandler<HTMLParagraphElement>
  errorMessage: string
}

const loginFormId = "loginForm"

export const LoginPageForm = ({ onSubmit, onClickGoToSignUpPage, onClickForgotPassword, errorMessage }: Props) => {
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

  const passwordInputField = register("password", {
    required: message.form.validation.password.required,
    pattern: {
      value: Password.pattern,
      message: message.form.validation.password.regex
    }
  })

  useEffect(() => {
    setFocus("email")
  }, [])

  return (
    <form className="space-y-6" id={loginFormId} onSubmit={handleSubmit(onSubmit)}>
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
      <FormPasswordInputSection
        fullWidth
        isCurrent
        title={message.word.password}
        id="password"
        placeholder="••••••••"
        reactHookForm={passwordInputField}
        errorMessage={errors.password?.message ?? ""}
      />
      <div className="flex items-center justify-end">
        <Text
          onClick={onClickForgotPassword}
          className="text-xs cursor-pointer hover:underline"
          data-testid="forgotPasswordTextButton"
        >
          {message.forgotPassword}
        </Text>
      </div>
      <FormResultErrorMessage message={errorMessage} />
      <Button fullWidth type="submit" form={loginFormId} data-testid="loginButton">
        {message.action.login}
      </Button>
      <Text>
        {message.notHaveAnAccountYet}
        <Text
          asChild
          onClick={onClickGoToSignUpPage}
          className="font-bold cursor-pointer text-primary-600 hover:underline dark:text-primary-500"
        >
          <span>{message.action.signUp}</span>
        </Text>
      </Text>
    </form>
  )
}
