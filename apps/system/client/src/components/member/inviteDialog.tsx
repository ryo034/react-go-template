import { PlusCircledIcon } from "@radix-ui/react-icons"
import { Plus, Trash } from "lucide-react"
import { useMemo } from "react"
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
  ScrollArea
} from "shared-ui"
import { Email } from "~/domain"
import { FormInputSection } from "../common/form/inputSection"
import { useInviteMembersFormMessage } from "./message"

export type InviteMembersFormValues = {
  members: [{ email: string; name: string }]
}

const inviteMembersFormId = "inviteMembersForm"

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

export const InviteMembersDialog = () => {
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

  useMemo(() => {
    append({ email: "", name: "" })
    setFocus("members.0.email")
  }, [])

  const onClickAddButton = () => {
    append({ email: "", name: "" })
  }

  const onClickDeleteButton = (index: number) => {
    remove(index)
  }

  const onSubmit = (data: InviteMembersFormValues) => {
    console.log(data)
  }

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="outline">
          <PlusCircledIcon className="mr-2 h-4 w-4" />
          {message.action.inviteMember}
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[580px] max-h-[580px]">
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
              <div className="w-full flex space-x-4 p-1" key={item.id}>
                <InviteMemberEmailInput
                  fieldArrayName={"members"}
                  fieldName="email"
                  index={index}
                  register={register}
                  errors={errors}
                />
                <Input id="name" placeholder={message.form.displayName.placeholder} />
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
          <Button type="submit" form={inviteMembersFormId}>
            {message.action.invite}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
