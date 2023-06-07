# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.1","image:base":"33b47501f2e31be55c1af6b21082de3c3efaae9ec0062beb1d3255a9fd739874","image:golang":"24b9f2aa8b5df78500ac7a816019fd2538905f6c04bea9bb6cebe39f8e1ad1aa"}

FROM reg.c5h.io/golang@sha256:24b9f2aa8b5df78500ac7a816019fd2538905f6c04bea9bb6cebe39f8e1ad1aa AS build

ARG TAG="v4.0.1"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-github /go/src/github.com/greboid/irc-github; \
    cd /go/src/github.com/greboid/irc-github; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/github; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:33b47501f2e31be55c1af6b21082de3c3efaae9ec0062beb1d3255a9fd739874

COPY --from=build /go/src/github.com/greboid/irc-github/main /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
