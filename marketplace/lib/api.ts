import { Category } from "@/types/product"

// API service layer for Go backend integration
export interface Product {
  productId: string
  name: string
  description: string
  price: number
  image: string
  category: string
  inStock: boolean
  rating?: number
  reviews?: number
}

export interface ApiResponse<T> {
  data: T
}

export interface PaginatedResponse<T> {
  data: T[]
  totalCount: number
  page: number;
  pageSize: number
}

export interface ProductFilters {
  categories?: string[]
  name?: string
  minPrice?: number
  maxPrice?: number
  inStock?: boolean
  sortBy?: "name" | "price" | "rating"
  sortOrder?: "asc" | "desc"
}

export interface ProductsParams extends ProductFilters {
  page?: number
  pageSize?: number
}

class ApiService {
  private baseUrl: string

  constructor() {
    // In production, this would come from environment variables
    this.baseUrl = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080"
  }

  private async request<T>(endpoint: string, options?: RequestInit): Promise<T> {
    const url = `${this.baseUrl}${endpoint}`

    try {
      const response = await fetch(url, {
        headers: {
          "Content-Type": "application/json",
          ...options?.headers,
        },
        ...options,
      })

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const data = await response.json()
      return data
    } catch (error) {
      console.error(`API request failed for ${endpoint}:`, error)
      throw error
    }
  }

  async getProducts(params?: ProductsParams): Promise<PaginatedResponse<Product>> {
    const searchParams = new URLSearchParams()

    if (params?.page) searchParams.set("page", params.page.toString())
    if (params?.pageSize) searchParams.set("pageSize", params.pageSize.toString())
    if (params?.categories?.length) searchParams.set("categories", String(params.categories))
    if (params?.name) searchParams.set("name", params.name)
    if (params?.minPrice) searchParams.set("minPrice", params.minPrice.toString())
    if (params?.maxPrice) searchParams.set("maxPrice", params.maxPrice.toString())
    if (params?.inStock !== undefined) searchParams.set("inStock", params.inStock.toString())
    if (params?.sortBy) searchParams.set("sortBy", params.sortBy)
    if (params?.sortOrder) searchParams.set("sortOrder", params.sortOrder)

    const queryString = searchParams.toString()
    const endpoint = `/api/v1/products${queryString ? `?${queryString}` : ""}`

    return this.request<PaginatedResponse<Product>>(endpoint)
  }

  async getProduct(id: string): Promise<ApiResponse<Product>> {
    return this.request<ApiResponse<Product>>(`/api/v1/products/${id}`)
  }

  async searchProducts(query: string): Promise<PaginatedResponse<Product>> {
    return this.request<PaginatedResponse<Product>>(`/api/v1/products?name=${encodeURIComponent(query)}`)
  }

  async getProductsByCategory(category: string): Promise<PaginatedResponse<Product>> {
    return this.request<PaginatedResponse<Product>>(`/api/v1/products?category=${encodeURIComponent(category)}`)
  }

  async getCategories(): Promise<ApiResponse<Category[]>> {
    return this.request<ApiResponse<Category[]>>(`/api/v1/categories`)  
  }
}

export const apiService = new ApiService()

// Mock data for development/demo purposes
export const mockProducts: Product[] = [
  {
    productId: "1",
    name: "Wireless Bluetooth Headphones",
    description: "Premium quality wireless headphones with noise cancellation and 30-hour battery life.",
    price: 199.99,
    image: "/wireless-bluetooth-headphones.jpg",
    category: "Electronics",
    inStock: true,
    rating: 4.5,
    reviews: 128,
  },
  {
    productId: "2",
    name: "Smart Fitness Watch",
    description: "Advanced fitness tracking with heart rate monitor, GPS, and waterproof design.",
    price: 299.99,
    image: "/smart-fitness-watch.png",
    category: "Electronics",
    inStock: true,
    rating: 4.3,
    reviews: 89,
  },
  {
    productId: "3",
    name: "Organic Cotton T-Shirt",
    description: "Comfortable and sustainable organic cotton t-shirt in various colors.",
    price: 29.99,
    image: "/organic-cotton-tshirt.png",
    category: "Clothing",
    inStock: true,
    rating: 4.7,
    reviews: 203,
  },
  {
    productId: "4",
    name: "Professional Camera Lens",
    description: "High-quality 50mm prime lens for professional photography.",
    price: 899.99,
    image: "/professional-camera-lens.jpg",
    category: "Photography",
    inStock: false,
    rating: 4.9,
    reviews: 45,
  },
  {
    productId: "5",
    name: "Ergonomic Office Chair",
    description: "Comfortable ergonomic office chair with lumbar support and adjustable height.",
    price: 449.99,
    image: "/ergonomic-office-chair.png",
    category: "Furniture",
    inStock: true,
    rating: 4.4,
    reviews: 167,
  },
  {
    productId: "6",
    name: "Stainless Steel Water Bottle",
    description: "Insulated stainless steel water bottle that keeps drinks cold for 24 hours.",
    price: 34.99,
    image: "/stainless-steel-bottle.png",
    category: "Lifestyle",
    inStock: true,
    rating: 4.6,
    reviews: 312,
  },
]

// Helper function to simulate API delay for development
export const delay = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms))

export const mockApiService = {
  async getProducts(params?: ProductsParams): Promise<PaginatedResponse<Product>> {
    await delay(500)

    let filteredProducts = [...mockProducts]

    if (params?.categories?.length) {
      filteredProducts = filteredProducts.filter(
        (product) => params?.categories?.map(c => c.toLowerCase()).includes(product.category.toLowerCase()),
      )
    }

    if (params?.name) {
      const searchTerm = params.name.toLowerCase()
      filteredProducts = filteredProducts.filter(
        (product) =>
          product.name.toLowerCase().includes(searchTerm)
      )
    }

    if (params?.minPrice) {
      filteredProducts = filteredProducts.filter((product) => product.price >= params.minPrice!)
    }

    if (params?.maxPrice) {
      filteredProducts = filteredProducts.filter((product) => product.price <= params.maxPrice!)
    }

    if (params?.inStock !== undefined) {
      filteredProducts = filteredProducts.filter((product) => product.inStock === params.inStock)
    }

    if (params?.sortBy) {
      filteredProducts.sort((a, b) => {
        let aValue: any = a[params.sortBy!]
        let bValue: any = b[params.sortBy!]

        if (params.sortBy === "name") {
          aValue = aValue.toLowerCase()
          bValue = bValue.toLowerCase()
        }

        if (params.sortOrder === "desc") {
          return bValue > aValue ? 1 : -1
        }
        return aValue > bValue ? 1 : -1
      })
    }

    const page = params?.page || 1
    const pageSize = params?.pageSize || 12
    const startIndex = (page - 1) * pageSize
    const endIndex = startIndex + pageSize
    const paginatedProducts = filteredProducts.slice(startIndex, endIndex)

    return {
      data: paginatedProducts,
      page,
      pageSize,
      totalCount: Math.ceil(filteredProducts.length / pageSize)
    }
  },

  async getProduct(id: string): Promise<Product> {
    await delay(300)
    const product = mockProducts.find((p) => p.productId === id)
    if (!product) {
      throw new Error(`Product with id ${id} not found`)
    }
    return product
  },

  async searchProducts(query: string): Promise<Product[]> {
    await delay(400)
    return mockProducts.filter(
      (product) =>
        product.name.toLowerCase().includes(query.toLowerCase())
    )
  },

  async getProductsByCategory(category: string): Promise<Product[]> {
    await delay(400)
    return mockProducts.filter((product) => product.category.toLowerCase() === category.toLowerCase())
  },
}
