# Generated from https://github.com/csmith/dockerfiles/blob/master/httpredirect/Dockerfile.gotpl
# BOM: {"github:greboid/httpredirect":"v1.0.3","image:base":"df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a","image:golang":"18a4339b2b4f079668919a005c982400f8cf8ce18f4752336395732498bddeef"}

FROM reg.c5h.io/golang@sha256:18a4339b2b4f079668919a005c982400f8cf8ce18f4752336395732498bddeef AS build

RUN set -eux; \
    git clone --depth=1 -b v1.0.3 --single-branch https://github.com/greboid/httpredirect /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a

COPY --from=build /work/out /httpredirect
COPY --from=build /notices /notices
ENTRYPOINT ["/httpredirect"]


EXPOSE 8080
