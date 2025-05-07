# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:cutlerydrawer/linx-server":"v3.0.0","image:base":"4fea2559a3c315f01c6241c689da902344f4694dbdbdab4f5c0a1d24a8a78363","image:golang":"d62ed02ee4a78007acc79a44388eb350701222f1aa76483fcfe363958d5845ec"}

FROM reg.c5h.io/golang@sha256:d62ed02ee4a78007acc79a44388eb350701222f1aa76483fcfe363958d5845ec AS build

RUN set -eux; \
    git clone --depth=1 -b v3.0.0 --single-branch https://github.com/cutlerydrawer/linx-server /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data/files /data/meta;



FROM reg.c5h.io/base@sha256:4fea2559a3c315f01c6241c689da902344f4694dbdbdab4f5c0a1d24a8a78363

COPY --from=build /work/out /linx-server
COPY --from=build /notices /notices
ENTRYPOINT ["/linx-server"]


COPY --from=build /data /data

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
