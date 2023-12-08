# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v1.3.1","image:base":"67c850424b327e8a75333dbe8fe00ffa4e37f10c114efa3c6884d3e687d6cb47","image:golang":"72e140a52eb30250acb9f89a7ca4c1ce30ffcffb46d5ba1070f5f73afdf7bf78"}

FROM reg.c5h.io/golang@sha256:72e140a52eb30250acb9f89a7ca4c1ce30ffcffb46d5ba1070f5f73afdf7bf78 AS build

ARG TAG="v1.3.1"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/dotege /go/src/github.com/csmith/dotege; \
    cd /go/src/github.com/csmith/dotege; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o main ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices; \
    # Make volume mount points
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data;

FROM reg.c5h.io/base@sha256:67c850424b327e8a75333dbe8fe00ffa4e37f10c114efa3c6884d3e687d6cb47

COPY --from=build /go/src/github.com/csmith/dotege/main /dotege
COPY --from=build /notices /data /
COPY --from=build /go/src/github.com/csmith/dotege/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output

ENTRYPOINT ["/dotege"]
