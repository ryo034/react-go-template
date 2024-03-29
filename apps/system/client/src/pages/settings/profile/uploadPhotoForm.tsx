"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { X } from "lucide-react"
import { type SubmitHandler, useForm } from "react-hook-form"
import { Button, Input, Label, Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "shared-ui"
import { z } from "zod"
import { AccountAvatar } from "~/components/account/avatar"
import type { Me } from "~/domain"

export type SettingsProfileUploadPhotoFormValues = {
  photo: File
}

interface SettingsProfileUploadPhotoFormProps {
  onSubmit: SubmitHandler<SettingsProfileUploadPhotoFormValues>
  onClickRemoveProfilePhotoButton: () => void
  isUpdating: boolean
  me: Me
}

const IMAGE_TYPES = ["image/jpeg", "image/png"]

export const SettingsProfileUploadPhotoForm = ({
  onSubmit,
  onClickRemoveProfilePhotoButton,
  isUpdating = false,
  me
}: SettingsProfileUploadPhotoFormProps) => {
  const schema = z.object({
    photo: z
      .custom<FileList>()
      .refine((file) => file.length !== 0, { message: "必須です" })
      .transform((file) => file[0])
      .refine((file) => file.size < 500000, { message: "ファイルサイズは最大5MBです" })
      .refine((file) => IMAGE_TYPES.includes(file.type), {
        message: ".jpgもしくは.pngのみ可能です"
      })
  })

  type ProfileUploadPhotoFormValues = z.infer<typeof schema>

  const form = useForm<ProfileUploadPhotoFormValues>({
    resolver: zodResolver(schema),
    mode: "onChange"
  })

  const photoInputField = form.register("photo", {
    required: "必須です"
  })

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (isUpdating) return
    const file = e.target.files
    if (file === null) return
    onSubmit({ photo: file[0] })
  }

  if (me.self === undefined || me.member === undefined) return <></>

  return (
    <form id="settingsProfileUploadPhotoForm" onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
      <div className="w-24 h-24 relative">
        <Label htmlFor="file">
          <AccountAvatar
            alt="avatar"
            url={me.self.photo?.photoURL || ""}
            fallbackString={me.self.name?.firstTwoCharacters || ""}
            size="xl"
            className="cursor-pointer"
            data-testid="avatarOnUpdateProfileForm"
          />
        </Label>
        <Input
          id="file"
          type="file"
          accept="image/png, image/jpeg, image/jpg"
          className="hidden"
          {...photoInputField}
          onChange={onChange}
        />
        {me.self.hasPhoto && (
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger asChild>
                <Button
                  className="absolute right-[-4px] top-0 p-0 h-6 w-6 hover:bg-none"
                  variant="outline"
                  data-testid="removeProfilePhotoIconButton"
                  onClick={onClickRemoveProfilePhotoButton}
                >
                  <X className="w-4 h-4" color="gray" />
                </Button>
              </TooltipTrigger>
              <TooltipContent>
                <p>画像を削除</p>
              </TooltipContent>
            </Tooltip>
          </TooltipProvider>
        )}
      </div>
    </form>
  )
}
