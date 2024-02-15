import { Skeleton } from "shared-ui"

interface Props {
  count: number
}

export const MemberCardListLoading = ({ count }: Props) => {
  return (
    <>
      {Array.from({ length: count }).map((_, i) => {
        return (
          <div className="flex flex-wrap ju text-left" key={i + 1}>
            <Skeleton className="mr-auto w-24 h-24 rounded-[36px]" />
            <div className="mt-4 w-full">
              <Skeleton className="h-4 w-full mb-4" />
              <Skeleton className="h-2 w-[80px]" />
            </div>
          </div>
        )
      })}
    </>
  )
}
