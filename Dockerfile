# Fase de build do frontend
FROM node:latest as build-stage
WORKDIR /app
COPY frontend/package*.json /app/
RUN npm install
COPY frontend/ /app/
RUN npm run build

# Fase de build do backend Go
FROM golang:latest as go-build-stage

ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /go/src/app
COPY backend/ /go/src/app/
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -installsuffix cgo -o /go/bin/server

# Fase final
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-stage /app/build /root/public
COPY --from=go-build-stage /go/bin/server /root
COPY --from=go-build-stage /go/src/app/data /root/data
RUN ls -la /root
RUN chmod +x /root/server
EXPOSE 8080
CMD ["/root/server"]