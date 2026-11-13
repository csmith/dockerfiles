# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"github:https://github.com/miniflux/v2":"2.3.5","image:base":"fb20443cb04c169a893b5d736e3f466a42ed299664a12d0dcf68aa37acdc62e9","image:golang":"638a1c914c1117254a593a283f0e51d6e425b007d1cefe9d7efac14779f622ef"}

FROM reg.c5h.io/golang@sha256:638a1c914c1117254a593a283f0e51d6e425b007d1cefe9d7efac14779f622ef AS build

RUN set -eux; \
    git clone --depth=1 -b 2.3.5 --single-branch https://github.com/miniflux/v2 /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:fb20443cb04c169a893b5d736e3f466a42ed299664a12d0dcf68aa37acdc62e9

COPY --from=build /work/out /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
