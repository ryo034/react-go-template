import { useIsDevCycleInitialized, withDevCycleProvider } from "@devcycle/react-client-sdk"
import { Toaster } from "shared-ui"
import { ContainerProvider } from "~/infrastructure/provider"
import { Router } from "~/infrastructure/route/router"

const devCycleApiKey = import.meta.env.VITE_DEV_CYCLE_API_KEY || "dummy"

const App = () => {
  const dvcReady = useIsDevCycleInitialized()
  if (!dvcReady) return <div />
  return (
    <ContainerProvider>
      <Router />
      <Toaster />
    </ContainerProvider>
  )
}

export const FeatureFlagProvider = withDevCycleProvider({ sdkKey: devCycleApiKey })(App)
