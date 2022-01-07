# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:csmith/linx-server":"v2.3.9","image:base":"e3b3d5f2e5fa7e5cbe763581a8da6b56c4498cb2107b9e385acc8f13d4de5903","image:golang":"52968a8b807cc37629d7bb688498c02998d5d2ef70caba40d5a048dbde4b7edf"}

FROM reg.c5h.io/golang@sha256:52968a8b807cc37629d7bb688498c02998d5d2ef70caba40d5a048dbde4b7edf AS build

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

FROM reg.c5h.io/base@sha256:e3b3d5f2e5fa7e5cbe763581a8da6b56c4498cb2107b9e385acc8f13d4de5903

COPY --from=build /go/src/github.com/csmith/linx-server/main /linx-server
COPY --from=build /data /data
COPY --from=build /notices /notices

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
