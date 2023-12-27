import { describe, expect, it } from "vitest"
import { DomainError, Phone } from "~/domain/shared"

describe("Phone", () => {
  describe("create", () => {
    it.each([
      "",
      "あああああ",
      "090123412",
      "1234567890",
      "00012341234",
      "01012341234",
      "02012341234",
      "03012341234",
      "04012341234",
      "05012341234",
      "06012341234"
    ])("DomainError %s", (actual) => {
      const result = Phone.create(actual)
      expect(result.isErr).toBeTruthy()
      result.mapErr((e) => expect(e).instanceOf(DomainError))
    })

    it("OK", () => {
      expect(Phone.create("09012341234").isOk).toBeTruthy()
      expect(Phone.create("09000000000").isOk).toBeTruthy()
    })
  })
})
