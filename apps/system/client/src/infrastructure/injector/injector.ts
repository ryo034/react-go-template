import ga4 from "react-ga4"
import { ApiErrorHandler } from "shared-network"
import { AuthDriver, ThemeDriver } from "~/driver"
import { FirebaseDriver } from "~/driver"
import { GoogleAnalyticsDriver } from "~/driver/analytics/ga/driver"
import { MeDriver } from "~/driver/me/driver"
import { MessageProvider } from "~/infrastructure/error/message"
import { firebaseAuth } from "~/infrastructure/firebase"
import { ReactI18nextProvider } from "~/infrastructure/i18n"
import { openapiFetchClient } from "~/infrastructure/openapi/client"
import "~/infrastructure/openapi/interceptor"
import { MeController, ThemeController } from "~/interface/controller"
import { AuthController } from "~/interface/controller/auth/controller"
import { MeGateway, MeGatewayAdapter } from "~/interface/gateway"
import { AuthGateway, AuthGatewayAdapter } from "~/interface/gateway/auth"
import { UserGatewayAdapter } from "~/interface/gateway/user"
import { WorkspaceGatewayAdapter } from "~/interface/gateway/workspace"
import { MemberGatewayAdapter } from "~/interface/gateway/workspace/member"
import { AuthPresenter } from "~/interface/presenter/auth/presenter"
import { MePresenter } from "~/interface/presenter/me/presenter"
import { ThemePresenter } from "~/interface/presenter/theme/presenter"
import { authStore, meStore, themeStore } from "~/store"
import { MeInteractor, ThemeInteractor } from "~/usecase"
import { AuthInteractor } from "~/usecase/auth"
import { SystemNetworkErrorInterpreter } from "../error"

const setupStore = () => {
  return {
    theme: themeStore,
    me: meStore,
    auth: authStore
  }
}

const store = setupStore()

const ls = localStorage

const apiErrorHandler = new ApiErrorHandler(new SystemNetworkErrorInterpreter().convertToSpecificError)

const setupDriver = () => {
  const firebase = new FirebaseDriver(firebaseAuth, apiErrorHandler)
  return {
    firebase,
    ga: new GoogleAnalyticsDriver(ga4, apiErrorHandler),
    theme: new ThemeDriver(ls),
    me: new MeDriver(openapiFetchClient, apiErrorHandler),
    auth: new AuthDriver(openapiFetchClient, apiErrorHandler)
  }
}

const driver = setupDriver()

const setupGatewayAdapter = () => {
  const user = new UserGatewayAdapter()
  const member = new MemberGatewayAdapter(user)
  const workspace = new WorkspaceGatewayAdapter()
  return {
    user,
    member,
    workspace,
    me: new MeGatewayAdapter(user, member, workspace),
    auth: new AuthGatewayAdapter()
  }
}

const gatewayAdapter = setupGatewayAdapter()

const setupGateway = () => {
  return {
    me: new MeGateway(driver.me, driver.firebase, gatewayAdapter.me),
    auth: new AuthGateway(driver.auth, driver.firebase, gatewayAdapter.auth)
  }
}

const gateway = setupGateway()

const setupPresenter = () => {
  return {
    theme: new ThemePresenter(store.theme),
    me: new MePresenter(store.me),
    auth: new AuthPresenter(store.auth)
  }
}

const presenter = setupPresenter()

const setupUseCase = () => {
  const me = new MeInteractor(gateway.me, presenter.me)
  return {
    theme: new ThemeInteractor(driver.theme, presenter.theme),
    me,
    auth: new AuthInteractor(gateway.auth, me, presenter.auth)
  }
}
const useCase = setupUseCase()

const setupController = () => {
  return {
    theme: new ThemeController(useCase.theme),
    me: new MeController(useCase.me),
    auth: new AuthController(useCase.auth)
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
