FROM golang:1.19.1-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

WORKDIR /app/cmd/web

# Build
RUN go build -o /snippetstore-docker

EXPOSE 4000

# Run
CMD [ "/snippetstore-docker" ]