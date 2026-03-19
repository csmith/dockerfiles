# Generated from https://github.com/csmith/dockerfiles/blob/master/soju-certwrapper/Dockerfile.gotpl
# BOM: {"github:csmith/certwrapper":"v4.3.0","github:emersion/soju":"v0.10.1","image:base":"84a319cda607326ba5aff905ae01191c4a3c793dbad0e058d9a816e3fb29036f","image:golang":"e76894c73b00fe5e379c1a2c5de8a721160d08ac76476b567e0ef9dccdc7d494"}

FROM reg.c5h.io/golang@sha256:e76894c73b00fe5e379c1a2c5de8a721160d08ac76476b567e0ef9dccdc7d494 AS build

RUN set -eux; \
    git clone --depth=1 -b v4.3.0 --single-branch https://github.com/csmith/certwrapper /go/src/github.com/csmith/certwrapper

WORKDIR /go/src/github.com/csmith/certwrapper


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath -tags lego_httpreq -o /out/certwrapper .; \
    go-licenses save ./... --save_path=/out/notices/certwrapper;


RUN set -eux; \
    git clone --depth=1 -b v0.10.1 --single-branch https://github.com/emersion/soju /go/src/github.com/emersion/soju

WORKDIR /go/src/github.com/emersion/soju


RUN set -eux; \
    GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -tags=moderncsqlite -o /out/soju ./cmd/soju; \
    GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -tags=moderncsqlite -o /out/sojuctl ./cmd/sojuctl; \
    # TODO: find a license tool that doesn't hardcode Google's opinion on "forbidden" licenses :/
    # go-licenses save ./... --save_path=/notices; \
    # Create folders for volumes
    install -o 65532 -g 65532 -d /out/data /out/certs;

FROM reg.c5h.io/base@sha256:84a319cda607326ba5aff905ae01191c4a3c793dbad0e058d9a816e3fb29036f

ENV CERTWRAPPER_CERTIFICATE_PATH=/certs/certificate.pem \
    CERTWRAPPER_DNS_PROVIDER=httpreq \
    CERTWRAPPER_ISSUER_PATH=/certs/issuer.perm \
    CERTWRAPPER_PRIVATE_KEY_PATH=/certs/privatekey.pem \
    CERTWRAPPER_USER_PATH=/certs/user.json

COPY --from=build /out/ /
VOLUME ["/data", "/certs"]

ENTRYPOINT ["/certwrapper"]
CMD ["/soju"]
