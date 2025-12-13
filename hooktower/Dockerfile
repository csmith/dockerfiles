# Generated from https://github.com/csmith/dockerfiles/blob/master/hooktower/Dockerfile.gotpl
# BOM: {"github:adnanh/webhook":"2.8.2","github:containrrr/watchtower":"v1.7.1","image:base":"df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a","image:golang":"18a4339b2b4f079668919a005c982400f8cf8ce18f4752336395732498bddeef"}

FROM reg.c5h.io/golang@sha256:18a4339b2b4f079668919a005c982400f8cf8ce18f4752336395732498bddeef AS build


RUN set -eux; \
    git clone --depth=1 -b v1.7.1 --single-branch https://github.com/containrrr/watchtower /go/src/github.com/containrrr/watchtower

WORKDIR /go/src/github.com/containrrr/watchtower


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o watchtower .; \
    go-licenses save ./... --save_path=/out/notices/watchtower;



RUN set -eux; \
    git clone --depth=1 -b 2.8.2 --single-branch https://github.com/adnanh/webhook /go/src/github.com/adnanh/webhook

WORKDIR /go/src/github.com/adnanh/webhook


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o webhook .; \
    go-licenses save ./... --save_path=/out/notices/webhook;


FROM reg.c5h.io/base@sha256:df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a

COPY --from=build /go/src/github.com/containrrr/watchtower/watchtower /watchtower
COPY --from=build /go/src/github.com/adnanh/webhook/webhook /webhook
COPY --from=build /out/notices /notices

ENTRYPOINT ["/webhook"]
