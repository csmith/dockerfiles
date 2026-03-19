# Generated from https://github.com/csmith/dockerfiles/blob/master/hooktower/Dockerfile.gotpl
# BOM: {"github:adnanh/webhook":"2.8.3","github:containrrr/watchtower":"v1.7.1","image:base":"84a319cda607326ba5aff905ae01191c4a3c793dbad0e058d9a816e3fb29036f","image:golang":"e76894c73b00fe5e379c1a2c5de8a721160d08ac76476b567e0ef9dccdc7d494"}

FROM reg.c5h.io/golang@sha256:e76894c73b00fe5e379c1a2c5de8a721160d08ac76476b567e0ef9dccdc7d494 AS build


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


FROM reg.c5h.io/base@sha256:84a319cda607326ba5aff905ae01191c4a3c793dbad0e058d9a816e3fb29036f

COPY --from=build /go/src/github.com/containrrr/watchtower/watchtower /watchtower
COPY --from=build /go/src/github.com/adnanh/webhook/webhook /webhook
COPY --from=build /out/notices /notices

ENTRYPOINT ["/webhook"]
