export class AdapterError extends Error {
  className: string
  method: string
  constructor(className: string, method: string, message: string) {
    super(`Class:${className} - method:${method} - message:${message}`)
    Object.setPrototypeOf(this, new.target.prototype)
    this.name = Error.name
    this.className = className
    this.method = method
    if (Error.captureStackTrace !== undefined) {
      Error.captureStackTrace(this)
    }
  }
}
