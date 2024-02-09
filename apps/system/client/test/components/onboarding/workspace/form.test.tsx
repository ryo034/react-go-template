import { render, screen, waitFor } from "@testing-library/react"
import userEvent from "@testing-library/user-event"
import { SubmitHandler } from "react-hook-form"
import { afterEach, beforeEach, it } from "vitest"
import { describe, expect, vi } from "vitest"
import {
  OnboardingSettingWorkspacePageForm,
  OnboardingSettingWorkspacePageFormValues
} from "~/components/onboarding/workspace/form"
import { OnboardingSettingWorkspacePageFormMessage } from "~/components/onboarding/workspace/message"
import { WorkspaceSubdomain } from "~/domain"

describe("OnboardingSettingWorkspacePageForm", () => {
  let mockOnSubmit: SubmitHandler<OnboardingSettingWorkspacePageFormValues>

  beforeEach(() => {
    mockOnSubmit = vi.fn()
    vi.mock("../../../../src/components/onboarding/workspace/message.tsx", () => ({
      useOnboardingSettingWorkspacePageFormMessage: () => {
        const res: OnboardingSettingWorkspacePageFormMessage = {
          word: { subdomain: "サブドメイン" },
          action: { submit: "送信" },
          form: {
            placeholder: {
              name: "アカウント名"
            },
            validation: {
              subdomain: {
                required: "サブドメインは必須です",
                max: `サブドメインは${WorkspaceSubdomain.max}文字以内で入力してください`,
                regex: "サブドメインは半角英数字とハイフンのみ使用できます"
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
    const elm = <OnboardingSettingWorkspacePageForm onSubmit={mockOnSubmit} errorMessage="" isLoading={false} />
    render(elm)
    const t = screen.getByTestId("subdomain")
    expect(t).toBeInTheDocument()
    await userEvent.type(t, "test")
    await screen.getByTestId("nextButton").click()
    await waitFor(() => expect(mockOnSubmit).toHaveBeenCalledTimes(1))
  })

  it("not displays error message under input when provided", async () => {
    const elm = <OnboardingSettingWorkspacePageForm onSubmit={mockOnSubmit} errorMessage={""} isLoading={false} />
    render(elm)
    const t = screen.getByTestId("subdomain")
    expect(t).toBeInTheDocument()
    await userEvent.type(t, "test test")
    expect(screen.queryByTestId("subdomain-error")).toBeNull()
  })

  it("displays error message when provided", () => {
    const errorMessage = "Error message"
    const elm = (
      <OnboardingSettingWorkspacePageForm onSubmit={mockOnSubmit} errorMessage={errorMessage} isLoading={false} />
    )
    const { getByText } = render(elm)
    expect(getByText(errorMessage)).toBeInTheDocument()
  })

  it("displays loading button when loading", () => {
    const elm = <OnboardingSettingWorkspacePageForm onSubmit={mockOnSubmit} errorMessage="" isLoading={true} />
    render(elm)
    const nextButton = screen.queryByTestId("nextButton")
    expect(nextButton).toBeNull()
    expect(screen.getByTestId("loadingButton")).toBeInTheDocument()
  })
})
