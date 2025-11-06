# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"github:miniflux/v2":"2.2.14","image:base":"a09ba1c5a9b2ec41293a8fa0cef935f6ab81978fcbb62271f490ebb07bbf6024","image:golang":"f08235e73a508e090d366ed3b7db00ec9b11c8bb53af3fe9d0b79ade8f886216"}

FROM reg.c5h.io/golang@sha256:f08235e73a508e090d366ed3b7db00ec9b11c8bb53af3fe9d0b79ade8f886216 AS build

RUN set -eux; \
    git clone --depth=1 -b 2.2.14 --single-branch https://github.com/miniflux/v2 /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:a09ba1c5a9b2ec41293a8fa0cef935f6ab81978fcbb62271f490ebb07bbf6024

COPY --from=build /work/out /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
