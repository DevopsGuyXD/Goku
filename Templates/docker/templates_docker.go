package templates_docker

import (
	utils "github.com/DevopsGuyXD/Goku/Utils"
)

// ============================================================================ DOCKER FILE
func DockerFile() string {

	var data string

	utils.Create_File([]string{"./dockerfile"})
	exists := utils.Folder_Exists("Sqlite")

	if exists {

		data =
			`FROM golang:latest AS builder
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

	// utils.Write_File(file, data)

	// fmt.Println("\nAdded dockerfile ")

	return data
}
