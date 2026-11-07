# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:https://github.com/cutlerydrawer/linx-server":"v3.4.1","image:base":"fb20443cb04c169a893b5d736e3f466a42ed299664a12d0dcf68aa37acdc62e9","image:golang":"d6cf3be941b0b3337e885dbfc83e765120838f1464d1422a3b68c3c485ce1be5"}

FROM reg.c5h.io/golang@sha256:d6cf3be941b0b3337e885dbfc83e765120838f1464d1422a3b68c3c485ce1be5 AS build

RUN set -eux; \
    git clone --depth=1 -b v3.4.1 --single-branch https://github.com/cutlerydrawer/linx-server /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data/files /data/meta;



FROM reg.c5h.io/base@sha256:fb20443cb04c169a893b5d736e3f466a42ed299664a12d0dcf68aa37acdc62e9

COPY --from=build /work/out /linx-server
COPY --from=build /notices /notices
ENTRYPOINT ["/linx-server"]


COPY --from=build /data /data

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
