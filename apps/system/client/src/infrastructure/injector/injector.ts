import ga4 from "react-ga4"
import { ApiErrorHandler, ErrorHandlingServiceMessageProvider } from "shared-network"
import { AuthDriver, ThemeDriver } from "~/driver"
import { FirebaseDriver } from "~/driver"
import { LoggerDriver } from "~/driver/logger/driver"
import { MeDriver } from "~/driver/me/driver"
import { WorkspaceDriver } from "~/driver/workspace/driver"
import { MessageProvider } from "~/infrastructure/error/message"
import { firebaseAuth } from "~/infrastructure/firebase"
import { ReactI18nextProvider, i18nKeys } from "~/infrastructure/i18n"
import { openapiFetchClient } from "~/infrastructure/openapi/client"
import { MeController, ThemeController } from "~/interface/controller"
import { AuthController } from "~/interface/controller/auth/controller"
import { WorkspaceController } from "~/interface/controller/workspace/controller"
import { InvitationGatewayAdapter, MeGateway, MeGatewayAdapter } from "~/interface/gateway"
import { AuthGateway, AuthGatewayAdapter } from "~/interface/gateway/auth"
import { UserGatewayAdapter } from "~/interface/gateway/user"
import { WorkspaceGatewayAdapter } from "~/interface/gateway/workspace"
import { WorkspaceGateway } from "~/interface/gateway/workspace/gateway"
import { MemberGatewayAdapter } from "~/interface/gateway/workspace/member"
import { AuthPresenter } from "~/interface/presenter/auth/presenter"
import { MePresenter } from "~/interface/presenter/me/presenter"
import { ThemePresenter } from "~/interface/presenter/theme/presenter"
import { WorkspacePresenter } from "~/interface/presenter/workspace/presenter"
import { authStore, invitationsStore, meStore, receivedInvitationStore, themeStore } from "~/store"
import { workspaceStore } from "~/store/workspace/store"
import { MeInteractor, ThemeInteractor } from "~/usecase"
import { AuthInteractor } from "~/usecase/auth"
import { WorkspaceInteractor } from "~/usecase/workspace"
import { SystemNetworkErrorInterpreter } from "../error"

const setupStore = () => {
  return {
    theme: themeStore,
    me: meStore,
    auth: authStore,
    workspace: workspaceStore,
    invitations: invitationsStore,
    receivedInvitation: receivedInvitationStore
  }
}

const store = setupStore()

const ls = localStorage

const apiErrorHandler = new ApiErrorHandler(new SystemNetworkErrorInterpreter().convertToSpecificError)

const setupDriver = () => {
  const firebase = new FirebaseDriver(firebaseAuth, apiErrorHandler)
  return {
    firebase,
    logger: new LoggerDriver(ga4, apiErrorHandler),
    theme: new ThemeDriver(ls),
    me: new MeDriver(openapiFetchClient, apiErrorHandler),
    auth: new AuthDriver(openapiFetchClient, apiErrorHandler),
    workspace: new WorkspaceDriver(openapiFetchClient, apiErrorHandler)
  }
}

const driver = setupDriver()

const setupGatewayAdapter = () => {
  const user = new UserGatewayAdapter()
  const member = new MemberGatewayAdapter(user)
  const workspace = new WorkspaceGatewayAdapter(member)
  const invitation = new InvitationGatewayAdapter()
  return {
    user,
    member,
    workspace,
    me: new MeGatewayAdapter(user, member, workspace, invitation),
    auth: new AuthGatewayAdapter(),
    invitation
  }
}

const gatewayAdapter = setupGatewayAdapter()

const setupGateway = () => {
  const me = new MeGateway(driver.me, driver.firebase, gatewayAdapter.me)
  return {
    me,
    auth: new AuthGateway(driver.auth, driver.firebase, gatewayAdapter.auth, gatewayAdapter.me),
    workspace: new WorkspaceGateway(
      driver.workspace,
      gatewayAdapter.workspace,
      gatewayAdapter.member,
      gatewayAdapter.invitation
    )
  }
}

const gateway = setupGateway()

const setupPresenter = () => {
  return {
    theme: new ThemePresenter(store.theme),
    me: new MePresenter(store.me),
    auth: new AuthPresenter(store.auth, store.receivedInvitation),
    workspace: new WorkspacePresenter(store.workspace, store.invitations)
  }
}

const presenter = setupPresenter()

const setupUseCase = () => {
  const me = new MeInteractor(gateway.me, presenter.me)
  return {
    theme: new ThemeInteractor(driver.theme, presenter.theme),
    me,
    auth: new AuthInteractor(gateway.auth, me, presenter.auth),
    workspace: new WorkspaceInteractor(gateway.workspace, me, presenter.workspace)
  }
}
const useCase = setupUseCase()

const setupController = () => {
  return {
    theme: new ThemeController(useCase.theme),
    auth: new AuthController(useCase.auth),
    me: new MeController(useCase.me),
    workspace: new WorkspaceController(useCase.workspace)
  }
}
const controller = setupController()

const i18n = new ReactI18nextProvider()
const systemErrorMessageProvider = new MessageProvider(i18n)

const errorMessageProvider = new ErrorHandlingServiceMessageProvider(
  {
    CannotConnectNetworkError: (_err) => i18n.translate(i18nKeys.network.cannotConnect),
    RequestTimeoutError: (_err) => i18n.translate(i18nKeys.network.requestTimeout),
    BadRequestError: (_err) => i18n.translate(i18nKeys.network.badRequest),
    ForbiddenError: (_err) => i18n.translate(i18nKeys.network.forbidden),
    AuthenticationError: (_err) => i18n.translate(i18nKeys.network.authentication),
    NotFoundError: (_err) => i18n.translate(i18nKeys.network.notFound),
    AlreadyExistError: (_err) => i18n.translate(i18nKeys.network.alreadyExist),
    InternalServerError: (_err) => i18n.translate(i18nKeys.network.internalServer)
  },
  systemErrorMessageProvider.translate
)

export const di = {
  driver,
  store,
  gateway,
  controller,
  i18n,
  errorMessageProvider
}

export type DI = typeof di
