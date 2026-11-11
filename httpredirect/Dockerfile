# Generated from https://github.com/csmith/dockerfiles/blob/master/httpredirect/Dockerfile.gotpl
# BOM: {"github:https://github.com/greboid/httpredirect":"v1.2.0","image:base":"fb20443cb04c169a893b5d736e3f466a42ed299664a12d0dcf68aa37acdc62e9","image:golang":"87be3283dc0e0b30e0e53bc9b791a09e466209855746c6981e2062626cee868c"}

FROM reg.c5h.io/golang@sha256:87be3283dc0e0b30e0e53bc9b791a09e466209855746c6981e2062626cee868c AS build

RUN set -eux; \
    git clone --depth=1 -b v1.2.0 --single-branch https://github.com/greboid/httpredirect /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:fb20443cb04c169a893b5d736e3f466a42ed299664a12d0dcf68aa37acdc62e9

COPY --from=build /work/out /httpredirect
COPY --from=build /notices /notices
ENTRYPOINT ["/httpredirect"]


EXPOSE 8080
