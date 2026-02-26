# Monorepo setup from
# https://volovoy.com/article/javascript-monorepo-with-docker-guide/
FROM node:24-alpine AS setup

WORKDIR /app

# Copy all packages
COPY apps apps
# COPY packages packages

# Find and remove non-package.json files
RUN find . \
    -mindepth 3 -maxdepth 3 \
    \! -name "package.json" \
    -exec rm -rf {} +

# Copy other files required for a build
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./

# Build the application
FROM node:24-alpine AS builder

WORKDIR /app

ARG PUBLIC_SPECTRAL_GRPC_CLIENT_ORIGIN
RUN if [ -z "$PUBLIC_SPECTRAL_GRPC_CLIENT_ORIGIN" ]; then \
    echo "Error: PUBLIC_SPECTRAL_GRPC_CLIENT_ORIGIN build argument is required." \
    exit 1; \
    fi

ARG NODE_ENV=production
ENV NODE_ENV=production

COPY --from=setup /app .

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN corepack enable pnpm

# Install packages with cache mount for pnpm store
RUN --mount=type=cache,target=/pnpm/store \
    pnpm install --frozen-lockfile --prefer-offline

COPY . .

RUN pnpm build:frontend

# Production stage
FROM nginx:alpine AS production

COPY ./infra/docker/frontend/nginx.conf /etc/nginx/nginx.conf

COPY --from=builder /app/apps/frontend/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
