package templates

import (
	"fmt"
	"os"

	utils "github.com/DevopsGuyXD/Goku/Utils"
)

func DockerFile() {

	var data string

	file, err := os.Create("./dockerfile")
	utils.CheckForNil(err)
	defer file.Close()

	folder := "Sqlite"

	exists := utils.FolderExists(folder)

	if exists {

		data = `FROM golang:1.24.4 AS builder
WORKDIR /app

    ENV CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64

    COPY go.mod go.sum ./
    RUN go mod download

    COPY . .

    RUN go build -o app

FROM alpine:latest
WORKDIR /app

    RUN addgroup -S appgroup && adduser -S appuser -G appgroup
    USER appuser

    COPY --from=builder /app/app .
    COPY --from=builder /app/.env .
    COPY --from=builder /app/Sqlite ./Sqlite

    EXPOSE 8000

    HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD wget --spider -q http://localhost:8000/health || exit 1

    CMD ["./app"]`
	} else {
		data = `FROM golang:1.23.5 AS builder
WORKDIR /app

    ENV CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64

    COPY go.mod go.sum ./
    RUN go mod download

    COPY . .

    RUN go build -o app

FROM alpine:latest
WORKDIR /app

    RUN addgroup -S appgroup && adduser -S appuser -G appgroup
    USER appuser

    COPY --from=builder /app/app .
    COPY --from=builder /app/.env .

    EXPOSE 8000

    HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD wget --spider -q http://localhost:8000/health || exit 1

    CMD ["./app"]`
	}
	utils.WriteFile(file, data)

	fmt.Println("\nAdded dockerfile ✅ ")
}
