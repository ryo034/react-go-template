export class EnvHandler {
  static isLocal(): boolean {
    return process.env.NODE_ENV === "development"
  }

  static isNotLocal(): boolean {
    return !EnvHandler.isLocal()
  }
}
