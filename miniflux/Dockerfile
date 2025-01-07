# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"github:miniflux/v2":"2.2.4","image:base":"a59bc827c80c4dbec2473c6b1a6f72b5b015d44e37981b112b9a60cb3c5180af","image:golang":"69bf99d272afa648a5758d3f6e435854fa6b6962df9418495dc5c2291e38f0df"}

FROM reg.c5h.io/golang@sha256:69bf99d272afa648a5758d3f6e435854fa6b6962df9418495dc5c2291e38f0df AS build

RUN set -eux; \
    git clone --depth=1 -b 2.2.4 --single-branch https://github.com/miniflux/v2 /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:a59bc827c80c4dbec2473c6b1a6f72b5b015d44e37981b112b9a60cb3c5180af

COPY --from=build /work/out /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
