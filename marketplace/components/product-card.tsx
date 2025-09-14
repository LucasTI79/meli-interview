"use client"

import Image from "next/image"
import Link from "next/link"
import { Star, ShoppingCart } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardFooter } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { useCart } from "@/contexts/cart-context"
import type { Product } from "@/types/product"
import { formatPrice } from "@/utils/format-price"

interface ProductCardProps {
  product: Product
}

export function ProductCard({ product }: ProductCardProps) {
  const { addItem } = useCart()

  const renderStars = (rating: number) => {
    const stars = []
    const fullStars = Math.floor(rating)
    const hasHalfStar = rating % 1 !== 0

    for (let i = 0; i < fullStars; i++) {
      stars.push(<Star key={i} className="h-4 w-4 fill-primary text-primary" />)
    }

    if (hasHalfStar) {
      stars.push(<Star key="half" className="h-4 w-4 fill-primary/50 text-primary" />)
    }

    const emptyStars = 5 - Math.ceil(rating)
    for (let i = 0; i < emptyStars; i++) {
      stars.push(<Star key={`empty-${i}`} className="h-4 w-4 text-muted-foreground" />)
    }

    return stars
  }

  return (
    <Card className="group overflow-hidden transition-all duration-300 hover:shadow-lg hover:-translate-y-1 flex flex-col justify-between">
      <div className="relative aspect-square overflow-hidden">
        <Link href={`/products/${product.productId}`}>
          <Image
            src={product.image || "/placeholder.svg"}
            alt={product.name}
            fill
            className="object-cover transition-transform duration-300 group-hover:scale-105"
            sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
          />
        </Link>
        {!product.inStock && (
          <Badge variant="destructive" className="absolute top-2 right-2">
            Esgotado
          </Badge>
        )}
        <Badge variant="secondary" className="absolute top-2 left-2">
          {product.category}
        </Badge>
      </div>

      <CardContent className="p-4">
        <Link href={`/products/${product.productId}`}>
          <h3 className="font-semibold text-lg mb-2 line-clamp-2 hover:text-primary transition-colors">
            {product.name}
          </h3>
        </Link>

        <p className="text-muted-foreground text-sm mb-3 line-clamp-2">{product.description}</p>

        {product.rating && (
          <div className="flex items-center gap-2 mb-3">
            <div className="flex items-center">{renderStars(product.rating)}</div>
            <span className="text-sm text-muted-foreground">({product.reviews})</span>
          </div>
        )}

        <div className="flex items-center justify-between">
          <span className="text-2xl font-bold text-primary">{formatPrice(product.price)}</span>
        </div>
      </CardContent>

      <CardFooter className="p-4 pt-0">
        <Button
          className="w-full pointer"
          disabled={!product.inStock}
          variant={product.inStock ? "default" : "secondary"}
          onClick={() => product.inStock && addItem(product)}
        >
          <ShoppingCart className="h-4 w-4 mr-2" />
          {product.inStock ? "Adicionar ao Carrinho" : "Indisponível"}
        </Button>
      </CardFooter>
    </Card>
  )
}
