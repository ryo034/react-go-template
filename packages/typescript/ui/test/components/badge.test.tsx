import "@testing-library/jest-dom"
import { render, screen } from "@testing-library/react"
import React from "react"
import { describe, expect, it } from "vitest"
import { Badge, badgeBaseClass } from "../../src/components/ui/badge"

describe("Badge Component", () => {
  it("renders default badge with correct classes", () => {
    render(<Badge>Default</Badge>)
    const expectClassString = `${badgeBaseClass} bg-primary text-primary-foreground`
    const badge = screen.getByText("Default")
    expect(badge).toBeInTheDocument()
    expect(badge).toHaveClass(...expectClassString.split(" "))
  })

  it.each([
    ["secondary", "bg-secondary", "text-secondary-foreground"],
    ["destructive", "bg-destructive", "text-destructive-foreground"],
    ["outline", "text-foreground"]
  ])("renders %s variant badge with correct classes", (variant, ...expectedClasses) => {
    render(<Badge variant={variant as any}>Test</Badge>)
    const badge = screen.getByText("Test")
    expectedClasses.forEach((expectedClass) => {
      expect(badge).toHaveClass(expectedClass)
    })
  })

  it("can accept and apply additional classes", () => {
    render(<Badge className="extra-class">Test</Badge>)
    const badge = screen.getByText("Test")
    expect(badge).toHaveClass("extra-class")
  })

  it("can accept and apply additional attributes", () => {
    render(<Badge data-testid="test-badge">Test</Badge>)
    const badge = screen.getByTestId("test-badge")
    expect(badge).toHaveAttribute("data-testid", "test-badge")
  })
})
