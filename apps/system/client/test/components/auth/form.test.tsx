import { render, screen, waitFor } from "@testing-library/react"
import userEvent from "@testing-library/user-event"
import type { MouseEventHandler } from "react"
import type { SubmitHandler } from "react-hook-form"
import { afterEach } from "vitest"
import { beforeEach, describe, expect, test, vi } from "vitest"
import { AuthPageForm, type LoginFormValues } from "~/components/auth/form"

const emailRequiredErrorMessage = "メールアドレスは必須です"
const emailRegexErrorMessage = "正しいメールアドレスを入力して下さい"

const emailPatterns = [
  ["", emailRequiredErrorMessage],
  ["test", emailRegexErrorMessage]
]

describe("LoginForm", () => {
  let mockOnSubmit: SubmitHandler<LoginFormValues>
  let mockOnClickGoogleLoginButton: MouseEventHandler<HTMLButtonElement>

  beforeEach(() => {
    mockOnSubmit = vi.fn()
    mockOnClickGoogleLoginButton = vi.fn()

    vi.mock("../../../src/components/auth/message.tsx", () => ({
      useAuthPageFormMessage: () => {
        return {
          word: {
            email: "メールアドレス"
          },
          action: {
            login: "ログイン",
            sendOneTimeCode: "ワンタイムパスワードを送信する",
            startWithEmail: "メールアドレスで始める",
            startWithGoogle: "Googleで始める"
          },
          form: {
            validation: {
              email: {
                required: emailRequiredErrorMessage,
                regex: emailRegexErrorMessage
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
          <AuthPageForm
            onSubmit={mockOnSubmit}
            errorMessage=""
            onClickGoogleLoginButton={mockOnClickGoogleLoginButton}
            isLoading={false}
          />
        )
        const { getByTestId } = render(elm)

        const submitButton = getByTestId("startButton")
        await userEvent.click(submitButton)

        const inputField = getByTestId(id)
        await userEvent.type(inputField, `${pattern}{enter}`)
        const errorMessageElm = screen.getByTestId(`${id}-errorMessage`)
        await waitFor(() => expect(errorMessageElm.textContent).toBe(message))
      })
    })
  }

  describe("displays error messages when fields are valid and submit is clicked", async () => {
    await validationCheckTest("email", emailPatterns)
  })

  test("focuses on email field when opened", async () => {
    render(
      <AuthPageForm
        onSubmit={mockOnSubmit}
        errorMessage=""
        onClickGoogleLoginButton={mockOnClickGoogleLoginButton}
        isLoading={false}
      />
    )
    expect(screen.getByTestId("email")).toEqual(document.activeElement)
  })

  test("submits correct values when form is filled out and submit is clicked", async () => {
    const elm = (
      <AuthPageForm
        onSubmit={mockOnSubmit}
        errorMessage=""
        onClickGoogleLoginButton={mockOnClickGoogleLoginButton}
        isLoading={false}
      />
    )
    const { getByTestId } = render(elm)
    const emailField = getByTestId("email")
    await userEvent.type(emailField, "test@example.com")

    const submitButton = getByTestId("startButton")
    await userEvent.click(submitButton)

    await waitFor(() => expect(mockOnSubmit).toHaveBeenCalledTimes(1))
  })

  test("displays loading button when loading", () => {
    const elm = (
      <AuthPageForm
        onSubmit={mockOnSubmit}
        errorMessage=""
        onClickGoogleLoginButton={mockOnClickGoogleLoginButton}
        isLoading={true}
      />
    )
    render(elm)
    const nextButton = screen.queryByTestId("startButton")
    expect(nextButton).toBeNull()
    expect(screen.getByTestId("loadingButton")).toBeInTheDocument()
  })
})
