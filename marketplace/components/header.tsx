"use client"

import type React from "react"

import Link from "next/link"
import { useState, useEffect } from "react"
import { useRouter, useSearchParams } from "next/navigation"
import { Search, ShoppingCart, User, Menu, Home, Package, Info, Tag } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Badge } from "@/components/ui/badge"
import { useCart } from "@/contexts/cart-context"
import { useDebounce } from "@/hooks/use-debounce"
import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
} from "@/components/ui/navigation-menu"
import { Sheet, SheetContent, SheetTrigger, SheetHeader, SheetTitle } from "@/components/ui/sheet"

export function Header() {
  const [isSearchOpen, setIsSearchOpen] = useState(false)
  const [searchQuery, setSearchQuery] = useState("")
  const router = useRouter()
  const searchParams = useSearchParams()
  const { state } = useCart()

  const debouncedSearchQuery = useDebounce(searchQuery, 250)

  useEffect(() => {
    const searchParam = searchParams.get("name")
    if (searchParam) {
      setSearchQuery(searchParam)
    }
  }, [searchParams])

  useEffect(() => {
    if (debouncedSearchQuery.trim() && debouncedSearchQuery !== searchParams.get("name")) {
      router.push(`/products?name=${encodeURIComponent(debouncedSearchQuery.trim())}`)
    } else if (!debouncedSearchQuery.trim() && searchParams.get("name")) {
      router.push("/products")
    }
  }, [debouncedSearchQuery, router, searchParams])

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault()
    if (searchQuery.trim()) {
      router.push(`/products?name=${encodeURIComponent(searchQuery.trim())}`)
      setIsSearchOpen(false)
    }
  }

  const categories = [
    { name: "Eletrônicos", href: "/products?category=Electronics", icon: Package },
    { name: "Roupas", href: "/products?category=Clothing", icon: Tag },
    { name: "Móveis", href: "/products?category=Furniture", icon: Home },
    { name: "Fotografia", href: "/products?category=Photography", icon: Package },
    { name: "Estilo de Vida", href: "/products?category=Lifestyle", icon: Package },
  ]

  return (
    <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container mx-auto px-4">
        <div className="flex h-16 items-center justify-between">
          <Link href="/" className="flex items-center space-x-2">
            <div className="h-8 w-8 bg-slate-900 dark:bg-slate-100 rounded-lg flex items-center justify-center shadow-sm">
              <span className="text-white dark:text-slate-900 font-bold text-sm">M</span>
            </div>
            <span className="font-bold text-xl bg-gradient-to-r from-foreground to-foreground/80 bg-clip-text text-transparent">
              Marketplace
            </span>
          </Link>

          <NavigationMenu className="hidden lg:flex">
            <NavigationMenuList>
              <NavigationMenuItem>
                <Link href="/products" legacyBehavior passHref>
                  <NavigationMenuLink className="group inline-flex h-10 w-max items-center justify-center rounded-md bg-background px-4 py-2 text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground focus:outline-none disabled:pointer-events-none disabled:opacity-50 data-[active]:bg-accent/50 data-[state=open]:bg-accent/50">
                    Produtos
                  </NavigationMenuLink>
                </Link>
              </NavigationMenuItem>

              <NavigationMenuItem>
                <NavigationMenuTrigger>Categorias</NavigationMenuTrigger>
                <NavigationMenuContent>
                  <div className="grid w-[400px] gap-3 p-4 md:w-[500px] md:grid-cols-2 lg:w-[600px]">
                    {categories.map((category) => {
                      const IconComponent = category.icon
                      return (
                        <Link
                          key={category.name}
                          href={category.href}
                          className="flex items-center space-x-3 select-none rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground"
                        >
                          <IconComponent className="h-4 w-4 text-muted-foreground" />
                          <div className="text-sm font-medium leading-none">{category.name}</div>
                        </Link>
                      )
                    })}
                  </div>
                </NavigationMenuContent>
              </NavigationMenuItem>

              <NavigationMenuItem>
                <Link href="/about" legacyBehavior passHref>
                  <NavigationMenuLink className="group inline-flex h-10 w-max items-center justify-center rounded-md bg-background px-4 py-2 text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground focus:outline-none disabled:pointer-events-none disabled:opacity-50 data-[active]:bg-accent/50 data-[state=open]:bg-accent/50">
                    Sobre
                  </NavigationMenuLink>
                </Link>
              </NavigationMenuItem>
            </NavigationMenuList>
          </NavigationMenu>

          <div className="hidden md:flex flex-1 max-w-sm mx-8">
            <form onSubmit={handleSearch} className="relative w-full">
              <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-muted-foreground" />
              <Input
                placeholder="Buscar produtos..."
                className="pl-10"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
              />
            </form>
          </div>

          <div className="flex items-center space-x-2">
            <Button variant="ghost" size="sm" className="md:hidden" onClick={() => setIsSearchOpen(!isSearchOpen)}>
              <Search className="h-5 w-5" />
            </Button>

            <Link href="/cart">
              <Button variant="ghost" size="sm" className="relative">
                <ShoppingCart className="h-5 w-5" />
                {state.itemCount > 0 && (
                  <Badge variant="destructive" className="absolute -top-1 -right-1 h-5 w-5 rounded-full p-0 text-xs">
                    {state.itemCount}
                  </Badge>
                )}
              </Button>
            </Link>

            <Button variant="ghost" size="sm">
              <User className="h-5 w-5" />
            </Button>

            <Sheet>
              <SheetTrigger asChild>
                <Button variant="ghost" size="sm" className="lg:hidden">
                  <Menu className="h-5 w-5" />
                </Button>
              </SheetTrigger>
              <SheetContent side="right" className="w-[320px] sm:w-[400px] p-0">
                <div className="flex flex-col h-full">
                  <SheetHeader className="px-6 py-4 border-b bg-gradient-to-r from-background to-muted/20">
                    <SheetTitle className="flex items-center space-x-2 text-left">
                      <div className="h-6 w-6 bg-slate-900 dark:bg-slate-100 rounded-md flex items-center justify-center">
                        <span className="text-white dark:text-slate-900 font-bold text-xs">M</span>
                      </div>
                      <span className="font-semibold">Menu</span>
                    </SheetTitle>
                  </SheetHeader>

                  <nav className="flex-1 px-6 py-6 space-y-6">
                    <div className="space-y-3">
                      <Link
                        href="/products"
                        className="flex items-center space-x-3 text-lg font-medium p-3 rounded-lg hover:bg-accent transition-colors group"
                      >
                        <Package className="h-5 w-5 text-muted-foreground group-hover:text-accent-foreground transition-colors" />
                        <span>Produtos</span>
                      </Link>

                      <Link
                        href="/about"
                        className="flex items-center space-x-3 text-lg font-medium p-3 rounded-lg hover:bg-accent transition-colors group"
                      >
                        <Info className="h-5 w-5 text-muted-foreground group-hover:text-accent-foreground transition-colors" />
                        <span>Sobre</span>
                      </Link>
                    </div>

                    <div className="space-y-3">
                      <div className="flex items-center space-x-2 px-3">
                        <Tag className="h-4 w-4 text-muted-foreground" />
                        <p className="text-sm font-semibold text-muted-foreground uppercase tracking-wider">
                          Categorias
                        </p>
                      </div>
                      <div className="space-y-1 pl-3">
                        {categories.map((category) => {
                          const IconComponent = category.icon
                          return (
                            <Link
                              key={category.name}
                              href={category.href}
                              className="flex items-center space-x-3 p-3 rounded-lg text-sm hover:bg-accent transition-colors group"
                            >
                              <IconComponent className="h-4 w-4 text-muted-foreground group-hover:text-accent-foreground transition-colors" />
                              <span className="text-muted-foreground group-hover:text-foreground transition-colors">
                                {category.name}
                              </span>
                            </Link>
                          )
                        })}
                      </div>
                    </div>
                  </nav>

                  <div className="px-6 py-4 border-t bg-muted/20">
                    <div className="flex items-center justify-between">
                      <div className="flex items-center space-x-3">
                        <div className="h-8 w-8 bg-muted rounded-full flex items-center justify-center">
                          <User className="h-4 w-4 text-muted-foreground" />
                        </div>
                        <div className="text-sm">
                          <p className="font-medium">Minha Conta</p>
                          <p className="text-muted-foreground text-xs">Faça login ou cadastre-se</p>
                        </div>
                      </div>
                      <Link href="/cart">
                        <Button variant="ghost" size="sm" className="relative">
                          <ShoppingCart className="h-4 w-4" />
                          {state.itemCount > 0 && (
                            <Badge
                              variant="destructive"
                              className="absolute -top-1 -right-1 h-4 w-4 rounded-full p-0 text-xs"
                            >
                              {state.itemCount}
                            </Badge>
                          )}
                        </Button>
                      </Link>
                    </div>
                  </div>
                </div>
              </SheetContent>
            </Sheet>
          </div>
        </div>

        {isSearchOpen && (
          <div className="md:hidden py-4 border-t">
            <form onSubmit={handleSearch} className="relative">
              <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-muted-foreground" />
              <Input
                placeholder="Buscar produtos..."
                className="pl-10"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                autoFocus
              />
            </form>
          </div>
        )}
      </div>
    </header>
  )
}
