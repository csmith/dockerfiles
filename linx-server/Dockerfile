# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:csmith/linx-server":"v2.3.9","image:base":"sha256:f7f27db7afb58bae23ad902072228f1b090b878af303d9f63bd2c1526b9b4f53","image:golang":"sha256:39697ca81312ea9a12b8ba883facf35ffd320ac0dcd8ae977fd03ab25c461fca"}

FROM reg.c5h.io/golang@sha256:39697ca81312ea9a12b8ba883facf35ffd320ac0dcd8ae977fd03ab25c461fca AS build

ARG TAG="v2.3.9"

RUN set -eux; \
    apk add --no-cache git; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/linx-server /go/src/github.com/csmith/linx-server; \
    cd /go/src/github.com/csmith/linx-server; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/csmith/linx-server/main;

FROM reg.c5h.io/base@sha256:f7f27db7afb58bae23ad902072228f1b090b878af303d9f63bd2c1526b9b4f53

COPY --from=build /go/src/github.com/csmith/linx-server/main /linx-server
COPY --from=build /data /data

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
