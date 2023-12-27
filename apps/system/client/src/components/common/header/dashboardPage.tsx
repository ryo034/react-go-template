import { cva } from "class-variance-authority"
import { HTMLAttributes } from "react"
import { cn } from "~/infrastructure/tailwindcss"

export interface Props extends HTMLAttributes<HTMLElement> {
  title: string
}

const variants = cva("px-8 pt-8", {
  variants: {}
})

export const DashboardPageHeader = ({ title, className }: Props) => {
  return (
    <header className={cn(variants({ className }))}>
      <div className="flex flex-col overflow-hidden max-w-full">
        <h1 className="text-2xl font-bold">{title}</h1>
      </div>
    </header>
  )
}
