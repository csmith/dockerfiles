# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v2.1.0","image:base":"3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99","image:golang":"fb31d040911eec2e5068e1b3283d39c60b916549a52671b1aabe091488fb66d0"}

FROM reg.c5h.io/golang@sha256:fb31d040911eec2e5068e1b3283d39c60b916549a52671b1aabe091488fb66d0 AS build

ARG TAG="v2.1.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/dotege /go/src/github.com/csmith/dotege; \
    cd /go/src/github.com/csmith/dotege; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o main ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices; \
    # Make volume mount points
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data;

FROM reg.c5h.io/base@sha256:3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99

COPY --from=build /go/src/github.com/csmith/dotege/main /dotege
COPY --from=build /notices /data /
COPY --from=build /go/src/github.com/csmith/dotege/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output

ENTRYPOINT ["/dotege"]
