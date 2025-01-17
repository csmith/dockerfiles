# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.2","image:base":"27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4","image:golang":"369aafeb4d799e6a42ca613e11bada8c88fe8fdc75226e5f95817f5a240591a9"}

FROM reg.c5h.io/golang@sha256:369aafeb4d799e6a42ca613e11bada8c88fe8fdc75226e5f95817f5a240591a9 AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.2 --single-branch https://github.com/greboid/irc-github /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/github; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4

COPY --from=build /work/out /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
