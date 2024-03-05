"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { SubmitHandler, useForm } from "react-hook-form"
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
  Button,
  Input,
  Label,
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger
} from "shared-ui"
import { z } from "zod"
import { Me } from "~/domain"

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
    <form id="settingsProfileUploadPhotoForm" onSubmit={form.handleSubmit(onSubmit)}>
      <TooltipProvider>
        <Tooltip>
          <TooltipTrigger asChild>
            <div className="w-24 h-24">
              <Label htmlFor="file">
                <Avatar className="w-24 h-24 rounded-[36px] cursor-pointer">
                  <AvatarImage src={me.self.photo?.photoURL} alt={"aa"} />
                  <AvatarFallback className="rounded-[36px]">
                    {me.member?.profile.displayName?.firstTwoCharacters}
                  </AvatarFallback>
                </Avatar>
              </Label>
              <Input
                id="file"
                type="file"
                accept="image/*"
                className="hidden"
                {...photoInputField}
                onChange={onChange}
              />
            </div>
          </TooltipTrigger>
          <TooltipContent className="cursor-pointer p-0">
            <Button variant="ghost" onClick={onClickRemoveProfilePhotoButton}>
              <span className="text-red-600">画像を削除</span>
            </Button>
          </TooltipContent>
        </Tooltip>
      </TooltipProvider>
    </form>
  )

  // return (
  //   <Form {...form}>
  //     <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
  //       <FormField
  //         control={form.control}
  //         name="photo"
  //         render={({ field }) => (
  //           <FormItem>
  //             <FormControl>
  //               <>
  //                 <Label htmlFor="file">
  //                   <Avatar className="mr-auto w-24 h-24 rounded-[36px] cursor-pointer" onClick={onAvatarClick}>
  //                     <AvatarImage src={me.self.photo?.photoURL} alt={"aa"} />
  //                     <AvatarFallback className="rounded-[36px]">
  //                       {me.member?.profile.displayName?.firstTwoCharacters}
  //                     </AvatarFallback>
  //                   </Avatar>
  //                 </Label>
  //                 <Input
  //                   id="file"
  //                   type="file"
  //                   accept="image/*"
  //                   className="hidden"
  //                   {...field}
  //                   value={field.value as any}
  //                 />
  //               </>
  //             </FormControl>
  //           </FormItem>
  //         )}
  //       />
  //       <div className="flex space-x-4">
  //         <Button type="button" onClick={onClickUpdateProfilePhotoButton} data-testid="updateProfilePhoto">
  //           {me.self.hasPhoto ? "画像を変更" : "画像を追加"}
  //         </Button>
  //         <Button
  //           type="button"
  //           variant="ghost"
  //           onClick={onClickRemoveProfilePhotoButton}
  //           data-testid="removeProfilePhoto"
  //         >
  //           <span className="text-red-600">画像を削除</span>
  //         </Button>
  //       </div>
  //     </form>
  //   </Form>
  // )
}
