# Generated from https://github.com/csmith/dockerfiles/blob/master/ergo-certwrapper/Dockerfile.gotpl
# BOM: {"github:csmith/certwrapper":"v4.2.0","github:ergochat/ergo":"v2.15.0","image:base":"4fea2559a3c315f01c6241c689da902344f4694dbdbdab4f5c0a1d24a8a78363","image:golang":"13c00b0ec8f2d15937134e1ed64661b5eb58ada268c1f90a33f039041e9df2e1"}

FROM reg.c5h.io/golang@sha256:13c00b0ec8f2d15937134e1ed64661b5eb58ada268c1f90a33f039041e9df2e1 AS build


RUN set -eux; \
    git clone --depth=1 -b v4.2.0 --single-branch https://github.com/csmith/certwrapper /go/src/github.com/csmith/certwrapper

WORKDIR /go/src/github.com/csmith/certwrapper


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath -tags lego_httpreq -o /out/certwrapper .; \
    go-licenses save ./... --save_path=/out/notices/certwrapper;



RUN set -eux; \
    git clone --depth=1 -b v2.15.0 --single-branch https://github.com/ergochat/ergo /go/src/github.com/ergochat/ergo

WORKDIR /go/src/github.com/ergochat/ergo


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o /out/ergo .; \
    go-licenses save ./... --save_path=/out/notices/ergo;


RUN set -eux; \
    mkdir -p /out/ircd /out/certs /out/ircd-bin; \
    chown -R 65532:65532 /out/ircd /out/certs /out/ircd-bin;



RUN set -eux; \
    # Copy over languages \
    cp -r /go/src/github.com/ergochat/ergo/languages /out/ircd-bin/languages;

FROM reg.c5h.io/base@sha256:4fea2559a3c315f01c6241c689da902344f4694dbdbdab4f5c0a1d24a8a78363

ENV CERTWRAPPER_CERTIFICATE_PATH=/certs/certificate.pem \
    CERTWRAPPER_DNS_PROVIDER=httpreq \
    CERTWRAPPER_ISSUER_PATH=/certs/issuer.perm \
    CERTWRAPPER_PRIVATE_KEY_PATH=/certs/privatekey.pem \
    CERTWRAPPER_USER_PATH=/certs/user.json

COPY --from=build /out/ /
VOLUME ["/ircd", "/certs"]

ENTRYPOINT ["/certwrapper", "/ergo", "run"]
CMD ["--conf", "/ircd/ircd.yaml"]
