import { Skeleton } from "shared-ui"

interface Props {
  count: number
}

const generateUUID = () => {
  return (Math.random() + 1).toString(36).substring(7)
}

const generateDummyList = (count: number) => {
  return Array.from({ length: count }).map((_) => generateUUID())
}

export const MemberCardListLoading = ({ count }: Props) => {
  return (
    <>
      {generateDummyList(count).map((id) => {
        return (
          <div className="flex flex-wrap ju text-left" key={id}>
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
