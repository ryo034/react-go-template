import { getRedirectResult } from "firebase/auth"
import { useContext, useLayoutEffect, useRef, useState } from "react"
import type { SubmitHandler } from "react-hook-form"
import { Link, createSearchParams, useNavigate } from "react-router-dom"
import { AuthPageForm, type LoginFormValues } from "~/components/auth/form"
import { useErrorMessageHandler } from "~/infrastructure/hooks/error"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"
import { routeMap } from "~/infrastructure/route/path"
import { serviceName } from "~/infrastructure/shared"

export const authPageRoute = "/"

export const AuthPage = () => {
  const { store, controller, i18n, driver } = useContext(ContainerContext)
  const { handleErrorMessage } = useErrorMessageHandler()
  const me = store.me((state) => state.me)
  const meRef = useRef(me)
  const authIsLoading = store.auth((state) => state.isLoading)
  const [errorMessage, setErrorMessage] = useState("")

  const navigate = useNavigate()

  useLayoutEffect(() => {
    store.me.subscribe((state) => {
      meRef.current = state.me
    })

    const handleRedirectResult = async () => {
      const result = await getRedirectResult(driver.firebase.getClient)
      if (result === null) {
        return
      }
      const err = await controller.auth.createByOAuth()
      if (err) {
        setErrorMessage(handleErrorMessage(err))
        return
      }

      if (meRef.current === null) {
        return
      }
      if (meRef.current.self.hasNotName) {
        navigate(routeMap.onboardingSettingName)
        return
      }
      if (meRef.current.hasReceivedInvitations) {
        navigate(routeMap.receivedInvitations)
        return
      }
      if (meRef.current.hasNotWorkspace) {
        navigate(routeMap.onboardingSettingWorkspace)
        return
      }
      navigate(routeMap.home)
    }
    handleRedirectResult()
  }, [])

  // if logged in, call useEffect getRedirectResult and navigate to home
  const onClickGoogleLoginButton = async () => {
    await driver.firebase.startWithGoogle()
  }

  const onSubmit: SubmitHandler<LoginFormValues> = async (d) => {
    const res = await controller.auth.startWithEmail(d.email)
    if (res) {
      setErrorMessage(handleErrorMessage(res))
      return
    }
    navigate({ pathname: routeMap.verifyOtp, search: createSearchParams({ email: d.email }).toString() })
  }

  return (
    <div className="container flex relative h-full flex-col items-center justify-center md:grid lg:max-w-none lg:grid-cols-2 lg:px-0">
      <div className="relative hidden h-full flex-col bg-muted p-10 text-white lg:flex dark:border-r">
        <div className="absolute inset-0 bg-zinc-900" />
        <div className="relative z-20 flex items-center text-lg font-medium">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
            className="mr-2 h-6 w-6"
          >
            <path d="M15 6v12a3 3 0 1 0 3-3H6a3 3 0 1 0 3 3V6a3 3 0 1 0-3 3h12a3 3 0 1 0-3-3" />
          </svg>
          {serviceName}
        </div>
        <div className="relative z-20 mt-auto">
          <blockquote className="space-y-2">
            <p className="text-lg">
              &ldquo;This library has saved me countless hours of work and helped me deliver stunning designs to my
              clients faster than ever before.&rdquo;
            </p>
            <footer className="text-sm">Sofia Davis</footer>
          </blockquote>
        </div>
      </div>
      <div className="lg:p-8 mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[420px]">
        <div className="flex flex-col space-y-2 text-center">
          <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
            {i18n.translate(i18nKeys.page.auth.title)}
          </h1>
          <p className="text-sm text-muted-foreground">Enter your email below to create your account</p>
        </div>
        <AuthPageForm
          onSubmit={onSubmit}
          onClickGoogleLoginButton={onClickGoogleLoginButton}
          errorMessage={errorMessage}
          isLoading={authIsLoading}
        />
        <p className="px-8 text-center text-sm text-muted-foreground">
          By clicking continue, you agree to our{" "}
          <Link to="/terms" className="underline underline-offset-4 hover:text-primary">
            Terms of Service
          </Link>{" "}
          and{" "}
          <Link to="/privacy" className="underline underline-offset-4 hover:text-primary">
            Privacy Policy
          </Link>
          .
        </p>
      </div>
    </div>
  )
}
