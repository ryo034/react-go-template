import { render, screen, waitFor } from "@testing-library/react"
import userEvent from "@testing-library/user-event"
import { SubmitHandler } from "react-hook-form"
import { afterEach, beforeEach, it } from "vitest"
import { describe, expect, vi } from "vitest"
import { OtpFormValues, VerifyOTPPageForm } from "~/components/auth/otp/form"

describe("VerifyOTPPageForm", () => {
  let mockOnSubmit: SubmitHandler<OtpFormValues>

  beforeEach(() => {
    mockOnSubmit = vi.fn()
    vi.mock("../../../../src/components/auth/otp/message.tsx", () => ({
      useVerifyOtpPageFormMessage: () => {
        return { word: { submit: "送信" } }
      }
    }))
  })

  afterEach(() => {
    vi.resetAllMocks()
  })

  it("fills and submits the form correctly", async () => {
    const elm = <VerifyOTPPageForm onSubmit={mockOnSubmit} errorMessage="" />
    render(elm)

    for (let idx = 0; idx < 6; idx++) {
      const t = screen.getByTestId(`otpInput${idx + 1}`)
      expect(t).toBeInTheDocument()
      await userEvent.type(t, `${idx + 1}`)
    }

    // check auto submit
    await waitFor(() => {
      expect(mockOnSubmit).toHaveBeenCalledTimes(1)
    })

    await screen.getByTestId("verifyOtpButton").click()
    await waitFor(() => {
      expect(mockOnSubmit).toHaveBeenCalledTimes(2)
    })
  })

  it("displays error message when provided", () => {
    const errorMessage = "Error message"
    const elm = <VerifyOTPPageForm onSubmit={mockOnSubmit} errorMessage={errorMessage} />
    const { getByText } = render(elm)
    expect(getByText(errorMessage)).toBeInTheDocument()
  })

  it("has a functional submit button", () => {
    const elm = <VerifyOTPPageForm onSubmit={mockOnSubmit} errorMessage="" />
    const { getByTestId } = render(elm)
    expect(getByTestId("verifyOtpButton")).toBeInTheDocument()
  })
})
