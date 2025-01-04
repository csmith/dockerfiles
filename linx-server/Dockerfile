# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:cutlerydrawer/linx-server":"v3.0.0","image:base":"a59bc827c80c4dbec2473c6b1a6f72b5b015d44e37981b112b9a60cb3c5180af","image:golang":"dd6035aab3b9568fd9cf56c1bfe223c945f1f7ed64f59c65230423bfb6fd136d"}

FROM reg.c5h.io/golang@sha256:dd6035aab3b9568fd9cf56c1bfe223c945f1f7ed64f59c65230423bfb6fd136d AS build

RUN set -eux; \
    git clone --depth=1 -b v3.0.0 --single-branch https://github.com/cutlerydrawer/linx-server /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data/files /data/meta;



FROM reg.c5h.io/base@sha256:a59bc827c80c4dbec2473c6b1a6f72b5b015d44e37981b112b9a60cb3c5180af

COPY --from=build /work/out /linx-server
COPY --from=build /notices /notices
ENTRYPOINT ["/linx-server"]


COPY --from=build /data /data

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
