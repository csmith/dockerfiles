# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"github:miniflux/v2":"2.2.13","image:base":"81ab55b27588f4854075acf470a64bea02b2c1c3535f8067d6f37fe6e9a308b0","image:golang":"d6762517c8ca7ab1dde07a3219a943f56648ad67678bf8c1bc5fd23c7e373de3"}

FROM reg.c5h.io/golang@sha256:d6762517c8ca7ab1dde07a3219a943f56648ad67678bf8c1bc5fd23c7e373de3 AS build

RUN set -eux; \
    git clone --depth=1 -b 2.2.13 --single-branch https://github.com/miniflux/v2 /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:81ab55b27588f4854075acf470a64bea02b2c1c3535f8067d6f37fe6e9a308b0

COPY --from=build /work/out /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
