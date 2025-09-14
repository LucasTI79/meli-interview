import Link from "next/link"
import { ArrowRight, ShoppingBag, Star, Truck } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { ProductGrid } from "@/components/product-grid"
import { apiService } from "@/lib/api"

// Server component that fetches featured products
async function FeaturedProducts() {
  try {
    const products = await apiService.getProducts({
        pageSize: 4
    })
    const featuredProducts = products.data
    return <ProductGrid products={featuredProducts} />
  } catch (error) {
    return (
      <div className="text-center py-8">
        <p className="text-muted-foreground">Erro ao carregar produtos em destaque</p>
      </div>
    )
  }
}

export default function HomePage() {
  return (
    <div className="min-h-screen">
      {/* Hero Section */}
      <section className="bg-gradient-to-br from-primary/5 via-background to-accent/5 py-20">
        <div className="container mx-auto px-4 text-center">
          <Badge variant="secondary" className="mb-4">
            Novo Marketplace
          </Badge>
          <h1 className="text-5xl md:text-6xl font-bold text-balance mb-6">
            Descubra Produtos
            <span className="text-primary block">Extraordinários</span>
          </h1>
          <p className="text-xl text-muted-foreground mb-8 max-w-2xl mx-auto text-balance">
            Uma curadoria especial de produtos de alta qualidade para transformar seu dia a dia
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Button size="lg" asChild>
              <Link href="/products">
                <ShoppingBag className="h-5 w-5 mr-2" />
                Explorar Produtos
              </Link>
            </Button>
            <Button size="lg" variant="outline" asChild>
              <Link href="/products">
                Ver Categorias
                <ArrowRight className="h-5 w-5 ml-2" />
              </Link>
            </Button>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-16 bg-muted/30">
        <div className="container mx-auto px-4">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <Card>
              <CardContent className="flex flex-col items-center text-center p-8">
                <div className="bg-primary/10 p-4 rounded-full mb-4">
                  <Truck className="h-8 w-8 text-primary" />
                </div>
                <h3 className="text-xl font-semibold mb-2">Entrega Rápida</h3>
                <p className="text-muted-foreground">
                  Frete grátis para compras acima de R$ 99 e entrega expressa disponível
                </p>
              </CardContent>
            </Card>

            <Card>
              <CardContent className="flex flex-col items-center text-center p-8">
                <div className="bg-primary/10 p-4 rounded-full mb-4">
                  <Star className="h-8 w-8 text-primary" />
                </div>
                <h3 className="text-xl font-semibold mb-2">Qualidade Garantida</h3>
                <p className="text-muted-foreground">
                  Produtos cuidadosamente selecionados com garantia de qualidade e satisfação
                </p>
              </CardContent>
            </Card>

            <Card>
              <CardContent className="flex flex-col items-center text-center p-8">
                <div className="bg-primary/10 p-4 rounded-full mb-4">
                  <ShoppingBag className="h-8 w-8 text-primary" />
                </div>
                <h3 className="text-xl font-semibold mb-2">Experiência Premium</h3>
                <p className="text-muted-foreground">
                  Interface intuitiva e suporte dedicado para uma experiência de compra única
                </p>
              </CardContent>
            </Card>
          </div>
        </div>
      </section>

      {/* Featured Products Section */}
      <section className="py-16">
        <div className="container mx-auto px-4">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold mb-4">Produtos em Destaque</h2>
            <p className="text-muted-foreground text-lg max-w-2xl mx-auto">
              Selecionamos os melhores produtos para você descobrir
            </p>
          </div>

          <FeaturedProducts />

          <div className="text-center mt-12">
            <Button size="lg" variant="outline" asChild>
              <Link href="/products">
                Ver Todos os Produtos
                <ArrowRight className="h-5 w-5 ml-2" />
              </Link>
            </Button>
          </div>
        </div>
      </section>
    </div>
  )
}
