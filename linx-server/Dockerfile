# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:https://github.com/cutlerydrawer/linx-server":"v3.4.1","image:base":"164f46a3cc86b9bdd26e51f8d4deb4688e97da19495b83cb1ff3bfc0ad3e31d1","image:golang":"0c7c8b0250d2f903a8c957d72e9a0c9172ac62f04a7c0594649db6b99549ece3"}

FROM reg.c5h.io/golang@sha256:0c7c8b0250d2f903a8c957d72e9a0c9172ac62f04a7c0594649db6b99549ece3 AS build

RUN set -eux; \
    git clone --depth=1 -b v3.4.1 --single-branch https://github.com/cutlerydrawer/linx-server /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data/files /data/meta;



FROM reg.c5h.io/base@sha256:164f46a3cc86b9bdd26e51f8d4deb4688e97da19495b83cb1ff3bfc0ad3e31d1

COPY --from=build /work/out /linx-server
COPY --from=build /notices /notices
ENTRYPOINT ["/linx-server"]


COPY --from=build /data /data

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
