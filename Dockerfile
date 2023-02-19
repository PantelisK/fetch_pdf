FROM golang:1.17-alpine AS build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o fetchdummy

FROM scratch
WORKDIR /
COPY --from=build /build/fetchdummy  /
ENTRYPOINT ["/fetchdummy"]