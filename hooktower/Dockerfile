# Generated from https://github.com/csmith/dockerfiles/blob/master/hooktower/Dockerfile.gotpl
# BOM: {"github:adnanh/webhook":"2.8.3","github:containrrr/watchtower":"v1.7.1","image:base":"00146a62a048620af617bc351091bce6105ca77ea883f65bf45aa22aecf62958","image:golang":"66b2fc993c20f1598d8f877bc6228c2f814d0d3c1055f744c1b1f3a3bc355023"}

FROM reg.c5h.io/golang@sha256:66b2fc993c20f1598d8f877bc6228c2f814d0d3c1055f744c1b1f3a3bc355023 AS build


RUN set -eux; \
    git clone --depth=1 -b v1.7.1 --single-branch https://github.com/containrrr/watchtower /go/src/github.com/containrrr/watchtower

WORKDIR /go/src/github.com/containrrr/watchtower


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o watchtower .; \
    go-licenses save ./... --save_path=/out/notices/watchtower;



RUN set -eux; \
    git clone --depth=1 -b 2.8.3 --single-branch https://github.com/adnanh/webhook /go/src/github.com/adnanh/webhook

WORKDIR /go/src/github.com/adnanh/webhook


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o webhook .; \
    go-licenses save ./... --save_path=/out/notices/webhook;


FROM reg.c5h.io/base@sha256:00146a62a048620af617bc351091bce6105ca77ea883f65bf45aa22aecf62958

COPY --from=build /go/src/github.com/containrrr/watchtower/watchtower /watchtower
COPY --from=build /go/src/github.com/adnanh/webhook/webhook /webhook
COPY --from=build /out/notices /notices

ENTRYPOINT ["/webhook"]
