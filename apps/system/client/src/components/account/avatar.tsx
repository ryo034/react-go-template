import { HTMLAttributes } from "react"
import { Avatar, AvatarFallback, AvatarImage } from "shared-ui"

const AccountAvatarSizeType = {
  sm: "sm",
  // md: "md",
  // lg: "lg",
  xl: "xl"
} as const

interface Props extends HTMLAttributes<HTMLElement> {
  size: keyof typeof AccountAvatarSizeType
  className?: string
  testId?: string
  url: string
  alt: string
  fallbackString: string
}

const accountAvatarSizeClassMap = {
  sm: "h-8 w-8 rounded-[12px]",
  // md: "h-12 w-12 rounded-[16px]",
  // lg: "h-16 w-16 rounded-[24px]",
  xl: "h-24 w-24 rounded-[36px]"
}

export const AccountAvatar = ({ size, className, url, alt, fallbackString, ...props }: Props) => {
  const addClass = `${accountAvatarSizeClassMap[size]} ${className}`
  return (
    <Avatar className={addClass} {...props}>
      <AvatarImage src={url} alt={alt} />
      <AvatarFallback data-testid="avatarFallback">{fallbackString}</AvatarFallback>
    </Avatar>
  )
}
