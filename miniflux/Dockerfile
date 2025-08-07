# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"github:miniflux/v2":"2.2.11","image:base":"81ab55b27588f4854075acf470a64bea02b2c1c3535f8067d6f37fe6e9a308b0","image:golang":"a31adfa8c662ab0ca6a11717b6bf73744d04c724b9e86fda87a8d11c58df8fc7"}

FROM reg.c5h.io/golang@sha256:a31adfa8c662ab0ca6a11717b6bf73744d04c724b9e86fda87a8d11c58df8fc7 AS build

RUN set -eux; \
    git clone --depth=1 -b 2.2.11 --single-branch https://github.com/miniflux/v2 /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:81ab55b27588f4854075acf470a64bea02b2c1c3535f8067d6f37fe6e9a308b0

COPY --from=build /work/out /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
