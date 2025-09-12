import { Suspense } from "react"
import { ProductsClient } from "@/components/products-client"
import { LoadingSpinner } from "@/components/loading-spinner"
import { ErrorMessage } from "@/components/error-message"
import { mockApiService } from "@/lib/api"

// Server component that fetches products
async function ProductsContent() {
  try {
    const products = await mockApiService.getProducts()
    return <ProductsClient products={products} />
  } catch (error) {
    return (
      <ErrorMessage
        title="Erro ao carregar produtos"
        message="Não foi possível carregar os produtos. Tente novamente mais tarde."
      />
    )
  }
}

export default function ProductsPage() {
  return (
    <div className="container mx-auto px-4 py-8">
      {/* Page Header */}
      <div className="mb-8">
        <h1 className="text-4xl font-bold text-balance mb-4">Nossos Produtos</h1>
        <p className="text-muted-foreground text-lg">Descubra nossa seleção cuidadosa de produtos de alta qualidade</p>
      </div>

      {/* Products with Interactive Features */}
      <Suspense
        fallback={
          <div className="flex justify-center py-12">
            <LoadingSpinner size="lg" />
          </div>
        }
      >
        <ProductsContent />
      </Suspense>
    </div>
  )
}

export const metadata = {
  title: "Produtos - Marketplace",
  description: "Explore nossa coleção completa de produtos de alta qualidade",
}
