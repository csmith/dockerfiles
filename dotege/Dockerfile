# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:https://github.com/csmith/dotege":"v2.4.1","image:base":"091cf7586a6b9d750c6e29255c7231896ce1ec4ef852f351068072ca00216b47","image:golang":"bd172a7af0e1c90db90419b5c1dcab346836d53029ca9574bb16127bb858c47f"}

FROM reg.c5h.io/golang@sha256:bd172a7af0e1c90db90419b5c1dcab346836d53029ca9574bb16127bb858c47f AS build

RUN set -eux; \
    git clone --depth=1 -b v2.4.1 --single-branch https://github.com/csmith/dotege /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o out ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data/config /data/certs /data/output;



FROM reg.c5h.io/base@sha256:091cf7586a6b9d750c6e29255c7231896ce1ec4ef852f351068072ca00216b47

COPY --from=build /work/out /dotege
COPY --from=build /notices /notices
ENTRYPOINT ["/dotege"]


COPY --from=build /data /
COPY --from=build /work/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output
