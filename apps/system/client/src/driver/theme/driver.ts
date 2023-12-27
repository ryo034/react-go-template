const key = "theme"
const darkThemeKey = "dark"
const lightThemeKey = "light"

export class ThemeDriver {
  constructor(private readonly client: Storage) {}

  get(): boolean {
    return this.client.getItem(key) === "dark"
  }

  set(isDark: boolean): void {
    this.client.setItem(key, isDark ? darkThemeKey : lightThemeKey)
  }
}
