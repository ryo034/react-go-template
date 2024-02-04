import * as SeparatorPrimitive from "@radix-ui/react-separator"
import * as React from "react"

import { cn } from "~/infrastructure/tailwindcss"

const Separator = React.forwardRef<
  React.ElementRef<typeof SeparatorPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof SeparatorPrimitive.Root>
>(({ className, orientation = "horizontal", decorative = true, ...props }, ref) => (
  <SeparatorPrimitive.Root
    ref={ref}
    decorative={decorative}
    orientation={orientation}
    className={cn("shrink-0 bg-border", orientation === "horizontal" ? "h-[1px] w-full" : "h-full w-[1px]", className)}
    {...props}
  />
))
Separator.displayName = SeparatorPrimitive.Root.displayName

interface SeparatorWithTitleProps {
  title: string
}

const SeparatorWithTitle = React.forwardRef<HTMLDivElement, SeparatorWithTitleProps>(({ title, ...props }, ref) => {
  return (
    <div className="relative" {...props} ref={ref}>
      <div className="absolute inset-0 flex items-center">
        <span className="w-full border-t" />
      </div>
      <div className="relative flex justify-center text-xs uppercase">
        <span className="bg-background px-2 text-muted-foreground">{title}</span>
      </div>
    </div>
  )
})
SeparatorWithTitle.displayName = "SeparatorWithTitle"

export { Separator, SeparatorWithTitle }
