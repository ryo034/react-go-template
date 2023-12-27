import { describe, expect, it } from "vitest"
import { DomainError, Password } from "~/domain/shared"

describe("Password", () => {
  describe("create", () => {
    it.each([
      "",
      "a",
      "12345",
      "testtest",
      "AAAAAAAA",
      "test123",
      "@test123",
      "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVA1"
    ])("DomainError %s", (actual) => {
      const result = Password.create(actual)
      expect(result.isErr).toBeTruthy()
      result.mapErr((e) => expect(e).instanceOf(DomainError))
    })
    it.each(["Test123", "@Test123"])("OK %s", (actual) => {
      const result = Password.create(actual)
      expect(result.isOk).toBeTruthy()
      result.map((e) => expect(e).instanceOf(Password))
    })
  })
})
