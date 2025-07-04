# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"github:miniflux/v2":"2.2.10","image:base":"c97e4a9e2aabe7541fb67a9d165b264bfb9d7cc15d954fc48b760e55156666d8","image:golang":"01da54b94fd7f32cc86399c9611d2686e22d0f4dd64a430d5441205e5074930d"}

FROM reg.c5h.io/golang@sha256:01da54b94fd7f32cc86399c9611d2686e22d0f4dd64a430d5441205e5074930d AS build

RUN set -eux; \
    git clone --depth=1 -b 2.2.10 --single-branch https://github.com/miniflux/v2 /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:c97e4a9e2aabe7541fb67a9d165b264bfb9d7cc15d954fc48b760e55156666d8

COPY --from=build /work/out /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
