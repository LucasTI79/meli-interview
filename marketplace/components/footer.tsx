import Link from "next/link"
import { Facebook, Twitter, Instagram, Mail, Phone, MapPin } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Separator } from "@/components/ui/separator"

export function Footer() {
  return (
    <footer className="bg-muted/30 border-t">
      <div className="container mx-auto px-4 py-12">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          <div className="space-y-4">
            <div className="flex items-center space-x-2">
              <div className="h-8 w-8 bg-primary rounded-lg flex items-center justify-center">
                <span className="text-primary-foreground font-bold text-sm">M</span>
              </div>
              <span className="font-bold text-xl">Marketplace</span>
            </div>
            <p className="text-muted-foreground text-sm">
              Sua loja online de confiança com produtos de alta qualidade e atendimento excepcional.
            </p>
            <div className="flex space-x-2">
              <Button variant="ghost" size="sm">
                <Facebook className="h-4 w-4" />
              </Button>
              <Button variant="ghost" size="sm">
                <Twitter className="h-4 w-4" />
              </Button>
              <Button variant="ghost" size="sm">
                <Instagram className="h-4 w-4" />
              </Button>
            </div>
          </div>

          <div className="space-y-4">
            <h3 className="font-semibold">Links Rápidos</h3>
            <nav className="flex flex-col space-y-2">
              <Link href="/products" className="text-sm text-muted-foreground hover:text-foreground">
                Produtos
              </Link>
              <Link href="/categories" className="text-sm text-muted-foreground hover:text-foreground">
                Categorias
              </Link>
              <Link href="/offers" className="text-sm text-muted-foreground hover:text-foreground">
                Ofertas
              </Link>
              <Link href="/about" className="text-sm text-muted-foreground hover:text-foreground">
                Sobre Nós
              </Link>
            </nav>
          </div>

          <div className="space-y-4">
            <h3 className="font-semibold">Atendimento</h3>
            <nav className="flex flex-col space-y-2">
              <Link href="/contact" className="text-sm text-muted-foreground hover:text-foreground">
                Contato
              </Link>
              <Link href="/shipping" className="text-sm text-muted-foreground hover:text-foreground">
                Entrega
              </Link>
              <Link href="/returns" className="text-sm text-muted-foreground hover:text-foreground">
                Trocas e Devoluções
              </Link>
              <Link href="/faq" className="text-sm text-muted-foreground hover:text-foreground">
                FAQ
              </Link>
            </nav>
          </div>

          <div className="space-y-4">
            <h3 className="font-semibold">Newsletter</h3>
            <p className="text-sm text-muted-foreground">Receba ofertas exclusivas e novidades em primeira mão.</p>
            <div className="flex space-x-2">
              <Input placeholder="Seu e-mail" className="flex-1" />
              <Button size="sm">Inscrever</Button>
            </div>
          </div>
        </div>

        <Separator className="my-8" />

        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-8">
          <div className="flex items-center space-x-2 text-sm text-muted-foreground">
            <Mail className="h-4 w-4" />
            <span>contato@marketplace.com</span>
          </div>
          <div className="flex items-center space-x-2 text-sm text-muted-foreground">
            <Phone className="h-4 w-4" />
            <span>(11) 9999-9999</span>
          </div>
          <div className="flex items-center space-x-2 text-sm text-muted-foreground">
            <MapPin className="h-4 w-4" />
            <span>São Paulo, SP</span>
          </div>
        </div>

        <Separator className="mb-8" />

        <div className="flex flex-col md:flex-row justify-between items-center space-y-4 md:space-y-0">
          <p className="text-sm text-muted-foreground">© 2024 Marketplace. Todos os direitos reservados.</p>
          <div className="flex space-x-4">
            <Link href="/privacy" className="text-sm text-muted-foreground hover:text-foreground">
              Privacidade
            </Link>
            <Link href="/terms" className="text-sm text-muted-foreground hover:text-foreground">
              Termos de Uso
            </Link>
          </div>
        </div>
      </div>
    </footer>
  )
}
