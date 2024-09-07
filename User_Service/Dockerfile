# Rasmiy Golang Docker rasmini yuklab oling
FROM golang:latest

# Konteyner ichida ishchi katalogini o'rnating
WORKDIR /app

# Lokal manba kodini konteyner ichidagi ishchi katalogga nusxalang
COPY . .

# Go dasturini tuzing va 'main' deb nomlangan bajariladigan faylni yarating
RUN go build -o main ./cmd/main.go

# Konteyner ichida 9000 portni oching
EXPOSE 4440

# Konteyner ishga tushirilganda 'main' bajariladigan faylni ishga tushiring
CMD ["./main"]
