# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v2.3.0","image:base":"66af67568b123f6ec0ff7360f9f2f08e274753f885288e5732d6e0c087c1c1dd","image:golang":"79c4a958212bde36a835f9b7c9dd2f721090e1a214699a11cfddb72858bfac6d"}

FROM reg.c5h.io/golang@sha256:79c4a958212bde36a835f9b7c9dd2f721090e1a214699a11cfddb72858bfac6d AS build

RUN set -eux; \
    git clone --depth=1 -b v2.3.0 --single-branch https://github.com/csmith/dotege /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o out ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data/config /data/certs /data/output;



FROM reg.c5h.io/base@sha256:66af67568b123f6ec0ff7360f9f2f08e274753f885288e5732d6e0c087c1c1dd

COPY --from=build /work/out /dotege
COPY --from=build /notices /notices
ENTRYPOINT ["/dotege"]


COPY --from=build /data /
COPY --from=build /work/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output
