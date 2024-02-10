# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.2","image:base":"fd06c9da248d637dfa37579addc089c993016cf3e21ee1142b5aea23095ad4f9","image:golang":"39d25f183dccb811df2f7cfe55b22c1cc5575ef8ac98116541a2a8edfea204ce"}

FROM reg.c5h.io/golang@sha256:39d25f183dccb811df2f7cfe55b22c1cc5575ef8ac98116541a2a8edfea204ce AS build

ARG TAG="v4.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-github /go/src/github.com/greboid/irc-github; \
    cd /go/src/github.com/greboid/irc-github; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/github; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:fd06c9da248d637dfa37579addc089c993016cf3e21ee1142b5aea23095ad4f9

COPY --from=build /go/src/github.com/greboid/irc-github/main /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
