import { render, screen, waitFor } from "@testing-library/react"
import userEvent from "@testing-library/user-event"
import { MouseEventHandler } from "react"
import { SubmitHandler } from "react-hook-form"
import { afterEach } from "vitest"
import { beforeEach, describe, expect, test, vi } from "vitest"
import { LoginFormValues, LoginPageForm } from "~/components/auth/login/form"

const emailRequiredErrorMessage = "メールアドレスは必須です"
const emailRegexErrorMessage = "正しいメールアドレスを入力して下さい"
const passwordRequiredErrorMessage = "パスワードは必須です"
const passwordRegexErrorMessage =
  "パスワードは大文字・小文字・半角英数字を1つずつ含む6文字以上128文字以下で入力して下さい"

const emailPatterns = [
  ["", emailRequiredErrorMessage],
  ["test", emailRegexErrorMessage]
]

const passwordPatterns = [
  ["", passwordRequiredErrorMessage],
  ["test", passwordRegexErrorMessage]
]

describe("LoginForm", () => {
  let mockOnSubmit: SubmitHandler<LoginFormValues>
  let onClickGoToSignUpPage: MouseEventHandler<HTMLParagraphElement>
  let onClickForgotPassword: MouseEventHandler<HTMLParagraphElement>

  beforeEach(() => {
    mockOnSubmit = vi.fn()
    onClickGoToSignUpPage = vi.fn()
    onClickForgotPassword = vi.fn()
    vi.mock("../../../../src/components/common/form/message.tsx", () => ({
      usePasswordInputComponentMessage: () => {
        return {
          action: {
            showPassword: "パスワードを表示する"
          }
        }
      }
    }))
    vi.mock("../../../../src/components/auth/login/message.tsx", () => ({
      useLoginPageFormMessage: () => {
        return {
          forgotPassword: "パスワードを忘れましたか？",
          notHaveAnAccountYet: "まだアカウントをお持ちではありませんか？",
          createNewAccount: "アカウントの新規作成",
          word: {
            email: "メールアドレス",
            password: "パスワード"
          },
          action: {
            login: "ログイン",
            signUp: "新規作成"
          },
          form: {
            validation: {
              email: {
                required: emailRequiredErrorMessage,
                regex: emailRegexErrorMessage
              },
              password: {
                required: passwordRequiredErrorMessage,
                regex: passwordRegexErrorMessage
              }
            }
          }
        }
      }
    }))
  })

  afterEach(() => {
    vi.resetAllMocks()
  })

  const validationCheckTest = async (id: string, patterns: string[][]) => {
    describe(id, () => {
      test.each(patterns)("input: '%s', expected error message: '%s'", async (pattern, message) => {
        const elm = (
          <LoginPageForm
            onClickForgotPassword={onClickForgotPassword}
            onSubmit={mockOnSubmit}
            errorMessage=""
            onClickGoToSignUpPage={onClickGoToSignUpPage}
          />
        )
        const { getByTestId } = render(elm)

        const submitButton = getByTestId("loginButton")
        await userEvent.click(submitButton)

        const inputField = getByTestId(id)
        await userEvent.type(inputField, `${pattern}{enter}`)
        const errorMessageElm = screen.getByTestId(`${id}-errorMessage`)
        await waitFor(() => expect(errorMessageElm.textContent).toBe(message))
      })
    })
  }

  test("focuses on email field when opened", async () => {
    render(
      <LoginPageForm
        onClickForgotPassword={onClickForgotPassword}
        onSubmit={mockOnSubmit}
        errorMessage=""
        onClickGoToSignUpPage={onClickGoToSignUpPage}
      />
    )
    expect(screen.getByTestId("email")).toEqual(document.activeElement)
  })

  describe("displays error messages when fields are valid and submit is clicked", async () => {
    validationCheckTest("email", emailPatterns)
    validationCheckTest("password", passwordPatterns)
  })

  test("submits correct values when form is filled out and submit is clicked", async () => {
    const elm = (
      <LoginPageForm
        onClickForgotPassword={onClickForgotPassword}
        onSubmit={mockOnSubmit}
        errorMessage=""
        onClickGoToSignUpPage={onClickGoToSignUpPage}
      />
    )
    const { getByTestId } = render(elm)
    const emailField = getByTestId("email")
    const passwordField = getByTestId("password")

    await userEvent.type(emailField, "test@example.com")
    await userEvent.type(passwordField, "Test123")

    const submitButton = getByTestId("loginButton")
    await userEvent.click(submitButton)

    await waitFor(() => expect(mockOnSubmit).toHaveBeenCalledTimes(1))
  })

  test("call onClickGoToSignUpPage when 新規作成 is clicked", async () => {
    const elm = (
      <LoginPageForm
        onClickForgotPassword={onClickForgotPassword}
        onSubmit={mockOnSubmit}
        errorMessage=""
        onClickGoToSignUpPage={onClickGoToSignUpPage}
      />
    )
    const { getByText } = render(elm)
    const target = getByText("新規作成")
    await userEvent.click(target)
    await waitFor(() => expect(onClickGoToSignUpPage).toHaveBeenCalledTimes(1))
  })

  test("call onClickForgotPassword when パスワードを忘れましたか？ is clicked", async () => {
    const elm = (
      <LoginPageForm
        onClickForgotPassword={onClickForgotPassword}
        onSubmit={mockOnSubmit}
        errorMessage=""
        onClickGoToSignUpPage={onClickGoToSignUpPage}
      />
    )
    const { getByTestId } = render(elm)
    const target = getByTestId("forgotPasswordTextButton")
    await userEvent.click(target)
    await waitFor(() => expect(onClickForgotPassword).toHaveBeenCalledTimes(1))
  })
})
