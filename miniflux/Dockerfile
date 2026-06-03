# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"github:miniflux/v2":"2.3.1","image:base":"00146a62a048620af617bc351091bce6105ca77ea883f65bf45aa22aecf62958","image:golang":"033ea84b4d6c3dd0ab46cc4f0cbf0a596b32f2db8313a44c0ac4bbb290d37f9c"}

FROM reg.c5h.io/golang@sha256:033ea84b4d6c3dd0ab46cc4f0cbf0a596b32f2db8313a44c0ac4bbb290d37f9c AS build

RUN set -eux; \
    git clone --depth=1 -b 2.3.1 --single-branch https://github.com/miniflux/v2 /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:00146a62a048620af617bc351091bce6105ca77ea883f65bf45aa22aecf62958

COPY --from=build /work/out /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
