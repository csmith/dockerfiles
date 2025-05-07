# Generated from https://github.com/csmith/dockerfiles/blob/master/httpredirect/Dockerfile.gotpl
# BOM: {"github:greboid/httpredirect":"v1.0.3","image:base":"4fea2559a3c315f01c6241c689da902344f4694dbdbdab4f5c0a1d24a8a78363","image:golang":"d62ed02ee4a78007acc79a44388eb350701222f1aa76483fcfe363958d5845ec"}

FROM reg.c5h.io/golang@sha256:d62ed02ee4a78007acc79a44388eb350701222f1aa76483fcfe363958d5845ec AS build

RUN set -eux; \
    git clone --depth=1 -b v1.0.3 --single-branch https://github.com/greboid/httpredirect /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:4fea2559a3c315f01c6241c689da902344f4694dbdbdab4f5c0a1d24a8a78363

COPY --from=build /work/out /httpredirect
COPY --from=build /notices /notices
ENTRYPOINT ["/httpredirect"]


EXPOSE 8080
