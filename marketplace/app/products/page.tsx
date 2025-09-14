import { ProductsClient } from "@/components/products-client"
import { HydrationBoundary, dehydrate } from "@tanstack/react-query"
import getQueryClient from "@/lib/get-query-client"
import { apiService } from "@/lib/api"

export default async function ProductsPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined }
}) {
  const categories = typeof searchParams.categories === "string" ? searchParams.categories : undefined
  const name = typeof searchParams.name === "string" ? searchParams.name : undefined
  const minPrice = typeof searchParams.minPrice === "string" ? Number(searchParams.minPrice) : undefined
  const maxPrice = typeof searchParams.maxPrice === "string" ? Number(searchParams.maxPrice) : undefined
  const page = typeof searchParams.page === "string" ? Number(searchParams.page) : 1

  const queryClient = getQueryClient()

  await queryClient.prefetchQuery({
    queryKey: ["products", { categories, name, minPrice, maxPrice, page }],
    queryFn: () => apiService.getProducts({ 
        categories: categories?.split(","), 
        name, 
        minPrice, 
        maxPrice, 
        page 
    }),
  })

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="mb-8">
        <h1 className="text-4xl font-bold text-balance mb-4">Nossos Produtos</h1>
        <p className="text-muted-foreground text-lg">Descubra nossa seleção cuidadosa de produtos de alta qualidade</p>
      </div>

      <HydrationBoundary state={dehydrate(queryClient)}>
        <ProductsClient />
      </HydrationBoundary>
    </div>
  )
}

export const metadata = {
  title: "Produtos - Marketplace",
  description: "Explore nossa coleção completa de produtos de alta qualidade",
}
