import "@testing-library/jest-dom"
import { render, screen } from "@testing-library/react"
import React from "react"
import { describe, expect, it } from "vitest"
import { Alert, AlertDescription, AlertTitle, baseAlertClass } from "../../src/components/ui/alert"

describe("Alert Component", () => {
  it("renders with default classes", () => {
    render(<Alert>Test</Alert>)
    const alert = screen.getByText("Test")
    const _expectClassString = `${baseAlertClass} border-gray-300 text-gray-800 bg-gray-50 [&>svg]:text-gray-800 dark:text-white dark:border-gray-400 dark:bg-transparent dark:[&>svg]:text-white`
    expect(alert).toBeInTheDocument()
    expect(alert).toHaveClass(...baseAlertClass.split(" "))
  })

  it("renders AlertTitle with correct classes", () => {
    render(<AlertTitle>Test</AlertTitle>)
    const title = screen.getByText("Test")
    expect(title).toHaveClass("mb-1", "font-medium", "leading-none", "tracking-tight")
  })

  it("renders AlertDescription with correct classes", () => {
    render(<AlertDescription>Test</AlertDescription>)
    const description = screen.getByText("Test")
    expect(description).toHaveClass("text-sm", "[&_p]:leading-relaxed")
  })
})
