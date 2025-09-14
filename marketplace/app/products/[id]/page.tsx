import { Suspense } from "react"
import { notFound } from "next/navigation"
import Link from "next/link"
import { ArrowLeft } from "lucide-react"
import { Button } from "@/components/ui/button"
import { ProductDetail } from "@/components/product-detail"
import { ProductGrid } from "@/components/product-grid"
import { LoadingSpinner } from "@/components/loading-spinner"
import { ErrorMessage } from "@/components/error-message"
import { ApiResponse, apiService } from "@/lib/api"
import type { Product } from "@/types/product"

interface ProductPageProps {
    params: {
        id: string
    }
}

async function ProductContent({ productId }: { productId: string }) {
    try {
        const product = await apiService.getProduct(productId)
        return <ProductDetail product={product.data} />
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
        const allProducts = await apiService.getProducts({
            pageSize: 5,
            categories: category ? [category] : []
        })
        // Filter out current product and get products from same category
        const relatedProducts = allProducts.data
            .filter((product) => product.productId !== currentProductId && (!category || product.category === category))

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
    let response: ApiResponse<Product> | null = null
    try {
        response = await apiService.getProduct(productId)
    } catch (error) {
        // Product will be null, handled in ProductContent
    }

    const product = response?.data

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

    if (productId) {
        return {
            title: "Produto não encontrado",
        }
    }

    try {
        const { data: product } = await apiService.getProduct(productId)
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
        const products = await apiService.getProducts()
        return products.data.map((product) => ({
            id: product.productId.toString(),
        }))
    } catch (error) {
        return []
    }
}
