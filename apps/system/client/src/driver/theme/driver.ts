import type { ThemeType } from "~/store"

const key = "theme"
const darkThemeKey = "dark"
const lightThemeKey = "light"

export class ThemeDriver {
  constructor(private readonly client: Storage) {}

  get(): ThemeType | null {
    const res = this.client.getItem(key)
    if (res === null) {
      return null
    }
    return res === darkThemeKey ? "dark" : "light"
  }

  set(theme: ThemeType): void {
    this.client.setItem(key, theme === "dark" ? darkThemeKey : lightThemeKey)
  }
}
