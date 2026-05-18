# Stage 1: Build Vue frontend
FROM node:22-alpine AS frontend
WORKDIR /app
COPY frontend/package*.json ./
RUN npm ci
COPY frontend .
RUN npm run build

# Stage 2: Build Go backend
FROM golang:alpine AS builder
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend .
COPY --from=frontend /app/dist ./dist
RUN go build -o server .

# Stage 3: Minimal final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server ./server
COPY --from=builder /app/dist ./dist
EXPOSE 8080
CMD ["./server"]
