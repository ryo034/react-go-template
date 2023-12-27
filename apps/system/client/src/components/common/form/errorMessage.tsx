interface Props {
  message: string
  dataTestId?: string
}

export const FormErrorMessage = ({ message, dataTestId }: Props) => {
  return (
    <span data-testid={dataTestId} className="text-red-600 text-xs">
      {message}
    </span>
  )
}
