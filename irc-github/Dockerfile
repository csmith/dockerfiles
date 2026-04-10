# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.3","image:base":"46cb85d037365ac7498d4cefa5882d5a6242d9907e66abcde3a2c2fc62a48987","image:golang":"d2c7d4b907467670c9dc335754a890a061167ede406e99020e7d7e1a43eada6a"}

FROM reg.c5h.io/golang@sha256:d2c7d4b907467670c9dc335754a890a061167ede406e99020e7d7e1a43eada6a AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.3 --single-branch https://github.com/greboid/irc-github /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/github; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:46cb85d037365ac7498d4cefa5882d5a6242d9907e66abcde3a2c2fc62a48987

COPY --from=build /work/out /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
