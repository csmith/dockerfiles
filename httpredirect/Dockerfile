# Generated from https://github.com/csmith/dockerfiles/blob/master/httpredirect/Dockerfile.gotpl
# BOM: {"github:greboid/httpredirect":"v1.0.3","image:base":"27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4","image:golang":"ee2aa9d15150ad6beb671f0febc38ef258eac10e247c32db8f9c12ea60eabf5e"}

FROM reg.c5h.io/golang@sha256:ee2aa9d15150ad6beb671f0febc38ef258eac10e247c32db8f9c12ea60eabf5e AS build

RUN set -eux; \
    git clone --depth=1 -b v1.0.3 --single-branch https://github.com/greboid/httpredirect /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4

COPY --from=build /work/out /httpredirect
COPY --from=build /notices /notices
ENTRYPOINT ["/httpredirect"]


EXPOSE 8080
