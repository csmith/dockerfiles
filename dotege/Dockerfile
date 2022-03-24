# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v1.2.0","image:base":"f85825cc6e579732f339638d1cd16d87f0673759127df47e8bcd3586412acb5d","image:golang":"aaf0cd2c7afeee9750b1b891ec05210afa4a508ab9525cd102e5df363b7bcb2a"}

FROM reg.c5h.io/golang@sha256:aaf0cd2c7afeee9750b1b891ec05210afa4a508ab9525cd102e5df363b7bcb2a AS build

ARG TAG="v1.2.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/dotege /go/src/github.com/csmith/dotege; \
    cd /go/src/github.com/csmith/dotege; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o main ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices; \
    # Make volume mount points
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data;

FROM reg.c5h.io/base@sha256:f85825cc6e579732f339638d1cd16d87f0673759127df47e8bcd3586412acb5d

COPY --from=build /go/src/github.com/csmith/dotege/main /dotege
COPY --from=build /notices /data /
COPY --from=build /go/src/github.com/csmith/dotege/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output

ENTRYPOINT ["/dotege"]
