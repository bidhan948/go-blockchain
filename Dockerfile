# 1. Base image with Go 1.22 (Alpine = tiny)
FROM golang:1.22-alpine

# 2. Set working directory inside the container
WORKDIR /app

# 3. Copy source code into the image
COPY . .

# 4. Compile the binary
RUN go build -o hello main.go

# 5. Set the default command (prints once and exits)
CMD ["./hello"]


# NEW – keep container alive, you can attach anytime
CMD ["sleep", "infinity"]