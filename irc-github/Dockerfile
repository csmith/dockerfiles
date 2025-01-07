# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.2","image:base":"a59bc827c80c4dbec2473c6b1a6f72b5b015d44e37981b112b9a60cb3c5180af","image:golang":"69bf99d272afa648a5758d3f6e435854fa6b6962df9418495dc5c2291e38f0df"}

FROM reg.c5h.io/golang@sha256:69bf99d272afa648a5758d3f6e435854fa6b6962df9418495dc5c2291e38f0df AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.2 --single-branch https://github.com/greboid/irc-github /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/github; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:a59bc827c80c4dbec2473c6b1a6f72b5b015d44e37981b112b9a60cb3c5180af

COPY --from=build /work/out /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
