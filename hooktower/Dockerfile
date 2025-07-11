# Generated from https://github.com/csmith/dockerfiles/blob/master/hooktower/Dockerfile.gotpl
# BOM: {"github:adnanh/webhook":"2.8.2","github:containrrr/watchtower":"v1.7.1","image:base":"c97e4a9e2aabe7541fb67a9d165b264bfb9d7cc15d954fc48b760e55156666d8","image:golang":"d51835bb4b9909f172e3d9b495e5aa14564a70a36fabeafba6832cf0d8ae631b"}

FROM reg.c5h.io/golang@sha256:d51835bb4b9909f172e3d9b495e5aa14564a70a36fabeafba6832cf0d8ae631b AS build


RUN set -eux; \
    git clone --depth=1 -b v1.7.1 --single-branch https://github.com/containrrr/watchtower /go/src/github.com/containrrr/watchtower

WORKDIR /go/src/github.com/containrrr/watchtower


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o watchtower .; \
    go-licenses save ./... --save_path=/out/notices/watchtower;



RUN set -eux; \
    git clone --depth=1 -b 2.8.2 --single-branch https://github.com/adnanh/webhook /go/src/github.com/adnanh/webhook

WORKDIR /go/src/github.com/adnanh/webhook


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o webhook .; \
    go-licenses save ./... --save_path=/out/notices/webhook;


FROM reg.c5h.io/base@sha256:c97e4a9e2aabe7541fb67a9d165b264bfb9d7cc15d954fc48b760e55156666d8

COPY --from=build /go/src/github.com/containrrr/watchtower/watchtower /watchtower
COPY --from=build /go/src/github.com/adnanh/webhook/webhook /webhook
COPY --from=build /out/notices /notices

ENTRYPOINT ["/webhook"]
