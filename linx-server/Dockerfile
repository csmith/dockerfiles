# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:csmith/linx-server":"v3.0.0","image:base":"06668724f8daa1866a72d07c44d9c0fe5c42bfd277d47fecc8e4ed037a5d0e84","image:golang":"8cca8e81ef263f66855ba7ad711e884e03e7670276e1cd2011d8f0f3f0eec7e0"}

FROM reg.c5h.io/golang@sha256:8cca8e81ef263f66855ba7ad711e884e03e7670276e1cd2011d8f0f3f0eec7e0 AS build

ARG TAG="v3.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/linx-server /go/src/github.com/csmith/linx-server; \
    cd /go/src/github.com/csmith/linx-server; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data;

FROM reg.c5h.io/base@sha256:06668724f8daa1866a72d07c44d9c0fe5c42bfd277d47fecc8e4ed037a5d0e84

COPY --from=build /go/src/github.com/csmith/linx-server/main /linx-server
COPY --from=build /data /data
COPY --from=build /notices /notices

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
