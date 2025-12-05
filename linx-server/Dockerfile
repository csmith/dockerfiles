# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:cutlerydrawer/linx-server":"v3.1.0","image:base":"46ea216adf9eb9615e03305be16d497ef5e1f3d936f3e986a8e4d25e599f27f7","image:golang":"18a4339b2b4f079668919a005c982400f8cf8ce18f4752336395732498bddeef"}

FROM reg.c5h.io/golang@sha256:18a4339b2b4f079668919a005c982400f8cf8ce18f4752336395732498bddeef AS build

RUN set -eux; \
    git clone --depth=1 -b v3.1.0 --single-branch https://github.com/cutlerydrawer/linx-server /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data/files /data/meta;



FROM reg.c5h.io/base@sha256:46ea216adf9eb9615e03305be16d497ef5e1f3d936f3e986a8e4d25e599f27f7

COPY --from=build /work/out /linx-server
COPY --from=build /notices /notices
ENTRYPOINT ["/linx-server"]


COPY --from=build /data /data

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
