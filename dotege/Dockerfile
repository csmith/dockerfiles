# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v2.3.0","image:base":"22cd0188413016b6119bf80c3e3be20bf03b54e9ce6b9ff1576ab07f64f6da66","image:golang":"5e07f9b2210c6981dca7906be8aa31aed0b75ae0bfc0d776b74e52b5b86b7736"}

FROM reg.c5h.io/golang@sha256:5e07f9b2210c6981dca7906be8aa31aed0b75ae0bfc0d776b74e52b5b86b7736 AS build

RUN set -eux; \
    git clone --depth=1 -b v2.3.0 --single-branch https://github.com/csmith/dotege /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o out ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data/config /data/certs /data/output;



FROM reg.c5h.io/base@sha256:22cd0188413016b6119bf80c3e3be20bf03b54e9ce6b9ff1576ab07f64f6da66

COPY --from=build /work/out /dotege
COPY --from=build /notices /notices
ENTRYPOINT ["/dotege"]


COPY --from=build /data /
COPY --from=build /work/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output
