// Shared TypeScript interfaces for the marketplace
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

export interface CartItem extends Product {
  quantity: number
}

export interface Category {
  id: string
  name: string
  slug: string
  description?: string
}

export interface ApiError {
  message: string
  status: number
  code?: string
}
