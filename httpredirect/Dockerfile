# Generated from https://github.com/csmith/dockerfiles/blob/master/httpredirect/Dockerfile.gotpl
# BOM: {"github:greboid/httpredirect":"v1.0.3","image:base":"27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4","image:golang":"ae7585d1d2bb33862ddea9d3e10f844e1bdd1ef16bad9e9b9343610f776fda42"}

FROM reg.c5h.io/golang@sha256:ae7585d1d2bb33862ddea9d3e10f844e1bdd1ef16bad9e9b9343610f776fda42 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.0.3 --single-branch https://github.com/greboid/httpredirect /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4

COPY --from=build /work/out /httpredirect
COPY --from=build /notices /notices
ENTRYPOINT ["/httpredirect"]


EXPOSE 8080
