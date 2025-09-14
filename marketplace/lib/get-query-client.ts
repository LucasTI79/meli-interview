import { QueryClient } from "@tanstack/react-query"
import { cache } from "react"

const getQueryClient = cache(
  () =>
    new QueryClient({
      defaultOptions: {
        queries: {
          staleTime: 60 * 1000, // 1 minute
          gcTime: 10 * 60 * 1000, // 10 minutes
        },
      },
    }),
)

export default getQueryClient
