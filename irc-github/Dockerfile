# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.1","image:base":"5d2c54fa31ef9b63b0499aabf0dd0091eee215ae556dd89f68189e7894cef381","image:golang":"d8fd3cc3b73ff5eb50c1e125e6fbf69775e75c4b881a89c035bc0a60372f3b3f"}

FROM reg.c5h.io/golang@sha256:d8fd3cc3b73ff5eb50c1e125e6fbf69775e75c4b881a89c035bc0a60372f3b3f AS build

ARG TAG="v4.0.1"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-github /go/src/github.com/greboid/irc-github; \
    cd /go/src/github.com/greboid/irc-github; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/github; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:5d2c54fa31ef9b63b0499aabf0dd0091eee215ae556dd89f68189e7894cef381

COPY --from=build /go/src/github.com/greboid/irc-github/main /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
