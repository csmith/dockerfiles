# Generated from https://github.com/csmith/dockerfiles/blob/master/ergo-certwrapper/Dockerfile.gotpl
# BOM: {"github:csmith/certwrapper":"v4.3.0","github:ergochat/ergo":"v2.18.0","image:base":"66af67568b123f6ec0ff7360f9f2f08e274753f885288e5732d6e0c087c1c1dd","image:golang":"79c4a958212bde36a835f9b7c9dd2f721090e1a214699a11cfddb72858bfac6d"}

FROM reg.c5h.io/golang@sha256:79c4a958212bde36a835f9b7c9dd2f721090e1a214699a11cfddb72858bfac6d AS build


RUN set -eux; \
    git clone --depth=1 -b v4.3.0 --single-branch https://github.com/csmith/certwrapper /go/src/github.com/csmith/certwrapper

WORKDIR /go/src/github.com/csmith/certwrapper


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath -tags lego_httpreq -o /out/certwrapper .; \
    go-licenses save ./... --save_path=/out/notices/certwrapper;



RUN set -eux; \
    git clone --depth=1 -b v2.18.0 --single-branch https://github.com/ergochat/ergo /go/src/github.com/ergochat/ergo

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

FROM reg.c5h.io/base@sha256:66af67568b123f6ec0ff7360f9f2f08e274753f885288e5732d6e0c087c1c1dd

ENV CERTWRAPPER_CERTIFICATE_PATH=/certs/certificate.pem \
    CERTWRAPPER_DNS_PROVIDER=httpreq \
    CERTWRAPPER_ISSUER_PATH=/certs/issuer.perm \
    CERTWRAPPER_PRIVATE_KEY_PATH=/certs/privatekey.pem \
    CERTWRAPPER_USER_PATH=/certs/user.json

COPY --from=build /out/ /
VOLUME ["/ircd", "/certs"]

ENTRYPOINT ["/certwrapper", "/ergo", "run"]
CMD ["--conf", "/ircd/ircd.yaml"]
