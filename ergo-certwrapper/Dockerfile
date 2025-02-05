# Generated from https://github.com/csmith/dockerfiles/blob/master/ergo-certwrapper/Dockerfile.gotpl
# BOM: {"github:csmith/certwrapper":"v4.2.0","github:ergochat/ergo":"v2.15.0","image:base":"27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4","image:golang":"ada7e34fe68a90eb6495bb95c3e19297ee91630356b1d4c50b496db8c3f9ad26"}

FROM reg.c5h.io/golang@sha256:ada7e34fe68a90eb6495bb95c3e19297ee91630356b1d4c50b496db8c3f9ad26 AS build


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

FROM reg.c5h.io/base@sha256:27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4

ENV CERTWRAPPER_CERTIFICATE_PATH=/certs/certificate.pem \
    CERTWRAPPER_DNS_PROVIDER=httpreq \
    CERTWRAPPER_ISSUER_PATH=/certs/issuer.perm \
    CERTWRAPPER_PRIVATE_KEY_PATH=/certs/privatekey.pem \
    CERTWRAPPER_USER_PATH=/certs/user.json

COPY --from=build /out/ /
VOLUME ["/ircd", "/certs"]

ENTRYPOINT ["/certwrapper", "/ergo", "run"]
CMD ["--conf", "/ircd/ircd.yaml"]
