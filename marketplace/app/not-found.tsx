"use client"

import Link from "next/link"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { Home, ArrowLeft, Search } from "lucide-react"

export default function NotFound() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100 dark:from-slate-900 dark:to-slate-800 flex items-center justify-center p-4">
      <Card className="w-full max-w-md mx-auto shadow-xl">
        <CardContent className="p-8 text-center space-y-6">
          {/* 404 Icon */}
          <div className="relative">
            <div className="w-24 h-24 mx-auto bg-gradient-to-br from-slate-100 to-slate-200 dark:from-slate-700 dark:to-slate-800 rounded-full flex items-center justify-center shadow-inner">
              <Search className="w-10 h-10 text-slate-400 dark:text-slate-500" />
            </div>
            <div className="absolute -top-2 -right-2 w-8 h-8 bg-red-100 dark:bg-red-900/30 rounded-full flex items-center justify-center">
              <span className="text-red-600 dark:text-red-400 text-sm font-bold">!</span>
            </div>
          </div>

          {/* Error Message */}
          <div className="space-y-2">
            <h1 className="text-3xl font-bold text-slate-900 dark:text-slate-100">404</h1>
            <h2 className="text-xl font-semibold text-slate-700 dark:text-slate-300">Página não encontrada</h2>
            <p className="text-slate-600 dark:text-slate-400 text-sm leading-relaxed">
              A página que você está procurando não existe ou ainda não foi implementada.
            </p>
          </div>

          {/* Action Buttons */}
          <div className="flex flex-col sm:flex-row gap-3 pt-4">
            <Button asChild variant="default" className="flex-1">
              <Link href="/" className="flex items-center justify-center gap-2">
                <Home className="w-4 h-4" />
                Início
              </Link>
            </Button>
            <Button asChild variant="outline" className="flex-1 bg-transparent">
              <Link href="/products" className="flex items-center justify-center gap-2">
                <Search className="w-4 h-4" />
                Produtos
              </Link>
            </Button>
          </div>

          {/* Back Button */}
          <Button
            variant="ghost"
            onClick={() => window.history.back()}
            className="w-full text-slate-600 dark:text-slate-400 hover:text-slate-900 dark:hover:text-slate-100"
          >
            <ArrowLeft className="w-4 h-4 mr-2" />
            Voltar
          </Button>
        </CardContent>
      </Card>
    </div>
  )
}
