# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"b92ffcd70a3cea70972a600a917a409451afe2732844d09e3afe212ad8e041ff","image:golang":"dfd9b80356ce2c039749992c5232aa4d05aadd00d8599357f3762891cf79e3ef"}

FROM reg.c5h.io/golang@sha256:dfd9b80356ce2c039749992c5232aa4d05aadd00d8599357f3762891cf79e3ef AS build

ARG TAG="v2.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-news /go/src/github.com/csmith/irc-news; \
    cd /go/src/github.com/csmith/irc-news; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:b92ffcd70a3cea70972a600a917a409451afe2732844d09e3afe212ad8e041ff

COPY --from=build /go/src/github.com/csmith/irc-news/main /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
