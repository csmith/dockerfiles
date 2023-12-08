# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.2","image:base":"67c850424b327e8a75333dbe8fe00ffa4e37f10c114efa3c6884d3e687d6cb47","image:golang":"72e140a52eb30250acb9f89a7ca4c1ce30ffcffb46d5ba1070f5f73afdf7bf78"}

FROM reg.c5h.io/golang@sha256:72e140a52eb30250acb9f89a7ca4c1ce30ffcffb46d5ba1070f5f73afdf7bf78 AS build

ARG TAG="v4.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-github /go/src/github.com/greboid/irc-github; \
    cd /go/src/github.com/greboid/irc-github; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/github; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:67c850424b327e8a75333dbe8fe00ffa4e37f10c114efa3c6884d3e687d6cb47

COPY --from=build /go/src/github.com/greboid/irc-github/main /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
