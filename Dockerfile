FROM golang AS build

WORKDIR /app

COPY . ./
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build -p $(nproc) -ldflags="-w -s" -o bin/ ./ofly

FROM alpine

RUN apk add -U --no-cache ca-certificates

COPY --from=build /app/bin/ /bin/

VOLUME /data
ENV OFLY_PATH=/data

ENTRYPOINT ["/bin/ofly"]
