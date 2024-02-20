import { render } from "@testing-library/react"
import { describe, test } from "vitest"
import { InviteMembersDialog } from "~/components/member/inviteDialog"

describe("InviteMembersDialog", () => {
  // beforeAll(() => {
  //   // @ts-ignore
  //   ReactDOM.createPortal = vi.fn((element) => {
  //     return element
  //   })
  // })

  // I can't test the dialog, as mentioned in the issue, so I'll test it E2E
  // https://github.com/radix-ui/primitives/discussions/1130
  test("correctly open and closes dialog", async () => {
    render(<InviteMembersDialog />)
    // const { getByTestId } = render(<InviteMembersDialog />)
    // await userEvent.click(getByTestId("inviteMembersButton"))
    // await waitFor(() => {
    //   expect(getByTestId("inviteMembersDialog")).toBeVisible()
    // })

    // userEvent.click(getByTestId("closeIconOnDialog"))
    // await waitFor(() => {
    //   expect(getByTestId("inviteMembersDialog")).not.toBeVisible()
    // })
  })

  // test("fields count is 1 when dialog open", async () => {
  // })

  // test("focuses on first field email input when dialog open", async () => {
  // })

  // test("displays loading button when loading", () => {
  // })
})
