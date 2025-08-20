# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.2","image:base":"81ab55b27588f4854075acf470a64bea02b2c1c3535f8067d6f37fe6e9a308b0","image:golang":"388bb46da994be8fbbc39ccf734f0e8a6e5d9131ebd877155d3c82dff4ac020b"}

FROM reg.c5h.io/golang@sha256:388bb46da994be8fbbc39ccf734f0e8a6e5d9131ebd877155d3c82dff4ac020b AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.2 --single-branch https://github.com/greboid/irc-github /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/github; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:81ab55b27588f4854075acf470a64bea02b2c1c3535f8067d6f37fe6e9a308b0

COPY --from=build /work/out /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
