// API service layer for Go backend integration
export interface Product {
  id: string
  name: string
  description: string
  price: number
  originalPrice?: number;
  image: string
  category: string
  inStock: boolean
  rating?: number
  reviews?: number
}

export interface ApiResponse<T> {
  data: T
  success: boolean
  message?: string
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

  // Get all products
  async getProducts(): Promise<Product[]> {
    return this.request<Product[]>("/api/products")
  }

  // Get single product by ID
  async getProduct(id: "number"): Promise<Product> {
    return this.request<Product>(`/api/products/${id}`)
  }

  // Search products (future enhancement)
  async searchProducts(query: string): Promise<Product[]> {
    return this.request<Product[]>(`/api/products/search?q=${encodeURIComponent(query)}`)
  }

  // Get products by category (future enhancement)
  async getProductsByCategory(category: string): Promise<Product[]> {
    return this.request<Product[]>(`/api/products/category/${encodeURIComponent(category)}`)
  }
}

export const apiService = new ApiService()

// Mock data for development/demo purposes
export const mockProducts: Product[] = [
  {
    id: "1",
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
    id: "2",
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
    id: "3",
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
    id: "4",
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
    id: "5",
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
    id: "6",
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

// Mock API functions for development
export const mockApiService = {
  async getProducts(): Promise<Product[]> {
    await delay(500) // Simulate network delay
    return mockProducts
  },

  async getProduct(id: string): Promise<Product> {
    await delay(300)
    const product = mockProducts.find((p) => p.id === id)
    if (!product) {
      throw new Error(`Product with id ${id} not found`)
    }
    return product
  },

  async searchProducts(query: string): Promise<Product[]> {
    await delay(400)
    return mockProducts.filter(
      (product) =>
        product.name.toLowerCase().includes(query.toLowerCase()) ||
        product.description.toLowerCase().includes(query.toLowerCase()) ||
        product.category.toLowerCase().includes(query.toLowerCase()),
    )
  },

  async getProductsByCategory(category: string): Promise<Product[]> {
    await delay(400)
    return mockProducts.filter((product) => product.category.toLowerCase() === category.toLowerCase())
  },
}
