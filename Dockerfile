# Fase de build do frontend
FROM node:latest as build-stage
WORKDIR /app/frontend
COPY frontend/package*.json /app/frontend/
RUN npm install
COPY frontend/ /app/frontend/
RUN npm run build

# Fase de build do backend Go
FROM golang:alpine as go-build-stage
RUN apk add --no-cache gcc musl-dev
WORKDIR /app/backend
COPY backend/ /app/backend/
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -a -o ./bin/server .

# Fase final
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-stage /app/frontend/build /root/public
COPY --from=go-build-stage /app/backend/bin/server /root
COPY --from=go-build-stage /app/backend/data /root/data
RUN chmod +x /root/server
EXPOSE 8080
CMD ["/root/server"]