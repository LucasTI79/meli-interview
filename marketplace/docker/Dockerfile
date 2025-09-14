# Stage 1: build
FROM node:20-alpine AS builder

# Instalar pnpm
RUN corepack enable && corepack prepare pnpm@latest --activate

WORKDIR /app

# Copiar package.json + pnpm-lock.yaml
COPY package.json pnpm-lock.yaml ./

# Instalar dependências
RUN pnpm install --frozen-lockfile

# Copiar código
COPY . .

CMD [ "pnpm", "dev" ]

# # Build da aplicação
# RUN pnpm build

# # Stage 2: produção
# FROM node:20-alpine AS runner

# WORKDIR /app

# ENV NODE_ENV=production

# # Copiar node_modules + build do builder
# COPY --from=builder /app/node_modules ./node_modules
# COPY --from=builder /app/.next ./.next
# COPY --from=builder /app/public ./public
# COPY --from=builder /app/package.json ./package.json

# EXPOSE 3000

# CMD ["pnpm", "start"]
