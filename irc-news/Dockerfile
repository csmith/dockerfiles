# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"dad97084335f003c2ec5251b8f9d22e59704f59c8e2ffcc6108470131341b2df","image:golang":"7f3e1f7047e21148ccfa49f487a71483f39023f756874f6e3cff3e0d2b8a045c"}

FROM reg.c5h.io/golang@sha256:7f3e1f7047e21148ccfa49f487a71483f39023f756874f6e3cff3e0d2b8a045c AS build

ARG TAG="v2.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-news /go/src/github.com/csmith/irc-news; \
    cd /go/src/github.com/csmith/irc-news; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:dad97084335f003c2ec5251b8f9d22e59704f59c8e2ffcc6108470131341b2df

COPY --from=build /go/src/github.com/csmith/irc-news/main /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
