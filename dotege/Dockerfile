# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v1.1.0","image:base":"b2e73bfd5ae1ce12341dd36f24573d1316bf5172809c65dd669ff2720582967f","image:golang":"beba80acf4c14def39eb627a988efb23102fb66a01b952445c7e77aea5f62bc4"}

FROM reg.c5h.io/golang@sha256:beba80acf4c14def39eb627a988efb23102fb66a01b952445c7e77aea5f62bc4 AS build

ARG TAG="v1.1.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/dotege /go/src/github.com/csmith/dotege; \
    cd /go/src/github.com/csmith/dotege; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o main .; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Make volume mount points
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/csmith/dotege/main /go/src/github.com/csmith/dotege/templates /notices;

FROM reg.c5h.io/base@sha256:b2e73bfd5ae1ce12341dd36f24573d1316bf5172809c65dd669ff2720582967f

COPY --from=build /go/src/github.com/csmith/dotege/main /dotege
COPY --from=build /notices /data /
COPY --from=build /go/src/github.com/csmith/dotege/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output

ENTRYPOINT ["/dotege"]
