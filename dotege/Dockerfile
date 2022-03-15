# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v1.2.0","image:base":"f47ed6eddc8491fd9683c0bb66751d16b0dd6f9e920d96c9d028ec0001ae5968","image:golang":"fd18f706f641c54676356176aca3a3879490e628c2697866d62557d1edeb3aa6"}

FROM reg.c5h.io/golang@sha256:fd18f706f641c54676356176aca3a3879490e628c2697866d62557d1edeb3aa6 AS build

ARG TAG="v1.2.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/dotege /go/src/github.com/csmith/dotege; \
    cd /go/src/github.com/csmith/dotege; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o main ./cmd/dotege; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Make volume mount points
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/csmith/dotege/main /go/src/github.com/csmith/dotege/templates /notices;

FROM reg.c5h.io/base@sha256:f47ed6eddc8491fd9683c0bb66751d16b0dd6f9e920d96c9d028ec0001ae5968

COPY --from=build /go/src/github.com/csmith/dotege/main /dotege
COPY --from=build /notices /data /
COPY --from=build /go/src/github.com/csmith/dotege/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output

ENTRYPOINT ["/dotege"]
