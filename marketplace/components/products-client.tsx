"use client"

import { useState, useMemo, useEffect } from "react"
import { useSearchParams } from "next/navigation"
import { Search, Filter, Grid, List, X, ChevronDown, Tag, DollarSign } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover"
import { Checkbox } from "@/components/ui/checkbox"
import { Label } from "@/components/ui/label"
import { Separator } from "@/components/ui/separator"
import { Badge } from "@/components/ui/badge"
import { ProductGrid } from "./product-grid"
import { ProductList } from "./product-list"
import type { Product } from "@/types/product"

interface ProductsClientProps {
  products: Product[]
}

type ViewMode = "grid" | "list"

interface Filters {
  categories: string[]
  priceRange: {
    min: number
    max: number
  }
}

export function ProductsClient({ products }: ProductsClientProps) {
  const searchParams = useSearchParams()
  const [searchTerm, setSearchTerm] = useState("")
  const [viewMode, setViewMode] = useState<ViewMode>("grid")
  const [filters, setFilters] = useState<Filters>({
    categories: [],
    priceRange: { min: 0, max: 1000 },
  })
  const [isFilterOpen, setIsFilterOpen] = useState(false)

  useEffect(() => {
    const categoryParam = searchParams.get("category")
    if (categoryParam) {
      setFilters((prev) => ({
        ...prev,
        categories: [categoryParam],
      }))
    }
  }, [searchParams])

  // Extract unique categories from products
  const categories = useMemo(() => {
    const uniqueCategories = [...new Set(products.map((p) => p.category))]
    return uniqueCategories.sort()
  }, [products])

  // Filter and search products
  const filteredProducts = useMemo(() => {
    return products.filter((product) => {
      // Search filter
      const matchesSearch =
        product.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        product.description.toLowerCase().includes(searchTerm.toLowerCase())

      // Category filter
      const matchesCategory = filters.categories.length === 0 || filters.categories.includes(product.category)

      // Price filter
      const matchesPrice = product.price >= filters.priceRange.min && product.price <= filters.priceRange.max

      return matchesSearch && matchesCategory && matchesPrice
    })
  }, [products, searchTerm, filters])

  const handleCategoryChange = (category: string, checked: boolean) => {
    setFilters((prev) => ({
      ...prev,
      categories: checked ? [...prev.categories, category] : prev.categories.filter((c) => c !== category),
    }))
  }

  const clearFilters = () => {
    setFilters({
      categories: [],
      priceRange: { min: 0, max: 1000 },
    })
    setSearchTerm("")
  }

  const removeCategoryFilter = (categoryToRemove: string) => {
    setFilters((prev) => ({
      ...prev,
      categories: prev.categories.filter((cat) => cat !== categoryToRemove),
    }))
  }

  const removeSearchFilter = () => {
    setSearchTerm("")
  }

  const activeFiltersCount = filters.categories.length + (searchTerm ? 1 : 0)

  return (
    <div>
      {/* Search and Controls */}
      <div className="flex flex-col sm:flex-row gap-4 mb-8">
        <div className="relative flex-1">
          <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-muted-foreground" />
          <Input
            placeholder="Buscar produtos..."
            className="pl-10 h-11 bg-background/50 border-border/50 focus:bg-background transition-colors"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>
        <div className="flex gap-2">
          <Popover open={isFilterOpen} onOpenChange={setIsFilterOpen}>
            <PopoverTrigger asChild>
              <Button
                variant="outline"
                size="sm"
                className="relative h-11 px-4 bg-gradient-to-r from-primary/5 to-primary/10 border-primary/20 hover:from-primary/10 hover:to-primary/20 transition-all duration-200"
              >
                <Filter className="h-4 w-4 mr-2" />
                Filtros
                <ChevronDown className="h-3 w-3 ml-2 opacity-50" />
                {activeFiltersCount > 0 && (
                  <Badge className="absolute -top-2 -right-2 h-5 w-5 p-0 flex items-center justify-center text-xs bg-primary">
                    {activeFiltersCount}
                  </Badge>
                )}
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-80 p-0 shadow-xl border-border/50" align="end">
              <div className="bg-gradient-to-b from-background to-background/95">
                {/* Header */}
                <div className="flex items-center justify-between p-6 pb-4">
                  <div className="flex items-center gap-2">
                    <Filter className="h-4 w-4 text-primary" />
                    <h4 className="font-semibold text-foreground">Filtros</h4>
                  </div>
                  <Button
                    variant="ghost"
                    size="sm"
                    onClick={clearFilters}
                    className="h-8 px-2 text-muted-foreground hover:text-foreground"
                  >
                    <X className="h-4 w-4 mr-1" />
                    Limpar
                  </Button>
                </div>

                <Separator className="mx-6" />

                {/* Categories Section */}
                <div className="p-6 pb-4">
                  <div className="flex items-center gap-2 mb-4">
                    <Tag className="h-4 w-4 text-primary" />
                    <h5 className="font-medium text-foreground">Categorias</h5>
                  </div>
                  <div className="space-y-3">
                    {categories.map((category) => (
                      <div key={category} className="flex items-center space-x-3 group">
                        <Checkbox
                          id={category}
                          checked={filters.categories.includes(category)}
                          onCheckedChange={(checked) => handleCategoryChange(category, checked as boolean)}
                          className="data-[state=checked]:bg-primary data-[state=checked]:border-primary"
                        />
                        <Label
                          htmlFor={category}
                          className="text-sm capitalize cursor-pointer group-hover:text-primary transition-colors flex-1"
                        >
                          {category}
                        </Label>
                        {filters.categories.includes(category) && (
                          <Badge variant="secondary" className="text-xs px-2 py-0.5">
                            ✓
                          </Badge>
                        )}
                      </div>
                    ))}
                  </div>
                </div>

                <Separator className="mx-6" />

                {/* Price Range Section */}
                <div className="p-6">
                  <div className="flex items-center gap-2 mb-4">
                    <DollarSign className="h-4 w-4 text-primary" />
                    <h5 className="font-medium text-foreground">Faixa de Preço</h5>
                  </div>
                  <div className="flex gap-3">
                    <div className="flex-1">
                      <Label className="text-xs text-muted-foreground mb-1 block">Mínimo</Label>
                      <Input
                        type="number"
                        placeholder="R$ 0"
                        value={filters.priceRange.min}
                        onChange={(e) =>
                          setFilters((prev) => ({
                            ...prev,
                            priceRange: { ...prev.priceRange, min: Number(e.target.value) || 0 },
                          }))
                        }
                        className="h-9 text-sm"
                      />
                    </div>
                    <div className="flex items-end pb-2">
                      <span className="text-muted-foreground text-sm">até</span>
                    </div>
                    <div className="flex-1">
                      <Label className="text-xs text-muted-foreground mb-1 block">Máximo</Label>
                      <Input
                        type="number"
                        placeholder="R$ 1000"
                        value={filters.priceRange.max}
                        onChange={(e) =>
                          setFilters((prev) => ({
                            ...prev,
                            priceRange: { ...prev.priceRange, max: Number(e.target.value) || 1000 },
                          }))
                        }
                        className="h-9 text-sm"
                      />
                    </div>
                  </div>
                  {/* Price range indicator */}
                  <div className="mt-3 p-2 bg-muted/50 rounded-md">
                    <p className="text-xs text-muted-foreground text-center">
                      R$ {filters.priceRange.min} - R$ {filters.priceRange.max}
                    </p>
                  </div>
                </div>
              </div>
            </PopoverContent>
          </Popover>

          <div className="flex border border-border/50 rounded-md overflow-hidden bg-background/50">
            <Button
              variant={viewMode === "grid" ? "default" : "ghost"}
              size="sm"
              onClick={() => setViewMode("grid")}
              className="h-11 px-4 rounded-none border-0"
            >
              <Grid className="h-4 w-4" />
            </Button>
            <Separator orientation="vertical" className="h-6 self-center" />
            <Button
              variant={viewMode === "list" ? "default" : "ghost"}
              size="sm"
              onClick={() => setViewMode("list")}
              className="h-11 px-4 rounded-none border-0"
            >
              <List className="h-4 w-4" />
            </Button>
          </div>
        </div>
      </div>

      {/* Results Info with Active Filters */}
      <div className="mb-6">
        <div className="flex flex-col sm:flex-row sm:items-center gap-3">
          <p className="text-sm text-muted-foreground">
            {filteredProducts.length} produto{filteredProducts.length !== 1 ? "s" : ""} encontrado
            {filteredProducts.length !== 1 ? "s" : ""}
            {searchTerm && ` para "${searchTerm}"`}
          </p>

          {(filters.categories.length > 0 || searchTerm) && (
            <div className="flex flex-wrap gap-2">
              {searchTerm && (
                <Badge variant="secondary" className="gap-1">
                  <Search className="h-3 w-3" />
                  {searchTerm}
                  <button onClick={removeSearchFilter} className="ml-1 hover:text-destructive transition-colors">
                    <X className="h-3 w-3" />
                  </button>
                </Badge>
              )}
              {filters.categories.map((category) => (
                <Badge key={category} variant="secondary" className="gap-1">
                  <Tag className="h-3 w-3" />
                  {category}
                  <button
                    onClick={() => removeCategoryFilter(category)}
                    className="ml-1 hover:text-destructive transition-colors"
                  >
                    <X className="h-3 w-3" />
                  </button>
                </Badge>
              ))}
            </div>
          )}
        </div>
      </div>

      {/* Products Display */}
      {viewMode === "grid" ? <ProductGrid products={filteredProducts} /> : <ProductList products={filteredProducts} />}
    </div>
  )
}
