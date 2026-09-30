# Generated from https://github.com/csmith/dockerfiles/blob/master/hooktower/Dockerfile.gotpl
# BOM: {"github:https://github.com/adnanh/webhook":"2.8.3","github:https://github.com/containrrr/watchtower":"v1.7.1","image:base":"68804ff33ba310905be47653301a7636727a352c4b820d9e5e672df96577a3dc","image:golang":"034439cb9fb2e8d3d2f43c37f0755580b215e299bcb3545dd3fcbef624ebd653"}

FROM reg.c5h.io/golang@sha256:034439cb9fb2e8d3d2f43c37f0755580b215e299bcb3545dd3fcbef624ebd653 AS build


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


FROM reg.c5h.io/base@sha256:68804ff33ba310905be47653301a7636727a352c4b820d9e5e672df96577a3dc

COPY --from=build /go/src/github.com/containrrr/watchtower/watchtower /watchtower
COPY --from=build /go/src/github.com/adnanh/webhook/webhook /webhook
COPY --from=build /out/notices /notices

ENTRYPOINT ["/webhook"]
