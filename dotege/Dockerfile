# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v2.1.0","image:base":"27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4","image:golang":"2a15c7a9681c2079ad5a8fc578d635e3685b6c441b3c79ba26bbcb1ef902fdf7"}

FROM reg.c5h.io/golang@sha256:2a15c7a9681c2079ad5a8fc578d635e3685b6c441b3c79ba26bbcb1ef902fdf7 AS build

RUN set -eux; \
    git clone --depth=1 -b v2.1.0 --single-branch https://github.com/csmith/dotege /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o out ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data/config /data/certs /data/output;



FROM reg.c5h.io/base@sha256:27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4

COPY --from=build /work/out /dotege
COPY --from=build /notices /notices
ENTRYPOINT ["/dotege"]


COPY --from=build /data /
COPY --from=build /work/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output
