import { render, waitFor } from "@testing-library/react"
import userEvent from "@testing-library/user-event"
import { beforeEach, describe, expect, test, vi } from "vitest"
import { FormPasswordInputSection } from "~/components/common/form/inputPassword"

describe("FormPasswordInputSection", () => {
  beforeEach(() => {
    vi.mock("../../../src/components/common/form/message.tsx", () => ({
      usePasswordInputComponentMessage: () => {
        return {
          action: {
            showPassword: "パスワードを表示する"
          }
        }
      }
    }))
  })

  test("switch password visible", async () => {
    const elm = (
      <FormPasswordInputSection showToggle title="パスワード" id="password" placeholder="••••••••" errorMessage={""} />
    )
    const { getByTestId } = render(elm)
    expect(getByTestId("password").getAttribute("type")).toBe("password")
    const checkElm = getByTestId("password-togglePasswordVisibility")
    const inputElm = checkElm.getElementsByTagName("input")[0]
    waitFor(() => expect(inputElm.checked).toBeFalsy())
    await userEvent.click(inputElm)
    waitFor(() => expect(inputElm.checked).toBeTruthy())
    waitFor(() => expect(getByTestId("password").getAttribute("type")).toBe("text"))
  })
})
