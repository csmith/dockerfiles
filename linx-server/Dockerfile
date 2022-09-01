# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:csmith/linx-server":"v3.0.0","image:base":"dad97084335f003c2ec5251b8f9d22e59704f59c8e2ffcc6108470131341b2df","image:golang":"290e92565f9e74b38a72d91bf56cca4c671b7e6de113b0ae98261b9f80c43d13"}

FROM reg.c5h.io/golang@sha256:290e92565f9e74b38a72d91bf56cca4c671b7e6de113b0ae98261b9f80c43d13 AS build

ARG TAG="v3.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/linx-server /go/src/github.com/csmith/linx-server; \
    cd /go/src/github.com/csmith/linx-server; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data;

FROM reg.c5h.io/base@sha256:dad97084335f003c2ec5251b8f9d22e59704f59c8e2ffcc6108470131341b2df

COPY --from=build /go/src/github.com/csmith/linx-server/main /linx-server
COPY --from=build /data /data
COPY --from=build /notices /notices

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
