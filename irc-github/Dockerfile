# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.1","image:base":"93252001f0ac448e4117c7203496cd54f7e8a73d779b9aa8b8f285d6acf31811","image:golang":"9112a9627dc59dcf9783f3043db276b996faeb93eaee4f6a3ac242b17b15c10a"}

FROM reg.c5h.io/golang@sha256:9112a9627dc59dcf9783f3043db276b996faeb93eaee4f6a3ac242b17b15c10a AS build

ARG TAG="v4.0.1"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-github /go/src/github.com/greboid/irc-github; \
    cd /go/src/github.com/greboid/irc-github; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/github; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:93252001f0ac448e4117c7203496cd54f7e8a73d779b9aa8b8f285d6acf31811

COPY --from=build /go/src/github.com/greboid/irc-github/main /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
