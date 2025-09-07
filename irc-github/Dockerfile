# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.2","image:base":"81ab55b27588f4854075acf470a64bea02b2c1c3535f8067d6f37fe6e9a308b0","image:golang":"c053db948fcfc72f52a70d832af007e0b1e877937d4d8f2876428e97931bb0e6"}

FROM reg.c5h.io/golang@sha256:c053db948fcfc72f52a70d832af007e0b1e877937d4d8f2876428e97931bb0e6 AS build

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
