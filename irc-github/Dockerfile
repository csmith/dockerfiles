# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.2","image:base":"cfa1d5f949b495a6e077bdddb1dbbc3d9a9d55036ab1267f4168a6a5c7e7fdce","image:golang":"ec103617eb5bb0afa4d2519d49bd8fdbe55bc462e513822b30a0970391eb3fcc"}

FROM reg.c5h.io/golang@sha256:ec103617eb5bb0afa4d2519d49bd8fdbe55bc462e513822b30a0970391eb3fcc AS build

ARG TAG="v4.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-github /go/src/github.com/greboid/irc-github; \
    cd /go/src/github.com/greboid/irc-github; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/github; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:cfa1d5f949b495a6e077bdddb1dbbc3d9a9d55036ab1267f4168a6a5c7e7fdce

COPY --from=build /go/src/github.com/greboid/irc-github/main /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
