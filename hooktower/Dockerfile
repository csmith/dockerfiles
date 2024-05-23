# Generated from https://github.com/csmith/dockerfiles/blob/master/hooktower/Dockerfile.gotpl
# BOM: {"github:adnanh/webhook":"2.8.1","github:containrrr/watchtower":"v1.7.1","image:base":"b53f7bd15a51e9e9e1c9d1ddea46a2f62223ed7cd8d8fa320c0d6b323f9f3f99","image:golang":"6609c254982953109f3d7c602d3e20165d6857d2773884e9295a9b0d2231672c"}

FROM reg.c5h.io/golang@sha256:6609c254982953109f3d7c602d3e20165d6857d2773884e9295a9b0d2231672c AS build

ARG WATCHTOWER_TAG="v1.7.1"
ARG WEBHOOK_TAG="2.8.1"

RUN set -eux; \
    git clone --depth=1 -b $WATCHTOWER_TAG --single-branch https://github.com/containrrr/watchtower /go/src/github.com/containrrr/watchtower; \
    cd /go/src/github.com/containrrr/watchtower; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static'" -trimpath -o watchtower .; \
    go-licenses save ./... --save_path=/notices; \
    git clone --depth=1 -b $WEBHOOK_TAG --single-branch https://github.com/adnanh/webhook /go/src/github.com/adnanh/webhook; \
    cd /go/src/github.com/adnanh/webhook; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static'" -trimpath -o webhook .; \
    go-licenses save ./... --save_path=/notices/webhook;

FROM reg.c5h.io/base@sha256:b53f7bd15a51e9e9e1c9d1ddea46a2f62223ed7cd8d8fa320c0d6b323f9f3f99

COPY --from=build /go/src/github.com/containrrr/watchtower/watchtower /watchtower
COPY --from=build /go/src/github.com/adnanh/webhook/webhook /webhook
COPY --from=build /notices /notices

ENTRYPOINT ["/webhook"]
