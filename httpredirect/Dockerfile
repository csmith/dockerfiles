# Generated from https://github.com/csmith/dockerfiles/blob/master/httpredirect/Dockerfile.gotpl
# BOM: {"github:greboid/httpredirect":"v1.1.0","image:base":"22cd0188413016b6119bf80c3e3be20bf03b54e9ce6b9ff1576ab07f64f6da66","image:golang":"e077549c644e3d9f04836265464d0b49f0eecef44b7a777274a66cda8830225b"}

FROM reg.c5h.io/golang@sha256:e077549c644e3d9f04836265464d0b49f0eecef44b7a777274a66cda8830225b AS build

RUN set -eux; \
    git clone --depth=1 -b v1.1.0 --single-branch https://github.com/greboid/httpredirect /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:22cd0188413016b6119bf80c3e3be20bf03b54e9ce6b9ff1576ab07f64f6da66

COPY --from=build /work/out /httpredirect
COPY --from=build /notices /notices
ENTRYPOINT ["/httpredirect"]


EXPOSE 8080
