import { render, screen } from "@testing-library/react"
import React from "react"
import { describe, expect, it } from "vitest"
import { Text, textBaseClass } from "../../src/components/ui/text"

describe("Text Component", () => {
  it("renders default text with correct classes", () => {
    render(<Text>Text</Text>)
    const expectClassString = textBaseClass
    const text = screen.getByText("Text")
    expect(text).toBeInTheDocument()
    for (const c of expectClassString.split(" ")) {
      expect(text).toHaveClass(c)
    }
  })

  it("applies fullWidth class when fullWidth prop is true", () => {
    render(<Text fullWidth>Text</Text>)
    const text = screen.getByText("Text")
    expect(text).toHaveClass("w-full")
  })

  it("applies additional classes passed via className prop", () => {
    render(<Text className="additional-class">Text</Text>)
    const text = screen.getByText("Text")
    expect(text).toHaveClass("additional-class")
  })

  it("renders as a child element when asChild prop is true", () => {
    render(
      <Text asChild>
        <div>Text</div>
      </Text>
    )
    const text = screen.getByText("Text")
    expect(text.tagName).toBe("DIV")
  })
})
