import "@testing-library/jest-dom"
import { render, screen } from "@testing-library/react"
import userEvent from "@testing-library/user-event"
import React from "react"
import { describe, expect, it } from "vitest"
import { Input, inputBaseClass } from "../../src/components/ui/input"

describe("Input Component", () => {
  it("renders correctly with base classes", () => {
    render(<Input id="test-input" />)
    const expectClassString = `${inputBaseClass}`
    const input = screen.getByTestId("test-input")
    expect(input).toBeInTheDocument()
    expect(input).toHaveClass(...expectClassString.split(" "))
  })

  it("can be focused and typed into", async () => {
    render(<Input id="test-input" />)
    const input = screen.getByTestId("test-input") as HTMLInputElement
    await userEvent.click(input)
    await userEvent.type(input, "Hello, world!")
    expect(input.value).toBe("Hello, world!")
  })

  it("applies fullWidth class when fullWidth prop is true", () => {
    render(<Input id="test-input" fullWidth />)
    const input = screen.getByTestId("test-input")
    expect(input).toHaveClass("w-full")
  })

  it("applies additional classes passed via className prop", () => {
    render(<Input id="test-input" className="additional-class" />)
    const input = screen.getByTestId("test-input")
    expect(input).toHaveClass("additional-class")
  })

  it("applies attributes passed to the input", () => {
    render(<Input id="test-input" placeholder="Enter text" disabled />)
    const input = screen.getByTestId("test-input") as HTMLInputElement
    expect(input.placeholder).toBe("Enter text")
    expect(input.disabled).toBeTruthy()
  })
})
