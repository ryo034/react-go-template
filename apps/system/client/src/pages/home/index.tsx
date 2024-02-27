import { Card } from "shared-ui"

export const homePageRoute = "/home"

export function HomePage() {
  return (
    <div className="flex justify-center items-center min-h-screen" data-testid="homePage">
      <Card>🎉🎉🎉Home🎉🎉🎉</Card>
    </div>
  )
}
