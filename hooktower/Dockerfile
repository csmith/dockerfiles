# Generated from https://github.com/csmith/dockerfiles/blob/master/hooktower/Dockerfile.gotpl
# BOM: {"github:adnanh/webhook":"2.8.1","github:containrrr/watchtower":"v1.7.1","image:base":"c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d","image:golang":"e8f65f65d61580cd3819e6d4aed29b3eca435d67cf0e09462569d214f9deea5a"}

FROM reg.c5h.io/golang@sha256:e8f65f65d61580cd3819e6d4aed29b3eca435d67cf0e09462569d214f9deea5a AS build

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

FROM reg.c5h.io/base@sha256:c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d

COPY --from=build /go/src/github.com/containrrr/watchtower/watchtower /watchtower
COPY --from=build /go/src/github.com/adnanh/webhook/webhook /webhook
COPY --from=build /notices /notices

ENTRYPOINT ["/webhook"]
