# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:https://github.com/greboid/irc-github":"v4.0.5","image:base":"091cf7586a6b9d750c6e29255c7231896ce1ec4ef852f351068072ca00216b47","image:golang":"5c4ad85bffae07edfdae6e170e1d6550833106a17297d826c30e9a3a1556b65c"}

FROM reg.c5h.io/golang@sha256:5c4ad85bffae07edfdae6e170e1d6550833106a17297d826c30e9a3a1556b65c AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.5 --single-branch https://github.com/greboid/irc-github /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/github; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:091cf7586a6b9d750c6e29255c7231896ce1ec4ef852f351068072ca00216b47

COPY --from=build /work/out /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
