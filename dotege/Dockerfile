# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v1.0.0","image:base":"sha256:0afb94d65692fdb671d51286d0bd62d65a6d3c3d3be755de0d15a6454c946430","image:golang":"sha256:d10adab145c50608e7fb95c0923830744dd674c981e383ac9d2300813f70e670"}

FROM reg.c5h.io/golang@sha256:d10adab145c50608e7fb95c0923830744dd674c981e383ac9d2300813f70e670 AS build

ARG TAG="v1.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/dotege /go/src/github.com/csmith/dotege; \
    cd /go/src/github.com/csmith/dotege; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -o main .; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Make volume mount points
    mkdir -p /data/{config,certs,output}; \
    chown -R 65532:65532 /data; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/csmith/dotege/{main,templates} /notices;

FROM reg.c5h.io/base@sha256:0afb94d65692fdb671d51286d0bd62d65a6d3c3d3be755de0d15a6454c946430

COPY --from=build /go/src/github.com/csmith/dotege/main /dotege
COPY --from=build /notices /data /
COPY --from=build /go/src/github.com/csmith/dotege/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output

ENTRYPOINT ["/dotege"]
