import { render, screen, waitFor } from "@testing-library/react"
import userEvent from "@testing-library/user-event"
import { SubmitHandler } from "react-hook-form"
import { afterEach, beforeEach, describe, expect, it, test, vi } from "vitest"
import { OnboardingSettingNamePageForm, OnboardingSettingNamePageFormValues } from "~/components/onboarding/name/form"
import { OnboardingSettingNamePageFormMessage } from "~/components/onboarding/name/message"

const nameErrorPatterns = [
  ["", "アカウント名は必須です"],
  [
    "あいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあ",
    "アカウント名は50文字以内で入力してください"
  ],
  ["test-test", "アカウント名は半角英数字とハイフンのみ使用できます"]
]

describe("OnboardingSettingNamePageForm", () => {
  let mockOnSubmit: SubmitHandler<OnboardingSettingNamePageFormValues>

  beforeEach(() => {
    mockOnSubmit = vi.fn()
    vi.mock("../../../../src/components/onboarding/name/message.tsx", () => ({
      useOnboardingSettingNamePageFormMessage: () => {
        const res: OnboardingSettingNamePageFormMessage = {
          word: { accountName: "アカウント名" },
          action: { submit: "送信" },
          form: {
            placeholder: { name: "アカウント名" },
            validation: {
              name: {
                required: "アカウント名は必須です",
                max: "アカウント名は50文字以内で入力してください",
                regex: "アカウント名は半角英数字とハイフンのみ使用できます"
              }
            }
          }
        }
        return res
      }
    }))
  })

  afterEach(() => {
    vi.resetAllMocks()
  })

  it("fills and submits the form correctly", async () => {
    const elm = <OnboardingSettingNamePageForm onSubmit={mockOnSubmit} errorMessage="" isLoading={false} />
    render(elm)

    const t = screen.getByTestId("name")
    expect(t).toBeInTheDocument()
    await userEvent.type(t, "test")
    await screen.getByTestId("nextButton").click()
    await waitFor(() => expect(mockOnSubmit).toHaveBeenCalledTimes(1))
  })

  test.each(nameErrorPatterns)("input: '%s', expected error message: '%s'", async (pattern, message) => {
    const elm = <OnboardingSettingNamePageForm onSubmit={mockOnSubmit} errorMessage={""} isLoading={false} />
    render(elm)
    const t = screen.getByTestId("name")
    expect(t).toBeInTheDocument()
    await userEvent.type(t, `${pattern}{enter}`)
    const errorMessageElm = screen.queryByTestId("name-errorMessage")
    if (!errorMessageElm) {
      throw new Error("errorMessageElm is null")
    }
    await waitFor(() => expect(errorMessageElm.textContent).toBe(message))
  })

  it("displays error message when provided", () => {
    const errorMessage = "Error message"
    const elm = <OnboardingSettingNamePageForm onSubmit={mockOnSubmit} errorMessage={errorMessage} isLoading={false} />
    const { getByText } = render(elm)
    expect(getByText(errorMessage)).toBeInTheDocument()
  })

  it("displays loading button when loading", () => {
    const elm = <OnboardingSettingNamePageForm onSubmit={mockOnSubmit} errorMessage="" isLoading={true} />
    render(elm)
    const nextButton = screen.queryByTestId("nextButton")
    expect(nextButton).toBeNull()
    expect(screen.getByTestId("loadingButton")).toBeInTheDocument()
  })
})
