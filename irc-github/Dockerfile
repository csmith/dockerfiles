# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.3","image:base":"df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a","image:golang":"fe297beab25063e7f09a6a45a799d4c877b9f393b6c67cca68469f9f3eb332a9"}

FROM reg.c5h.io/golang@sha256:fe297beab25063e7f09a6a45a799d4c877b9f393b6c67cca68469f9f3eb332a9 AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.3 --single-branch https://github.com/greboid/irc-github /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/github; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a

COPY --from=build /work/out /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
