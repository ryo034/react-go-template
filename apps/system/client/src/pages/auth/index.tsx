import { useContext, useLayoutEffect, useRef, useState } from "react"
import { SubmitHandler } from "react-hook-form"
import { Link, useNavigate } from "react-router-dom"
import { AuthPageForm, LoginFormValues } from "~/components/auth/form"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"
import { routeMap } from "~/infrastructure/route/path"

export const authPageRoute = "/"

export const AuthPage = () => {
  const { store, controller, i18n, errorMessageProvider } = useContext(ContainerContext)
  const me = store.me((state) => state.me)
  const meRef = useRef(me)
  const authIsLoading = store.auth((state) => state.isLoading)
  const [errorMessage, setErrorMessage] = useState("")

  const navigate = useNavigate()

  useLayoutEffect(() => {
    store.me.subscribe((state) => {
      meRef.current = state.me
    })
  }, [])

  const onSubmit: SubmitHandler<LoginFormValues> = async (d) => {
    const res = await controller.auth.startWithEmail(d.email)
    console.log("res", res)
    if (res) {
      setErrorMessage(errorMessageProvider.translate(res))
      return
    }
    navigate(routeMap.verifyOtp)
  }

  const onClickGoogleLoginButton = async () => {
    console.log("onClickGoogleLoginButton")
  }

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-background/95">
      <section className="">
        <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
          <div className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
            <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
              <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
                {i18n.translate(i18nKeys.page.auth.title)}
              </h1>
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
        </div>
      </section>
    </div>
  )
}
