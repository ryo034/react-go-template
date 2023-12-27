import { describe, expect, it } from "vitest"
import { Email } from "~/domain/shared/email"

describe("Email", () => {
  describe("create", () => {
    it("required", () => {
      expect(Email.create("").isErr).toBeTruthy()
    })
    it("max", () => {
      expect(Email.create("a".repeat(Email.max + 1)).isErr).toBeTruthy()
    })
    it("regex", () => {
      expect(Email.create("test").isErr).toBeTruthy()
      expect(Email.create("testtesttest.com").isErr).toBeTruthy()
    })
    it("OK", () => {
      expect(Email.create("test@example.com").isOk).toBeTruthy()
    })
  })
})
