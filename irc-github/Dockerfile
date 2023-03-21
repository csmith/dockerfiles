# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.1","image:base":"45ece244ca0ace29f890ad27114ff83a36c00e74e6cbec3d1e226b7e62d0de7a","image:golang":"ad26d081d498687f8c1e4874a6da33d2ecbdaa7c62113e6c15ce694740f669dd"}

FROM reg.c5h.io/golang@sha256:ad26d081d498687f8c1e4874a6da33d2ecbdaa7c62113e6c15ce694740f669dd AS build

ARG TAG="v4.0.1"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-github /go/src/github.com/greboid/irc-github; \
    cd /go/src/github.com/greboid/irc-github; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/github; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:45ece244ca0ace29f890ad27114ff83a36c00e74e6cbec3d1e226b7e62d0de7a

COPY --from=build /go/src/github.com/greboid/irc-github/main /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
