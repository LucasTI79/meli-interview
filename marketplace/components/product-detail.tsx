"use client"

import Image from "next/image"
import { Star, ShoppingCart, Heart, Share2, Truck, Shield, RotateCcw } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Separator } from "@/components/ui/separator"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { useCart } from "@/contexts/cart-context"
import type { Product } from "@/types/product"
import { formatPrice } from "@/utils/format-price"

interface ProductDetailProps {
  product: Product
}

export function ProductDetail({ product }: ProductDetailProps) {
  const { addItem } = useCart()

  const handleAddToCart = () => {
    addItem(product)
  }

  const renderStars = (rating: number) => {
    const stars = []
    const fullStars = Math.floor(rating)
    const hasHalfStar = rating % 1 !== 0

    for (let i = 0; i < fullStars; i++) {
      stars.push(<Star key={i} className="h-5 w-5 fill-primary text-primary" />)
    }

    if (hasHalfStar) {
      stars.push(<Star key="half" className="h-5 w-5 fill-primary/50 text-primary" />)
    }

    const emptyStars = 5 - Math.ceil(rating)
    for (let i = 0; i < emptyStars; i++) {
      stars.push(<Star key={`empty-${i}`} className="h-5 w-5 text-muted-foreground" />)
    }

    return stars
  }

  return (
    <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <div className="space-y-4">
        <div className="relative aspect-square overflow-hidden rounded-lg border">
          <Image
            src={product.image || "/placeholder.svg"}
            alt={product.name}
            fill
            className="object-cover"
            sizes="(max-width: 768px) 100vw, 50vw"
            priority
          />
          {!product.inStock && (
            <Badge variant="destructive" className="absolute top-4 right-4">
              Esgotado
            </Badge>
          )}
        </div>
      </div>

      <div className="space-y-6">
        <div>
          <Badge variant="secondary" className="mb-2">
            {product.category}
          </Badge>
          <h1 className="text-3xl font-bold text-balance mb-4">{product.name}</h1>

          {product.rating ? (
            <div className="flex items-center gap-3 mb-4">
              <div className="flex items-center">{renderStars(product.rating)}</div>
              <span className="text-lg font-medium">{product.rating}</span>
              <span className="text-muted-foreground">({product.reviews} avaliações)</span>
            </div>
          ) : null}

          <div className="flex items-baseline gap-2 mb-6">
            <span className="text-4xl font-bold text-primary">{formatPrice(product.price)}</span>
          </div>
        </div>

        <Separator />

        <div className="space-y-4">
          <div className="flex gap-3">
            <Button size="lg" className="flex-1" disabled={!product.inStock} onClick={handleAddToCart}>
              <ShoppingCart className="h-5 w-5 mr-2" />
              {product.inStock ? "Adicionar ao Carrinho" : "Indisponível"}
            </Button>
            <Button size="lg" variant="outline">
              <Heart className="h-5 w-5" />
            </Button>
            <Button size="lg" variant="outline">
              <Share2 className="h-5 w-5" />
            </Button>
          </div>

          {product.inStock && (
            <Button size="lg" variant="secondary" className="w-full">
              Comprar Agora
            </Button>
          )}
        </div>

        <Separator />

        <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
          <Card>
            <CardContent className="flex items-center gap-3 p-4">
              <Truck className="h-5 w-5 text-primary" />
              <div>
                <p className="font-medium text-sm">Frete Grátis</p>
                <p className="text-xs text-muted-foreground">Acima de R$ 99</p>
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardContent className="flex items-center gap-3 p-4">
              <Shield className="h-5 w-5 text-primary" />
              <div>
                <p className="font-medium text-sm">Garantia</p>
                <p className="text-xs text-muted-foreground">12 meses</p>
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardContent className="flex items-center gap-3 p-4">
              <RotateCcw className="h-5 w-5 text-primary" />
              <div>
                <p className="font-medium text-sm">Troca</p>
                <p className="text-xs text-muted-foreground">30 dias</p>
              </div>
            </CardContent>
          </Card>
        </div>

        <Separator />

        <Tabs defaultValue="description" className="w-full">
          <TabsList className="grid w-full grid-cols-3">
            <TabsTrigger value="description">Descrição</TabsTrigger>
            <TabsTrigger value="specifications">Especificações</TabsTrigger>
            <TabsTrigger value="reviews">Avaliações</TabsTrigger>
          </TabsList>

          <TabsContent value="description" className="mt-6">
            <div className="prose prose-sm max-w-none">
              <p className="text-muted-foreground leading-relaxed">{product.description}</p>
            </div>
          </TabsContent>

          <TabsContent value="specifications" className="mt-6">
            <div className="space-y-3">
              <div className="flex justify-between py-2 border-b">
                <span className="font-medium">Categoria</span>
                <span className="text-muted-foreground">{product.category}</span>
              </div>
              <div className="flex justify-between py-2 border-b">
                <span className="font-medium">Disponibilidade</span>
                <span className="text-muted-foreground">{product.inStock ? "Em estoque" : "Esgotado"}</span>
              </div>
              <div className="flex justify-between py-2 border-b">
                <span className="font-medium">SKU</span>
                <span className="text-muted-foreground">PRD-{product.productId.toString().padStart(6, "0")}</span>
              </div>
            </div>
          </TabsContent>

          <TabsContent value="reviews" className="mt-6">
            <div className="text-center py-8">
              <p className="text-muted-foreground">
                {product.reviews ? `${product.reviews} avaliações disponíveis` : "Nenhuma avaliação ainda"}
              </p>
            </div>
          </TabsContent>
        </Tabs>
      </div>
    </div>
  )
}
