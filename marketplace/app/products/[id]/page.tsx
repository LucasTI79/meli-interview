import { Suspense } from "react"
import { notFound } from "next/navigation"
import Link from "next/link"
import { ArrowLeft } from "lucide-react"
import { Button } from "@/components/ui/button"
import { ProductDetail } from "@/components/product-detail"
import { ProductGrid } from "@/components/product-grid"
import { LoadingSpinner } from "@/components/loading-spinner"
import { ErrorMessage } from "@/components/error-message"
import { mockApiService } from "@/lib/api"
import type { Product } from "@/types/product"

interface ProductPageProps {
  params: {
    id: string
  }
}

// Server component that fetches the product
async function ProductContent({ productId }: { productId: string }) {
  try {
    const product = await mockApiService.getProduct(productId)
    return <ProductDetail product={product} />
  } catch (error) {
    return (
      <ErrorMessage
        title="Produto não encontrado"
        message="O produto que você está procurando não existe ou foi removido."
      />
    )
  }
}

// Server component that fetches related products
async function RelatedProducts({ currentProductId, category }: { currentProductId: string; category?: string }) {
  try {
    const allProducts = await mockApiService.getProducts()
    // Filter out current product and get products from same category
    const relatedProducts = allProducts
      .filter((product) => product.id !== currentProductId && (!category || product.category === category))
      .slice(0, 4)

    if (relatedProducts.length === 0) {
      return null
    }

    return (
      <section className="mt-16">
        <h2 className="text-2xl font-bold mb-8">Produtos Relacionados</h2>
        <ProductGrid products={relatedProducts} />
      </section>
    )
  } catch (error) {
    return null
  }
}

export default async function ProductPage({ params }: ProductPageProps) {
  const productId = params.id

  // Get product for metadata and related products
  let product: Product | null = null
  try {
    product = await mockApiService.getProduct(productId)
  } catch (error) {
    // Product will be null, handled in ProductContent
  }

  return (
    <div className="container mx-auto px-4 py-8">
      {/* Breadcrumb */}
      <div className="mb-8">
        <Button variant="ghost" asChild className="mb-4">
          <Link href="/products">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Voltar aos Produtos
          </Link>
        </Button>
      </div>

      {/* Product Details */}
      <Suspense
        fallback={
          <div className="flex justify-center py-12">
            <LoadingSpinner size="lg" />
          </div>
        }
      >
        <ProductContent productId={productId} />
      </Suspense>

      {/* Related Products */}
      {product && (
        <Suspense fallback={<div className="mt-16 h-64" />}>
          <RelatedProducts currentProductId={productId} category={product.category} />
        </Suspense>
      )}
    </div>
  )
}

// Generate metadata for SEO
export async function generateMetadata({ params }: ProductPageProps) {
  const productId = params.id


  try {
    const product = await mockApiService.getProduct(productId)
    return {
      title: `${product.name} - Marketplace`,
      description: product.description,
      openGraph: {
        title: product.name,
        description: product.description,
        images: [product.image],
      },
    }
  } catch (error) {
    return {
      title: "Produto não encontrado",
    }
  }
}

// Generate static params for static generation (optional)
export async function generateStaticParams() {
  try {
    const products = await mockApiService.getProducts()
    return products.map((product) => ({
      id: product.id.toString(),
    }))
  } catch (error) {
    return []
  }
}
