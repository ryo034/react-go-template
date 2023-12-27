import { FC, useContext, useState } from "react"
import { useNavigate } from "react-router-dom"
import { Button, Card, CardContent, CardHeader, CardTitle, LoadingButton, Text, cn, toast } from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"
import { accountInitialPagePath } from "~/infrastructure/route/router"

export const ConfirmEmailPage: FC = () => {
  const { driver, controller } = useContext(ContainerContext)

  const [isCompleting, setIsCompleting] = useState<boolean>(false)

  const navigate = useNavigate()
  const currentUser = driver.firebase.currentUser

  const onClickConfirmButton = async () => {
    setIsCompleting(true)
    const res = await controller.me.verifyEmail()
    setIsCompleting(false)
    if (!res) {
      navigate(accountInitialPagePath)
      return
    }
    toast({
      variant: "destructive",
      title: "確認に失敗しました。",
      description: "メールアドレスを確認してください。"
    })
  }

  const onClickResendConfirmButton = async () => {
    const res = await controller.me.sendEmailVerification()
    if (!res) {
      toast({ title: "確認メールを再送しました。" })
      return
    }
    toast({
      variant: "destructive",
      title: "確認メールの再送に失敗しました。",
      description: "お手数ですが、時間をおいて再度お試しください。"
    })
  }

  const completeButton = () => {
    if (isCompleting) {
      return <LoadingButton fullWidth />
    }
    return (
      <Button fullWidth onClick={onClickConfirmButton}>
        確認が完了したらこちら
      </Button>
    )
  }

  const onClickLogoutButton = async () => {
    await controller.me.signOut()
  }

  return (
    <>
      <main className="min-h-screen bg-gray-50 dark:bg-background/95">
        <section className="">
          <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
            <Card className={cn("w-[380px]")}>
              <CardHeader>
                <CardTitle>メールアドレスを確認してください</CardTitle>
              </CardHeader>
              <CardContent className="grid gap-4">
                <Text>{currentUser?.email}に確認メールを送信しました。</Text>
                <Text>送られたメールに記載されているリンクをクリックしてください。</Text>
                {completeButton()}
                <Text>
                  確認メールが送られていませんか？
                  <Text asChild className="cursor-pointer font-bold" onClick={onClickResendConfirmButton}>
                    <span>確認メールを再送する</span>
                  </Text>
                </Text>

                <Text className="cursor-pointer underline">ヘルプセンター</Text>
              </CardContent>
            </Card>
            <Button className="mt-6" variant={"ghost"} onClick={onClickLogoutButton}>
              ログアウト
            </Button>
          </div>
        </section>
      </main>
    </>
  )
}
