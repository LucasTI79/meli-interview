"use client"

import { useState, useMemo, useEffect } from "react"
import { useSearchParams, useRouter, usePathname } from "next/navigation"
import { useQuery } from "@tanstack/react-query"
import { Search, Filter, Grid, List, X, ChevronDown, Tag, DollarSign, ChevronLeft, ChevronRight, Package, RefreshCw, Home } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover"
import { Checkbox } from "@/components/ui/checkbox"
import { Label } from "@/components/ui/label"
import { Separator } from "@/components/ui/separator"
import { Badge } from "@/components/ui/badge"
import { ProductGrid } from "./product-grid"
import { ProductList } from "./product-list"
import { apiService, mockApiService, type ProductsParams } from "@/lib/api"
import { formatPrice } from "@/utils/format-price"
import Link from "next/link"
import { useDebounce } from "@/hooks/use-debounce"

type ViewMode = "grid" | "list"

interface Filters {
    categories: string[]
    priceRange: {
        min: number
        max: number
    }
}


export function ProductsClient() {
    const searchParams = useSearchParams()
    const router = useRouter()
    const pathname = usePathname()

    const [viewMode, setViewMode] = useState<ViewMode>("grid")
    const [isFilterOpen, setIsFilterOpen] = useState(false)

    const [searchTerm, setSearchTerm] = useState(
        searchParams.get("name") || "",
    )
    const searchDebounce = useDebounce(searchTerm, 500)
    const [filters, setFilters] = useState<Filters>({
        categories: [],
        priceRange: { min: 0, max: 1000 },
    })

    const [currentPage, setCurrentPage] = useState(
        Number(searchParams.get("page")) || 1,
    )
    const itemsPerPage = 12

    const queryParams: ProductsParams = useMemo(
        () => ({
            page: currentPage,
            limit: itemsPerPage,
            name: searchDebounce || undefined,
            categories: filters.categories || undefined,
            minPrice: filters.priceRange.min > 0 ? filters.priceRange.min : undefined,
            maxPrice: filters.priceRange.max < 1000 ? filters.priceRange.max : undefined,
            sortBy: "name",
            sortOrder: "asc",
        }),
        [currentPage, searchDebounce, filters],
    )

    const { data, isLoading, error, refetch } = useQuery({
        queryKey: ["products", queryParams],
        queryFn: () => apiService.getProducts(queryParams),
        staleTime: 5 * 60 * 1000, // 5 minutes
    })

    const products = data?.data || []
    const pagination = {
        page: data?.page ?? 1,
        pageSize: data?.pageSize,
        totalCount: data?.totalCount,
        totalPages: data?.totalCount ? Math.ceil(data?.totalCount / data?.pageSize) : 0
    }

    const hasActiveFilters = filters.categories?.length || filters.priceRange.max || filters.priceRange.min

    const updateURL = (newParams: Record<string, string | undefined>) => {
        const params = new URLSearchParams(searchParams.toString())

        Object.entries(newParams).forEach(([key, value]) => {
            if (value) {
                params.set(key, value)
            } else {
                params.delete(key)
            }
        })

        router.push(`${pathname}?${params.toString()}`, { scroll: false })
    }

    // Extract unique categories from all products (for filter options)
    const { data: categories } = useQuery({
        queryKey: ["categories"],
        queryFn: () => apiService.getCategories(),
        staleTime: 10 * 60 * 1000, // 10 minutes
    })

    const handleCategoryChange = (category: string, checked: boolean) => {
        const newCategories = checked ? [...filters.categories, category] : filters.categories.filter((c) => c !== category)

        setFilters((prev) => ({
            ...prev,
            categories: newCategories,
        }))

        setCurrentPage(1)
        updateURL({
            categories: newCategories.join(",") || undefined,
            page: undefined,
        })
    }

    useEffect(() => {
        setCurrentPage(1)
        updateURL({
            name: searchDebounce || undefined,
            page: undefined,
        })
    }, [searchDebounce])

    const handlePageChange = (page: number) => {
        setCurrentPage(page)
        updateURL({ page: page.toString() })
        window.scrollTo({ top: 0, behavior: "smooth" })
    }

    const clearFilters = () => {
        setFilters({
            categories: [],
            priceRange: { min: 0, max: 1000 },
        })
        setSearchTerm("")
        setCurrentPage(1)
        router.push(pathname)
    }

    const EmptyState = () => (
        <div className="flex flex-col items-center justify-center py-16 px-4 text-center">
            <div className="relative mb-6">
                <div className="w-24 h-24 bg-muted rounded-full flex items-center justify-center">
                    <Package className="w-12 h-12 text-muted-foreground" />
                </div>
                {hasActiveFilters && (
                    <div className="absolute -top-2 -right-2 w-8 h-8 bg-orange-100 dark:bg-orange-900/20 rounded-full flex items-center justify-center">
                        <Search className="w-4 h-4 text-orange-600 dark:text-orange-400" />
                    </div>
                )}
            </div>

            <h3 className="text-2xl font-semibold mb-2">
                {hasActiveFilters ? "Nenhum produto encontrado" : "Nenhum produto disponível"}
            </h3>

            <p className="text-muted-foreground mb-6 max-w-md">
                {hasActiveFilters
                    ? "Não encontramos produtos que correspondam aos seus filtros. Tente ajustar os critérios de busca."
                    : "Não há produtos disponíveis no momento. Volte em breve para ver novidades!"}
            </p>

            {hasActiveFilters && (
                <div className="flex flex-col sm:flex-row gap-3 mb-6">
                    <Button onClick={clearFilters} variant="outline" className="flex items-center gap-2 bg-transparent">
                        <X className="w-4 h-4" />
                        Limpar filtros
                    </Button>
                    <Button onClick={() => refetch()} variant="outline" className="flex items-center gap-2">
                        <RefreshCw className="w-4 h-4" />
                        Tentar novamente
                    </Button>
                </div>
            )}

            <div className="flex flex-col sm:flex-row gap-3">
                <Link href="/">
                    <Button variant="default" className="flex items-center gap-2">
                        <Home className="w-4 h-4" />
                        Voltar ao início
                    </Button>
                </Link>
                {hasActiveFilters && (
                    <Link href="/products">
                        <Button variant="outline">Ver todos os produtos</Button>
                    </Link>
                )}
            </div>

            {hasActiveFilters && (
                <div className="mt-8 p-4 bg-muted/50 rounded-lg max-w-md">
                    <p className="text-sm text-muted-foreground mb-2 font-medium">Dicas de busca:</p>
                    <ul className="text-sm text-muted-foreground space-y-1">
                        <li>• Tente termos mais gerais</li>
                        <li>• Verifique a ortografia</li>
                        <li>• Use filtros de categoria diferentes</li>
                    </ul>
                </div>
            )}
        </div>
    )

    const removeCategoryFilter = (categoryToRemove: string) => {
        const newCategories = filters.categories.filter((cat) => cat !== categoryToRemove)
        setFilters((prev) => ({
            ...prev,
            categories: newCategories,
        }))
        setCurrentPage(1)
        updateURL({
            category: newCategories[0] || undefined,
            page: undefined,
        })
    }

    const removeSearchFilter = () => {
        setSearchTerm("")
        setCurrentPage(1)
        updateURL({
            name: undefined,
            page: undefined,
        })
    }

    const activeFiltersCount = filters.categories.length + (searchTerm ? 1 : 0)

    if (error) {
        return (
            <div className="text-center py-12">
                <p className="text-muted-foreground">Erro ao carregar produtos. Tente novamente.</p>
                <Button onClick={() => window.location.reload()} className="mt-4">
                    Recarregar
                </Button>
            </div>
        )
    }

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
                                        {categories?.data?.map(({ name: category }) => (
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
                                    <div className="mt-3 p-2 bg-muted/50 rounded-md">
                                        <p className="text-xs text-muted-foreground text-center">
                                            {formatPrice(filters.priceRange.min)} - {formatPrice(filters.priceRange.max)}
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
                        {isLoading ? (
                            "Carregando..."
                        ) : (
                            <>
                                {pagination?.totalCount || 0} produto{(pagination?.totalCount || 0) !== 1 ? "s" : ""} encontrado
                                {(pagination?.totalCount || 0) !== 1 ? "s" : ""}
                                {searchTerm && ` para "${searchTerm}"`}
                                {pagination && pagination.totalPages > 1 && (
                                    <>
                                        {" "}
                                        • Página {pagination.page} de {pagination.totalPages}
                                    </>
                                )}
                            </>
                        )}
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

            {/* Loading State */}
            {isLoading && (
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                    {Array.from({ length: 8 }).map((_, i) => (
                        <div key={i} className="animate-pulse">
                            <div className="bg-muted rounded-lg h-64 mb-4"></div>
                            <div className="bg-muted rounded h-4 mb-2"></div>
                            <div className="bg-muted rounded h-4 w-2/3"></div>
                        </div>
                    ))}
                </div>
            )}

            {/* Products Display */}
            {!isLoading && products.length > 0 ? (
                <>
                    {viewMode === "grid" ? <ProductGrid products={products} /> : <ProductList products={products} />}

                    {pagination && pagination.totalPages > 1 && (
                        <div className="flex items-center justify-center gap-2 mt-12">
                            <Button
                                variant="outline"
                                size="sm"
                                onClick={() => handlePageChange(currentPage - 1)}
                                disabled={currentPage === 1}
                                className="gap-2"
                            >
                                <ChevronLeft className="h-4 w-4" />
                                Anterior
                            </Button>

                            <div className="flex gap-1">
                                {Array.from({ length: Math.min(5, pagination.totalPages) }, (_, i) => {
                                    let pageNum
                                    if (pagination.totalPages <= 5) {
                                        pageNum = i + 1
                                    } else if (currentPage <= 3) {
                                        pageNum = i + 1
                                    } else if (currentPage >= pagination.totalPages - 2) {
                                        pageNum = pagination.totalPages - 4 + i
                                    } else {
                                        pageNum = currentPage - 2 + i
                                    }

                                    return (
                                        <Button
                                            key={pageNum}
                                            variant={currentPage === pageNum ? "default" : "outline"}
                                            size="sm"
                                            onClick={() => handlePageChange(pageNum)}
                                            className="w-10 h-10 p-0"
                                        >
                                            {pageNum}
                                        </Button>
                                    )
                                })}
                            </div>

                            <Button
                                variant="outline"
                                size="sm"
                                onClick={() => handlePageChange(currentPage + 1)}
                                disabled={currentPage === pagination.totalPages}
                                className="gap-2"
                            >
                                Próxima
                                <ChevronRight className="h-4 w-4" />
                            </Button>
                        </div>
                    )}
                </>
            ) : (
                <EmptyState />
            )}
        </div>
    )
}
