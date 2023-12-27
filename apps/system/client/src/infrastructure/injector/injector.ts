import { createPromiseClient } from "@connectrpc/connect"
import { createConnectTransport } from "@connectrpc/connect-web"
import ga4 from "react-ga4"
import { ThemeDriver } from "~/driver"
import { GoogleAnalyticsDriver } from "~/driver/analytics/ga/driver"
import { FirebaseDriver } from "~/driver/firebase"
import { MeDriver } from "~/driver/me/driver"
import { MeService } from "~/generated/schema/api/me/v1/me_connect"
import { MessageProvider } from "~/infrastructure/error/message"
import { firebaseAuth } from "~/infrastructure/firebase"
import { ReactI18nextProvider } from "~/infrastructure/i18n"
import { DriverAuthMiddleware } from "~/infrastructure/middleware/driver"
import { MeController, ThemeController } from "~/interface/controller"
import { MeGateway, MeGatewayAdapter } from "~/interface/gateway"
import { MePresenter } from "~/interface/presenter/me/presenter"
import { ThemePresenter } from "~/interface/presenter/theme/presenter"
import { meStore } from "~/store/me/store"
import { themeStore } from "~/store/theme/store"
import { MeInteractor, ThemeInteractor } from "~/usecase"

const driverAuthInterceptor = new DriverAuthMiddleware(firebaseAuth)

const transport = createConnectTransport({
  baseUrl: import.meta.env.VITE_API_BASE_URL,
  interceptors: [driverAuthInterceptor.intercept]
})

const setupStore = () => {
  return {
    me: meStore,
    theme: themeStore
  }
}

const store = setupStore()

const ls = localStorage
const setupDriver = () => {
  const firebase = new FirebaseDriver(firebaseAuth)
  return {
    firebase,
    ga: new GoogleAnalyticsDriver(ga4),
    me: new MeDriver(createPromiseClient(MeService, transport)),
    theme: new ThemeDriver(ls)
  }
}

const driver = setupDriver()

const setupGatewayAdapter = () => {
  return {
    me: new MeGatewayAdapter()
  }
}

const gatewayAdapter = setupGatewayAdapter()

const setupGateway = () => {
  return {
    me: new MeGateway(driver.me, driver.firebase, gatewayAdapter.me)
  }
}

const gateway = setupGateway()

const setupPresenter = () => {
  return {
    theme: new ThemePresenter(store.theme),
    me: new MePresenter(store.me)
  }
}

const presenter = setupPresenter()

const setupUseCase = () => {
  return {
    theme: new ThemeInteractor(driver.theme, presenter.theme),
    me: new MeInteractor(gateway.me, presenter.me)
  }
}
const useCase = setupUseCase()

const setupController = () => {
  return {
    theme: new ThemeController(useCase.theme),
    me: new MeController(useCase.me)
  }
}
const controller = setupController()

const i18n = new ReactI18nextProvider()
const errorMessageProvider = new MessageProvider(i18n)

export const di = {
  driver,
  store,
  gateway,
  controller,
  i18n,
  errorMessageProvider
}

export type DI = typeof di
