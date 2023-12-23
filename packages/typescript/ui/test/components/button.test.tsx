import "@testing-library/jest-dom"
import { render, screen } from "@testing-library/react"
import React from "react"
import { describe, expect, it } from "vitest"
import { Button, buttonBaseClass } from "../../src/components/ui/button"

describe("Button component", () => {
  it("renders button with the correct text", async () => {
    render(<Button>Test</Button>)
    const buttonElement = screen.getByText("Test")
    expect(buttonElement).toBeInTheDocument()
  })

  it("renders full width button", () => {
    render(<Button fullWidth>Test Button</Button>)
    const button = screen.getByText("Test Button")
    expect(button).toHaveClass("w-full")
  })

  it("button has the correct default class", () => {
    render(<Button variant="default">Test</Button>)
    const expectClassString = `${buttonBaseClass} bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2 flex items-center justify-center`
    const buttonElement = screen.getByText("Test")
    expect(buttonElement).toHaveClass(...expectClassString.split(" "))
  })

  it("renders button as a child element", () => {
    render(
      <Button asChild>
        <div>Test</div>
      </Button>
    )
    const button = screen.getByText("Test")
    expect(button.tagName).toBe("DIV")
  })

  it("can accept and apply additional attributes", () => {
    render(<Button data-testid="test-id">Test</Button>)
    const buttonElement = screen.getByText("Test")
    expect(buttonElement).toHaveAttribute("data-testid", "test-id")
  })
})
