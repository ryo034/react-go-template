import { Settings } from "lucide-react"
import { useContext, useMemo } from "react"
import {
  Button,
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
  Dialog,
  DialogContent,
  DialogTrigger,
  Input,
  Label,
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger
} from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"

export const MembersSettingDialog = () => {
  const { controller } = useContext(ContainerContext)

  useMemo(() => {
    const fetchInvitations = async () => await controller.workspace.findAllInvitations()

    return async () => {
      fetchInvitations()
    }
  }, [])

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="ghost" data-testid="openMembersSettingDialogButton">
          <Settings className="h-4 w-4" />
        </Button>
      </DialogTrigger>
      <DialogContent className="max-h-[715px] wide-dialog" data-testid="inviteMembersDialog" hideCloseButton={true}>
        <Tabs defaultValue="account" className="w-full">
          <TabsList className="grid w-full grid-cols-2">
            <TabsTrigger value="account">メンバー</TabsTrigger>
            <TabsTrigger value="password">招待者</TabsTrigger>
          </TabsList>
          <TabsContent value="account">
            <Card>
              <CardHeader>
                <CardTitle>Account</CardTitle>
                <CardDescription>Make changes to your account here. Click save when you're done.</CardDescription>
              </CardHeader>
              <CardContent className="space-y-2">
                <div className="space-y-1">
                  <Label htmlFor="name">Name</Label>
                  <Input id="name" defaultValue="Pedro Duarte" />
                </div>
                <div className="space-y-1">
                  <Label htmlFor="username">Username</Label>
                  <Input id="username" defaultValue="@peduarte" />
                </div>
              </CardContent>
              <CardFooter>
                <Button>Save changes</Button>
              </CardFooter>
            </Card>
          </TabsContent>
          <TabsContent value="password">
            <p>aa</p>
          </TabsContent>
        </Tabs>
      </DialogContent>
    </Dialog>
  )
}
