import { render, screen, waitFor } from "@testing-library/react"
import userEvent from "@testing-library/user-event"
import { SubmitHandler } from "react-hook-form"
import { afterEach, beforeEach, it } from "vitest"
import { describe, expect, vi } from "vitest"
import { OtpFormValues, VerifyOTPPageForm, verifyOtpFormId } from "~/components/auth/otp/form"

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

    const otpInput1 = screen.getByTestId("otpInput1")
    expect(otpInput1).toHaveFocus()

    for (let idx = 0; idx < 6; idx++) {
      const t = screen.getByTestId(`otpInput${idx + 1}`)
      expect(t).toBeInTheDocument()
      await userEvent.type(t, `${idx + 1}`)
    }

    // check auto submit
    await waitFor(() => expect(mockOnSubmit).toHaveBeenCalledTimes(1))

    await screen.getByTestId("verifyOtpButton").click()
    await waitFor(() => expect(mockOnSubmit).toHaveBeenCalledTimes(2))
  })

  it("automatically filled all field when paste OTP from clipboard", async () => {
    Object.assign(navigator, {
      clipboard: {
        readText() {
          return Promise.resolve("123456")
        }
      }
    })

    const elm = <VerifyOTPPageForm onSubmit={mockOnSubmit} errorMessage="" />
    render(elm)

    await userEvent.paste("123456")
    // auto fill and auto submit
    await waitFor(() => {
      expect(screen.getByTestId("otpInput1")).toHaveValue(1)
      expect(screen.getByTestId("otpInput2")).toHaveValue(2)
      expect(screen.getByTestId("otpInput3")).toHaveValue(3)
      expect(screen.getByTestId("otpInput4")).toHaveValue(4)
      expect(screen.getByTestId("otpInput5")).toHaveValue(5)
      expect(screen.getByTestId("otpInput6")).toHaveValue(6)
      expect(mockOnSubmit).toHaveBeenCalledTimes(1)
    })
  })

  it("displays error message when provided", () => {
    const errorMessage = "Error message"
    const elm = <VerifyOTPPageForm onSubmit={mockOnSubmit} errorMessage={errorMessage} />
    const { getByText } = render(elm)
    expect(getByText(errorMessage)).toBeInTheDocument()
  })

  it("When pasting a non OTP from the clipboard, the field is not automatically filled in and an error toast appears", async () => {
    const clipboardFailValue = "12345あ"
    Object.assign(navigator, {
      clipboard: {
        readText() {
          return Promise.resolve(clipboardFailValue)
        }
      }
    })

    const elm = <VerifyOTPPageForm onSubmit={mockOnSubmit} errorMessage="" />
    render(elm)
    const form = screen.getByTestId(verifyOtpFormId)
    await userEvent.paste(clipboardFailValue)
    expect(form).toHaveFormValues({
      otpInput1: null,
      otpInput2: null,
      otpInput3: null,
      otpInput4: null,
      otpInput5: null,
      otpInput6: null
    })
    await waitFor(() => expect(mockOnSubmit).toHaveBeenCalledTimes(0))
  })
})
