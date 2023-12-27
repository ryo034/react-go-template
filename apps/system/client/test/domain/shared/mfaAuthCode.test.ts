import { describe, expect, it } from "vitest"
import { DomainError, MFAAuthCode } from "~/domain/shared"

describe("mfaAuthCode", () => {
  describe("create", () => {
    it.each(["1", "12", "123", "1234", "12345"])("DomainError %s", (actual) => {
      const result = MFAAuthCode.create(actual)
      expect(result.isErr).toBeTruthy()
      result.mapErr((e) => expect(e).instanceOf(DomainError))
    })
    it("OK", () => {
      expect(MFAAuthCode.create("123456").isOk).toBeTruthy()
    })
  })
})
