import { PlusCircledIcon } from "@radix-ui/react-icons"
import { Plus, Trash } from "lucide-react"
import { useContext, useEffect, useState } from "react"
import { useFieldArray, useForm } from "react-hook-form"
import {
  Button,
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  Input,
  LoadingButton,
  ScrollArea,
  useToast
} from "shared-ui"
import { Email } from "~/domain/shared"
import { ContainerContext } from "~/infrastructure/injector/context"
import { useInviteMembersFormMessage } from "./message"

export type InviteMembersFormValues = {
  members: [{ email: string; name: string }]
}

export const inviteMembersFormId = "inviteMembersForm"

interface InviteMemberEmailInputProps {
  register: any
  fieldArrayName: string
  fieldName: string
  index: number
  errors: any
}

const InviteMemberEmailInput = ({
  fieldArrayName,
  fieldName,
  index,
  errors,
  register
}: InviteMemberEmailInputProps) => {
  const message = useInviteMembersFormMessage()
  const fn = `${fieldArrayName}.${index}.${fieldName}`
  const errorMessage = errors?.[fieldArrayName]?.[index]?.[fieldName]?.message
  return (
    <Input
      className={errorMessage ? "border-red-500 border-2" : ""}
      fullWidth
      autoComplete={"email"}
      id={fn}
      type={"email"}
      placeholder={message.form.email.placeholder}
      reactHookForm={register(fn, {
        required: message.form.email.required,
        pattern: {
          value: Email.pattern,
          message: message.form.email.regex
        }
      })}
    />
  )
}

interface InviteMemberNameInputProps {
  register: any
  fieldArrayName: string
  fieldName: string
  index: number
}

const InviteMemberNameInput = ({ fieldArrayName, fieldName, index, register }: InviteMemberNameInputProps) => {
  const message = useInviteMembersFormMessage()
  const fn = `${fieldArrayName}.${index}.${fieldName}`
  return (
    <Input
      autoComplete={"name"}
      id={fn}
      type={"text"}
      placeholder={message.form.displayName.placeholder}
      reactHookForm={register(fn)}
    />
  )
}

export const InviteMembersDialog = () => {
  const { controller } = useContext(ContainerContext)
  const [open, setOpen] = useState(false)
  const [inviting, setInviting] = useState(false)
  const { toast } = useToast()

  const {
    register,
    control,
    handleSubmit,
    setFocus,
    formState: { errors }
  } = useForm<InviteMembersFormValues>()

  const { fields, append, remove } = useFieldArray({
    control,
    name: "members"
  })

  const message = useInviteMembersFormMessage()

  useEffect(() => {
    remove()
    append({ email: "", name: "" })
    setFocus("members.0.email")
  }, [open, setFocus, append, remove])

  const onClickAddButton = () => {
    append({ email: "", name: "" })
  }

  const onClickDeleteButton = (index: number) => {
    remove(index)
  }

  const close = () => {
    setInviting(false)
    setOpen(false)
  }

  const onSubmit = async (data: InviteMembersFormValues) => {
    setInviting(true)
    const res = await controller.workspace.inviteMembers({ invitees: data.members })
    toast({ title: res ? message.action.failedInvite : message.action.successInvite })
    close()
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant="outline" data-testid="inviteMembersButton">
          <PlusCircledIcon className="mr-2 h-4 w-4" />
          {message.action.inviteMember}
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[580px] max-h-[580px]" data-testid="inviteMembersDialog">
        <DialogHeader>
          <DialogTitle className="text-center">{message.action.inviteMember}</DialogTitle>
        </DialogHeader>
        <ScrollArea className="h-72 w-full rounded-md">
          <form
            id={inviteMembersFormId}
            data-testid={inviteMembersFormId}
            className="flex flex-wrap justify-center space-y-4"
            onSubmit={handleSubmit(onSubmit)}
          >
            {fields.map((item, index) => (
              <div className="w-full flex space-x-4 p-1" key={item.id} data-testid={`inviteMemberFiledRow-${index}`}>
                <InviteMemberEmailInput
                  fieldArrayName={"members"}
                  fieldName="email"
                  index={index}
                  register={register}
                  errors={errors}
                />
                <InviteMemberNameInput fieldArrayName={"members"} fieldName="name" index={index} register={register} />
                <Button variant="ghost" type="button" onClick={() => onClickDeleteButton(index)}>
                  <Trash color="red" className="h-4 w-4" />
                </Button>
              </div>
            ))}
            <Button variant="ghost" type="button" onClick={onClickAddButton}>
              <Plus className="mr-2 h-4 w-4" />
              {message.action.add}
            </Button>
          </form>
        </ScrollArea>
        <DialogFooter>
          {inviting ? (
            <LoadingButton text={`${message.action.inviting}...`} />
          ) : (
            <Button type="submit" form={inviteMembersFormId} disabled={fields.length === 0}>
              {message.action.invite}
            </Button>
          )}
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
