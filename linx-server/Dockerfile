# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:cutlerydrawer/linx-server":"v3.1.0","image:base":"c97e4a9e2aabe7541fb67a9d165b264bfb9d7cc15d954fc48b760e55156666d8","image:golang":"c21f45fbbe31dc43f9197824a2897bb4eca86f45f6a91f265be5995d1cdb3d3c"}

FROM reg.c5h.io/golang@sha256:c21f45fbbe31dc43f9197824a2897bb4eca86f45f6a91f265be5995d1cdb3d3c AS build

RUN set -eux; \
    git clone --depth=1 -b v3.1.0 --single-branch https://github.com/cutlerydrawer/linx-server /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data/files /data/meta;



FROM reg.c5h.io/base@sha256:c97e4a9e2aabe7541fb67a9d165b264bfb9d7cc15d954fc48b760e55156666d8

COPY --from=build /work/out /linx-server
COPY --from=build /notices /notices
ENTRYPOINT ["/linx-server"]


COPY --from=build /data /data

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
