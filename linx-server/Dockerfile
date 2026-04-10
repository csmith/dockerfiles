# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:cutlerydrawer/linx-server":"v3.4.1","image:base":"46cb85d037365ac7498d4cefa5882d5a6242d9907e66abcde3a2c2fc62a48987","image:golang":"d2c7d4b907467670c9dc335754a890a061167ede406e99020e7d7e1a43eada6a"}

FROM reg.c5h.io/golang@sha256:d2c7d4b907467670c9dc335754a890a061167ede406e99020e7d7e1a43eada6a AS build

RUN set -eux; \
    git clone --depth=1 -b v3.4.1 --single-branch https://github.com/cutlerydrawer/linx-server /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;



RUN set -eux; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data/files /data/meta;



FROM reg.c5h.io/base@sha256:46cb85d037365ac7498d4cefa5882d5a6242d9907e66abcde3a2c2fc62a48987

COPY --from=build /work/out /linx-server
COPY --from=build /notices /notices
ENTRYPOINT ["/linx-server"]


COPY --from=build /data /data

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
