import type React from "react"
import type { Metadata } from "next"
import { GeistSans } from "geist/font/sans"
import { GeistMono } from "geist/font/mono"
import { Analytics } from "@vercel/analytics/next"
import { Header } from "@/components/header"
import { Footer } from "@/components/footer"
import { CartProvider } from "@/contexts/cart-context"
import { QueryProvider } from "@/providers/query-provider"
import { Suspense } from "react"
import "./globals.css"

export const metadata: Metadata = {
  title: "Marketplace - Produtos de Qualidade",
  description: "Descubra produtos extraordinários em nosso marketplace de alta qualidade",
  generator: "v0.app",
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang="pt-BR">
      <body className={`font-sans ${GeistSans.variable} ${GeistMono.variable}`}>
        <QueryProvider>
          <CartProvider>
            <Suspense fallback={<div>Loading...</div>}>
              <div className="min-h-screen flex flex-col">
                <Header />
                <main className="flex-1">{children}</main>
                <Footer />
              </div>
            </Suspense>
          </CartProvider>
        </QueryProvider>
        <Analytics />
      </body>
    </html>
  )
}
