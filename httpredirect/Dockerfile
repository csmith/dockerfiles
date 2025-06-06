# Generated from https://github.com/csmith/dockerfiles/blob/master/httpredirect/Dockerfile.gotpl
# BOM: {"github:greboid/httpredirect":"v1.0.3","image:base":"c97e4a9e2aabe7541fb67a9d165b264bfb9d7cc15d954fc48b760e55156666d8","image:golang":"075478d1b4114af160e921b8673d3419a444a35b947d86856d7b3c4a74a8f445"}

FROM reg.c5h.io/golang@sha256:075478d1b4114af160e921b8673d3419a444a35b947d86856d7b3c4a74a8f445 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.0.3 --single-branch https://github.com/greboid/httpredirect /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:c97e4a9e2aabe7541fb67a9d165b264bfb9d7cc15d954fc48b760e55156666d8

COPY --from=build /work/out /httpredirect
COPY --from=build /notices /notices
ENTRYPOINT ["/httpredirect"]


EXPOSE 8080
