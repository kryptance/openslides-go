ARG CONTEXT=prod

FROM golang:1.25.4-alpine AS base

## Setup
ARG CONTEXT
WORKDIR /app/openslides-backend-service
ENV APP_CONTEXT=${CONTEXT}

## Install
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

## External Information
EXPOSE 9002

## Command
HEALTHCHECK CMD ["/app/openslides-backend-service/openslides-backend-service", "health"]

# Development Image

FROM base AS dev

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]

CMD CompileDaemon -log-prefix=false -build="go build" -command="./openslides-backend-service"

# Production Image

FROM base AS builder
RUN go build

FROM scratch AS prod

## Setup
ARG CONTEXT
ENV APP_CONTEXT=prod

COPY --from=builder /app/openslides-backend-service/openslides-backend-service /

## External Information
LABEL org.opencontainers.image.title="OpenSlides Backend Service"
LABEL org.opencontainers.image.description="The OpenSlides Backend Service handles actions and presenters."
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/OpenSlides/openslides-backend-service"

EXPOSE 9002

## Command
ENTRYPOINT ["/openslides-backend-service"]

HEALTHCHECK CMD ["/openslides-backend-service", "health"]
