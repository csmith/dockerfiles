# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v2.3.0","image:base":"996f481af28d13a3c5069644629fbb02aace36a81a912b688db0771db8669424","image:golang":"0e37894507254f68c0fb756ee281a65a5432789184cb0de26d91c8f7d2f9f709"}

FROM reg.c5h.io/golang@sha256:0e37894507254f68c0fb756ee281a65a5432789184cb0de26d91c8f7d2f9f709 AS build

RUN set -eux; \
    git clone --depth=1 -b v2.3.0 --single-branch https://github.com/csmith/dotege /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o out ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data/config /data/certs /data/output;



FROM reg.c5h.io/base@sha256:996f481af28d13a3c5069644629fbb02aace36a81a912b688db0771db8669424

COPY --from=build /work/out /dotege
COPY --from=build /notices /notices
ENTRYPOINT ["/dotege"]


COPY --from=build /data /
COPY --from=build /work/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output
