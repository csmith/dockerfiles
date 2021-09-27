# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:csmith/linx-server":"v2.3.9","image:base":"sha256:daa3cfedaabdc411afb9e58693526d5e8d283fbc767abc6c2bcef929ea782566","image:golang":"sha256:2495578a7ef96327cfc4b2971f4264b0006c18d2494a015a03ab03476f86ca70"}

FROM reg.c5h.io/golang@sha256:2495578a7ef96327cfc4b2971f4264b0006c18d2494a015a03ab03476f86ca70 AS build

ARG TAG="v2.3.9"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/linx-server /go/src/github.com/csmith/linx-server; \
    cd /go/src/github.com/csmith/linx-server; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/csmith/linx-server/main;

FROM reg.c5h.io/base@sha256:daa3cfedaabdc411afb9e58693526d5e8d283fbc767abc6c2bcef929ea782566

COPY --from=build /go/src/github.com/csmith/linx-server/main /linx-server
COPY --from=build /data /data
COPY --from=build /notices /notices

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
