FROM golang AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /backend

# Production
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /backend /backend
EXPOSE 8081
USER nonroot:nonroot
ENV BROKER_TOKEN=FIIKZQz9LdJsQkrttw5nBQ
ENTRYPOINT ["/backend"]
