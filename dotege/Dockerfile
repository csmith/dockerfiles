# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:https://github.com/csmith/dotege":"v2.4.2","image:base":"fb20443cb04c169a893b5d736e3f466a42ed299664a12d0dcf68aa37acdc62e9","image:golang":"d6cf3be941b0b3337e885dbfc83e765120838f1464d1422a3b68c3c485ce1be5"}

FROM reg.c5h.io/golang@sha256:d6cf3be941b0b3337e885dbfc83e765120838f1464d1422a3b68c3c485ce1be5 AS build

RUN set -eux; \
    git clone --depth=1 -b v2.4.2 --single-branch https://github.com/csmith/dotege /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o out ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data/config /data/certs /data/output;



FROM reg.c5h.io/base@sha256:fb20443cb04c169a893b5d736e3f466a42ed299664a12d0dcf68aa37acdc62e9

COPY --from=build /work/out /dotege
COPY --from=build /notices /notices
ENTRYPOINT ["/dotege"]


COPY --from=build /data /
COPY --from=build /work/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output
