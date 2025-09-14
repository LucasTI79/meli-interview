import Image from "next/image"
import Link from "next/link"
import { Star, ShoppingCart } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import type { Product } from "@/types/product"
import { formatPrice } from "@/utils/format-price"

interface ProductListItemProps {
    product: Product
}

export function ProductListItem({ product }: ProductListItemProps) {
    return (
        <div className="flex gap-4 p-4 border rounded-lg hover:shadow-md transition-shadow">
            <Link href={`/products/${product.productId}`} className="flex-shrink-0">
                <Image
                    src={product.image || "/placeholder.svg"}
                    alt={product.name}
                    width={96}
                    height={96}
                    className="rounded-lg object-cover w-24 h-24"
                />
            </Link>

            <div className="flex-1 min-w-0">
                <div className="flex flex-col sm:flex-row sm:items-start sm:justify-between gap-2">
                    <div className="flex-1">
                        <Link href={`/products/${product.productId}`}>
                            <h3 className="font-semibold text-lg hover:text-primary transition-colors line-clamp-1">
                                {product.name}
                            </h3>
                        </Link>

                        <Badge variant="secondary" className="mt-1 capitalize">
                            {product.category}
                        </Badge>

                        <p className="text-muted-foreground mt-2 line-clamp-2 text-sm">{product.description}</p>

                        <div className="flex items-center gap-2 mt-2">
                            <div className="flex items-center">
                                {Array.from({ length: 5 }).map((_, i) => (
                                    <Star
                                        key={i}
                                        className={`h-4 w-4 ${i < Math.floor(product?.rating ?? 0) ? "fill-yellow-400 text-yellow-400" : "text-muted-foreground"
                                            }`}
                                    />
                                ))}
                            </div>
                            <span className="text-sm text-muted-foreground">({product.rating})</span>
                        </div>
                    </div>

<div className="flex flex-col justify-between items-center gap-2 self-stretch">
                        <div className="text-right">
                            <p className="text-2xl font-bold text-primary">R$ {product.price.toFixed(2)}</p>
                            {product.originalPrice ? product.originalPrice > 0 && product.originalPrice > product.price && (
                                <p className="text-sm text-muted-foreground line-through">
                                    {formatPrice(product.originalPrice)}
                                </p>
                            ) : null}

                        </div>

                        <Button size="sm" className="w-full sm:w-auto">
                            <ShoppingCart className="h-4 w-4 mr-2" />
                            Adicionar
                        </Button>
                    </div>
                </div>
            </div>
        </div>
    )
}
