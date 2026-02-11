# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:cutlerydrawer/linx-server":"v3.1.0","image:base":"df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a","image:golang":"77752e4278043c67ff36b4335f475e90e765540450f428669249c4441cf287db"}

FROM reg.c5h.io/golang@sha256:77752e4278043c67ff36b4335f475e90e765540450f428669249c4441cf287db AS build

RUN set -eux; \
    git clone --depth=1 -b v3.1.0 --single-branch https://github.com/cutlerydrawer/linx-server /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data/files /data/meta;



FROM reg.c5h.io/base@sha256:df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a

COPY --from=build /work/out /linx-server
COPY --from=build /notices /notices
ENTRYPOINT ["/linx-server"]


COPY --from=build /data /data

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
