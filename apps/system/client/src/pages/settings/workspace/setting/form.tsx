"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { useEffect } from "react"
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
import { WorkspaceName } from "~/domain"

export type SettingsWorkspaceSettingFormValues = {
  name: string
  // subdomain: string
}

interface SettingsWorkspaceSettingFormProps {
  isUpdating: boolean
  onSubmit: SubmitHandler<SettingsWorkspaceSettingFormValues>
  defaultValues: {
    name: string
  }
}

export const SettingsWorkspaceSettingForm = ({
  isUpdating,
  onSubmit,
  defaultValues
}: SettingsWorkspaceSettingFormProps) => {
  const schema = z.object({
    name: z
      .string()
      .min(1, { message: "Name is required." })
      .max(WorkspaceName.max, {
        message: `Name must not be longer than ${WorkspaceName.max} characters.`
      })
  })

  const form = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
    defaultValues
  })

  useEffect(() => {
    form.reset(defaultValues)
  }, [defaultValues])

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          disabled={isUpdating}
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input placeholder="Workspace name" {...field} />
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
