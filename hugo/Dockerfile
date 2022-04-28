# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.37-r3","apk:g++":"10.3.1_git20211027-r0","apk:gcc":"10.3.1_git20211027-r0","apk:gmp":"6.2.1-r1","apk:isl22":"0.22-r0","apk:libatomic":"10.3.1_git20211027-r0","apk:libc-dev":"0.7.2-r3","apk:libgcc":"10.3.1_git20211027-r0","apk:libgomp":"10.3.1_git20211027-r0","apk:libgphobos":"10.3.1_git20211027-r0","apk:libstdc++":"10.3.1_git20211027-r0","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.2-r7","apk:musl-dev":"1.2.2-r7","apk:zlib":"1.2.12-r1","github:gohugoio/hugo":"v0.98.0","github:magefile/mage":"v1.13.0","image:base-glibc":"902eeaa2ba8128b63e4247b947911c41df14620da6d44fef797179ff80827347","image:golang":"9eb0fe2aa6f733756bd07464419bfa32e312bcf03e4dd4a3ec368b18cbd5c73a"}

FROM reg.c5h.io/golang@sha256:9eb0fe2aa6f733756bd07464419bfa32e312bcf03e4dd4a3ec368b18cbd5c73a AS build

ARG TAG="v0.98.0"
ARG MAGE="v1.13.0"

RUN set -eux; \
    apk add --no-cache \
        binutils=2.37-r3 \
        g++=10.3.1_git20211027-r0 \
        gcc=10.3.1_git20211027-r0 \
        gmp=6.2.1-r1 \
        isl22=0.22-r0 \
        libatomic=10.3.1_git20211027-r0 \
        libc-dev=0.7.2-r3 \
        libgcc=10.3.1_git20211027-r0 \
        libgomp=10.3.1_git20211027-r0 \
        libgphobos=10.3.1_git20211027-r0 \
        libstdc++=10.3.1_git20211027-r0 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.2-r7 \
        musl-dev=1.2.2-r7 \
        zlib=1.2.12-r1 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/gohugoio/hugo /tmp/work; \
    cd /tmp/work; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE hugo; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE install; \
    mkdir -p /tmp/site;

FROM reg.c5h.io/base-glibc@sha256:902eeaa2ba8128b63e4247b947911c41df14620da6d44fef797179ff80827347

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
