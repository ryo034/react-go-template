import "@testing-library/jest-dom"
import { act } from "@testing-library/react"

let listener: ((rect: any) => void) | undefined = undefined
global.ResizeObserver = class ResizeObserver {
  constructor(ls: any) {
    listener = ls
  }
  observe() {}
  unobserve() {}
  disconnect() {}
}

// then in your test case you can trigger the callback with the following code
act(() => {
  listener?.([
    {
      target: {
        clientWidth: 100,
        scrollWidth: 200,
        clientHeight: 100,
        scrollHeight: 200
      }
    }
  ])
})
