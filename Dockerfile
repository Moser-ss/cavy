FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build .
  
FROM gcr.io/distroless/static-debian11:nonroot
EXPOSE 3000
WORKDIR /app
COPY --from=build /app/cavy /
COPY config.yaml /app/
CMD ["/cavy"]