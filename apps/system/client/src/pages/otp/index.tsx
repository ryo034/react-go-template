import { useContext, useLayoutEffect, useRef, useState } from "react"
import { SubmitHandler } from "react-hook-form"
import { useNavigate, useSearchParams } from "react-router-dom"
import { OtpFormValues, VerifyOTPPageForm } from "~/components/auth/otp/form"
import { Email } from "~/domain/shared"
import { useAuthenticator } from "~/infrastructure/hooks/auth"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"
import { routeMap } from "~/infrastructure/route/path"

export const verifyOtpPageRoute = "/verify-otp"

export const VerifyOtpPage = () => {
  const { store, controller, i18n, errorMessageProvider } = useContext(ContainerContext)
  const [searchParams] = useSearchParams()
  const [errorMessage, setErrorMessage] = useState("")
  const { nextNavigate } = useAuthenticator()
  const navigate = useNavigate()
  const email = searchParams.get("email")
  const ema = Email.create(email ?? "")
  const me = store.me((state) => state.me)
  const meRef = useRef(me)

  useLayoutEffect(() => {
    store.me.subscribe((state) => {
      meRef.current = state.me
    })

    if (ema.isErr) {
      navigate({ pathname: routeMap.auth })
    }
  }, [email])

  const onSubmit: SubmitHandler<OtpFormValues> = async (d) => {
    if (ema.isErr) {
      setErrorMessage(errorMessageProvider.resolve(ema.error))
      return
    }
    const opt = `${d.otpInput1}${d.otpInput2}${d.otpInput3}${d.otpInput4}${d.otpInput5}${d.otpInput6}`
    const err = await controller.auth.verifyOtp(ema.value, opt)
    if (err) {
      setErrorMessage(errorMessageProvider.resolve(err))
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
      navigate(routeMap.receivedInvitationsPageRoute)
      return
    }
    if (meRef.current.hasNotWorkspace) {
      navigate(routeMap.onboardingSettingWorkspace)
      return
    }
    navigate(routeMap.home)
  }

  return (
    <div className="flex justify-center items-center min-h-screen">
      <section className="w-full max-w-md mx-auto py-12 px-4 sm:px-6 lg:px-8">
        <div className="space-y-6">
          <h2 className="text-center text-3xl font-extrabold text-gray-900">
            {i18n.translate(i18nKeys.action.enter, { field: i18n.translate(i18nKeys.word.otp) })}
          </h2>
          <p className="text-center text-sm text-muted-foreground">
            {i18n.translate(i18nKeys.page.verifyOtp.enterOtpMessage)}
            <span>{email}</span>
          </p>
          <VerifyOTPPageForm onSubmit={onSubmit} errorMessage={errorMessage} />
        </div>
      </section>
    </div>
  )
}
