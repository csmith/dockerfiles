# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"github:https://github.com/miniflux/v2":"2.3.3","image:base":"996f481af28d13a3c5069644629fbb02aace36a81a912b688db0771db8669424","image:golang":"2486b1a3072181a988fd9f42069d094bf3056ff19d675520656adbd623a5278e"}

FROM reg.c5h.io/golang@sha256:2486b1a3072181a988fd9f42069d094bf3056ff19d675520656adbd623a5278e AS build

RUN set -eux; \
    git clone --depth=1 -b 2.3.3 --single-branch https://github.com/miniflux/v2 /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:996f481af28d13a3c5069644629fbb02aace36a81a912b688db0771db8669424

COPY --from=build /work/out /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
