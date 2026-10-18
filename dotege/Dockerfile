# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:https://github.com/csmith/dotege":"v2.4.0","image:base":"fbcd09463f4b0824d709a18598637c2f7f7266d0591ca7b61311b2a12cca6f11","image:golang":"2a596927c391e5688e97229dbddd5d65429b1a7459a530477522f8fd1d8d29bd"}

FROM reg.c5h.io/golang@sha256:2a596927c391e5688e97229dbddd5d65429b1a7459a530477522f8fd1d8d29bd AS build

RUN set -eux; \
    git clone --depth=1 -b v2.4.0 --single-branch https://github.com/csmith/dotege /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o out ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data/config /data/certs /data/output;



FROM reg.c5h.io/base@sha256:fbcd09463f4b0824d709a18598637c2f7f7266d0591ca7b61311b2a12cca6f11

COPY --from=build /work/out /dotege
COPY --from=build /notices /notices
ENTRYPOINT ["/dotege"]


COPY --from=build /data /
COPY --from=build /work/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output
