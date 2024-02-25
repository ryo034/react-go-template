"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { SubmitHandler, useForm } from "react-hook-form"
import {
  Button,
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
  Input,
  LoadingButton,
  Textarea,
} from "shared-ui"
import { z } from "zod"
import { MemberDisplayName } from "~/domain"
import { MemberBio } from "~/domain/workspace/member/bio"

export type SettingsProfileFormValues = {
  displayName: string
  bio: string
}

interface SettingsProfileFormProps {
  isUpdating: boolean
  onSubmit: SubmitHandler<SettingsProfileFormValues>
  defaultValues: {
    displayName: string
    bio: string
  }
}

export const SettingsProfileForm = ({ defaultValues, onSubmit, isUpdating = false }: SettingsProfileFormProps) => {
  const profileFormSchema = z.object({
    displayName: z
      .string()
      .min(MemberDisplayName.min, {
        message: "Username must be at least 2 characters."
      })
      .max(MemberDisplayName.max, {
        message: "Username must not be longer than 30 characters."
      }),
    bio: z.string().max(MemberBio.max)
  })

  type ProfileFormValues = z.infer<typeof profileFormSchema>

  const form = useForm<ProfileFormValues>({
    resolver: zodResolver(profileFormSchema),
    defaultValues,
    mode: "onChange"
  })

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="displayName"
          render={({ field }) => (
            <FormItem>
              <FormLabel>DisplayName</FormLabel>
              <FormControl>
                <Input placeholder="display name" {...field} />
              </FormControl>
              <FormDescription>
                This is your public display name. It can be your real name or a pseudonym. You can only change this once
                every 30 days.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="bio"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Bio</FormLabel>
              <FormControl>
                <Textarea placeholder="Tell us a little bit about yourself" className="resize-none h-32" {...field} />
              </FormControl>
              <FormDescription>
                You can <span>@mention</span> other users and organizations to link to them.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        {isUpdating ? <LoadingButton /> : <Button type="submit">Update profile</Button>}
      </form>
    </Form>
  )
}
