# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:csmith/linx-server":"v3.0.0","image:base":"3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99","image:golang":"2d5332dbd96c2017fc9f481e608f709ca1b3c9a5c5725a66f9dba26ba5432021"}

FROM reg.c5h.io/golang@sha256:2d5332dbd96c2017fc9f481e608f709ca1b3c9a5c5725a66f9dba26ba5432021 AS build

ARG TAG="v3.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/linx-server /go/src/github.com/csmith/linx-server; \
    cd /go/src/github.com/csmith/linx-server; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data;

FROM reg.c5h.io/base@sha256:3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99

COPY --from=build /go/src/github.com/csmith/linx-server/main /linx-server
COPY --from=build /data /data
COPY --from=build /notices /notices

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
