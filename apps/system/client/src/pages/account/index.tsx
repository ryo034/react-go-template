import { LogOut } from "lucide-react"
import { useContext } from "react"
import { Button, Card, CardContent, CardFooter, CardHeader, CardTitle, useToast } from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"
import { cn } from "~/infrastructure/tailwindcss"

export const AccountPage = () => {
  const { controller, store } = useContext(ContainerContext)
  const me = store.me((state) => state.me)

  const { toast } = useToast()

  const onClick = async () => {
    const res = await controller.me.signOut()
    if (!res) {
      toast({ title: "ログアウトしました👋" })
      return
    }
  }

  if (me === null) {
    return <></>
  }

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-background/95">
      <section className="">
        <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
          <Card className={cn("w-[380px]")}>
            <CardHeader>
              <CardTitle>Congratulations🎉🎉🎉</CardTitle>
            </CardHeader>
            <CardContent className="grid gap-4">
              <h2>Hello!! {me.user.name}</h2>
            </CardContent>
            <CardFooter>
              <Button className="w-full" onClick={onClick}>
                <LogOut className="mr-2 h-4 w-4" /> Logout
              </Button>
            </CardFooter>
          </Card>
        </div>
      </section>
    </div>
  )
}
