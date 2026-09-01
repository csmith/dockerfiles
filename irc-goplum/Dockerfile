# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:https://github.com/greboid/irc-goplum":"v4.0.2","image:base":"996f481af28d13a3c5069644629fbb02aace36a81a912b688db0771db8669424","image:golang":"f2ab88d9eba5aec87eaec29d063016f27ae3e19511d040e849f7b5033b956f8a"}

FROM reg.c5h.io/golang@sha256:f2ab88d9eba5aec87eaec29d063016f27ae3e19511d040e849f7b5033b956f8a AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.2 --single-branch https://github.com/greboid/irc-goplum /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/goplum; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:996f481af28d13a3c5069644629fbb02aace36a81a912b688db0771db8669424

COPY --from=build /work/out /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
