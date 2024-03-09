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
  LoadingButton
} from "shared-ui"
import { z } from "zod"
import { AccountFullName } from "~/domain"

export type SettingsAccountFormValues = {
  name: string
}

interface SettingsAccountFormProps {
  isUpdating: boolean
  onSubmit: SubmitHandler<SettingsAccountFormValues>
  defaultValues: {
    name: string
  }
}

export const SettingsAccountForm = ({ isUpdating, onSubmit, defaultValues }: SettingsAccountFormProps) => {
  const accountFormSchema = z.object({
    name: z
      .string()
      .max(AccountFullName.max, {
        message: "Name must not be longer than 30 characters."
      })
      .regex(AccountFullName.pattern, {
        message: "Name must contain only letters, numbers, and spaces."
      })
  })

  const form = useForm<z.infer<typeof accountFormSchema>>({
    resolver: zodResolver(accountFormSchema),
    defaultValues
  })

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input placeholder="Your name" {...field} />
              </FormControl>
              <FormDescription>This is the name that will be displayed on your profile and in emails.</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        {isUpdating ? <LoadingButton /> : <Button type="submit">Update</Button>}
      </form>
    </Form>
  )
}
