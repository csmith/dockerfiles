# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.3","image:base":"932d79427476d41a4cc015dae576efd446b46fe22463ecf15a62f8a5b54b9929","image:golang":"bb6a061c510ddf8aaf4d480c1eabe858d1a996bf52ba05108d4da3da16a91d68"}

FROM reg.c5h.io/golang@sha256:bb6a061c510ddf8aaf4d480c1eabe858d1a996bf52ba05108d4da3da16a91d68 AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.3 --single-branch https://github.com/greboid/irc-github /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/github; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:932d79427476d41a4cc015dae576efd446b46fe22463ecf15a62f8a5b54b9929

COPY --from=build /work/out /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
