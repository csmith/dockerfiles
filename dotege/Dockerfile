# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v1.3.1","image:base":"45ece244ca0ace29f890ad27114ff83a36c00e74e6cbec3d1e226b7e62d0de7a","image:golang":"fec1aa63edf7ff6121cff227ea449b450caedf2f1bd8756dcb0748191c55f5c5"}

FROM reg.c5h.io/golang@sha256:fec1aa63edf7ff6121cff227ea449b450caedf2f1bd8756dcb0748191c55f5c5 AS build

ARG TAG="v1.3.1"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/dotege /go/src/github.com/csmith/dotege; \
    cd /go/src/github.com/csmith/dotege; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o main ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices; \
    # Make volume mount points
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data;

FROM reg.c5h.io/base@sha256:45ece244ca0ace29f890ad27114ff83a36c00e74e6cbec3d1e226b7e62d0de7a

COPY --from=build /go/src/github.com/csmith/dotege/main /dotege
COPY --from=build /notices /data /
COPY --from=build /go/src/github.com/csmith/dotege/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output

ENTRYPOINT ["/dotege"]
