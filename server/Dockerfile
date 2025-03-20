# Builder stage
FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main main.go

# -----------------------------------------------

# Runner stage to run the binary
FROM scratch AS runner

# Copy the binary from builder stage to this stage
COPY --from=builder ./app/main ./main

EXPOSE 8998

ENTRYPOINT [ "./main" ]
