export const ForbiddenPage = () => {
  return (
    <div className="flex flex-col items-center w-full min-h-[75vh] py-6">
      <div className="flex flex-col items-center justify-center space-y-4">
        <div className="space-y-2 text-center">
          <h1 className="text-3xl font-bold tracking-tighter">Uh oh. This page isn't available.</h1>
          <p className="text-gray-500 dark:text-gray-400">
            You may have mistyped the address or the page may have moved.
          </p>
        </div>
      </div>
    </div>
  )
}
