# Generated from https://github.com/csmith/dockerfiles/blob/master/httpredirect/Dockerfile.gotpl
# BOM: {"github:https://github.com/greboid/httpredirect":"v1.1.0","image:base":"996f481af28d13a3c5069644629fbb02aace36a81a912b688db0771db8669424","image:golang":"3204d08f9b5525b3641537d5e1159e8e751622811967f646de0e108b8148fa59"}

FROM reg.c5h.io/golang@sha256:3204d08f9b5525b3641537d5e1159e8e751622811967f646de0e108b8148fa59 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.1.0 --single-branch https://github.com/greboid/httpredirect /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:996f481af28d13a3c5069644629fbb02aace36a81a912b688db0771db8669424

COPY --from=build /work/out /httpredirect
COPY --from=build /notices /notices
ENTRYPOINT ["/httpredirect"]


EXPOSE 8080
