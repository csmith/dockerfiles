# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v1.3.2","image:base":"fa9cdd83581affde03cb2234170235ff287588e4da3786133d72846838210621","image:golang":"7785c8e5e1fd85efb640224f81d501d9b8d425b1fafeed4e535c5711cd05119d"}

FROM reg.c5h.io/golang@sha256:7785c8e5e1fd85efb640224f81d501d9b8d425b1fafeed4e535c5711cd05119d AS build

ARG TAG="v1.3.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/dotege /go/src/github.com/csmith/dotege; \
    cd /go/src/github.com/csmith/dotege; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o main ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices; \
    # Make volume mount points
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data;

FROM reg.c5h.io/base@sha256:fa9cdd83581affde03cb2234170235ff287588e4da3786133d72846838210621

COPY --from=build /go/src/github.com/csmith/dotege/main /dotege
COPY --from=build /notices /data /
COPY --from=build /go/src/github.com/csmith/dotege/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output

ENTRYPOINT ["/dotege"]
