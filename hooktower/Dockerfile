# Generated from https://github.com/csmith/dockerfiles/blob/master/hooktower/Dockerfile.gotpl
# BOM: {"github:adnanh/webhook":"2.8.1","github:containrrr/watchtower":"v1.7.1","image:base":"fd06c9da248d637dfa37579addc089c993016cf3e21ee1142b5aea23095ad4f9","image:golang":"e9bc02ea14a83a22df783d2d6244c8fa4349ccc91d4d202784c2b569b747ab65"}

FROM reg.c5h.io/golang@sha256:e9bc02ea14a83a22df783d2d6244c8fa4349ccc91d4d202784c2b569b747ab65 AS build

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

FROM reg.c5h.io/base@sha256:fd06c9da248d637dfa37579addc089c993016cf3e21ee1142b5aea23095ad4f9

COPY --from=build /go/src/github.com/containrrr/watchtower/watchtower /watchtower
COPY --from=build /go/src/github.com/adnanh/webhook/webhook /webhook
COPY --from=build /notices /notices

ENTRYPOINT ["/webhook"]
