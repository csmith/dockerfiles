# Generated from https://github.com/csmith/dockerfiles/blob/master/linx-server/Dockerfile.gotpl
# BOM: {"github:csmith/linx-server":"v2.3.9","image:base":"sha256:bba8c74ee389cc37da171979067085b53a8561caa66f49719db79e2fbb5d96cd","image:golang":"sha256:d10adab145c50608e7fb95c0923830744dd674c981e383ac9d2300813f70e670"}

FROM reg.c5h.io/golang@sha256:d10adab145c50608e7fb95c0923830744dd674c981e383ac9d2300813f70e670 AS build

ARG TAG="v2.3.9"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/linx-server /go/src/github.com/csmith/linx-server; \
    cd /go/src/github.com/csmith/linx-server; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    mkdir -p /data/files /data/meta; \
    chown -R 65532:65532 /data; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/csmith/linx-server/main;

FROM reg.c5h.io/base@sha256:bba8c74ee389cc37da171979067085b53a8561caa66f49719db79e2fbb5d96cd

COPY --from=build /go/src/github.com/csmith/linx-server/main /linx-server
COPY --from=build /data /data
COPY --from=build /notices /notices

VOLUME ["/data/files", "/data/meta"]
EXPOSE 8080
ENTRYPOINT ["/linx-server", "-bind=0.0.0.0:8080", "-filespath=/data/files/", "-metapath=/data/meta/"]
CMD ["-sitename=linx", "-allowhotlink"]
