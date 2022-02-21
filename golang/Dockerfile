# Generated from https://github.com/csmith/dockerfiles/blob/master/golang/Dockerfile.gotpl
# BOM: {"apk:bash":"5.1.16-r0","apk:binutils":"2.37-r3","apk:brotli-libs":"1.0.9-r5","apk:busybox":"1.34.1-r4","apk:ca-certificates":"20211220-r0","apk:curl":"7.80.0-r0","apk:expat":"2.4.5-r0","apk:gcc":"10.3.1_git20211027-r0","apk:git":"2.34.1-r0","apk:gmp":"6.2.1-r1","apk:go":"1.17.4-r0","apk:isl22":"0.22-r0","apk:libatomic":"10.3.1_git20211027-r0","apk:libcrypto1.1":"1.1.1l-r8","apk:libcurl":"7.80.0-r0","apk:libgcc":"10.3.1_git20211027-r0","apk:libgomp":"10.3.1_git20211027-r0","apk:libgphobos":"10.3.1_git20211027-r0","apk:libssl1.1":"1.1.1l-r8","apk:libstdc++":"10.3.1_git20211027-r0","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.2-r7","apk:musl-dev":"1.2.2-r7","apk:ncurses-libs":"6.3_p20211120-r0","apk:ncurses-terminfo-base":"6.3_p20211120-r0","apk:nghttp2-libs":"1.46.0-r0","apk:pcre2":"10.39-r0","apk:readline":"8.1.1-r0","apk:zlib":"1.2.11-r3","golang":"go1.17.7","image:alpine":"ef95559f475cdf56e55db71c168f1616d497e2215bbd9f6bc63398036ae546a1"}

FROM reg.c5h.io/alpine@sha256:ef95559f475cdf56e55db71c168f1616d497e2215bbd9f6bc63398036ae546a1

ARG ARCHIVE_URL="https:/golang.org/dl/go1.17.7.src.tar.gz"
ARG ARCHIVE_SUM="c108cd33b73b1911a02b697741df3dea43e01a5c4e08e409e8b3a0e3745d2b4d"

ENV GOPATH=/go \
    PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -eux; \
    # Dependencies and source for bootstrapping go
    apk add --no-cache --virtual .build-deps \
        bash=5.1.16-r0 \
        binutils=2.37-r3 \
        brotli-libs=1.0.9-r5 \
        busybox=1.34.1-r4 \
        ca-certificates=20211220-r0 \
        curl=7.80.0-r0 \
        gcc=10.3.1_git20211027-r0 \
        gmp=6.2.1-r1 \
        go=1.17.4-r0 \
        isl22=0.22-r0 \
        libatomic=10.3.1_git20211027-r0 \
        libcrypto1.1=1.1.1l-r8 \
        libcurl=7.80.0-r0 \
        libgcc=10.3.1_git20211027-r0 \
        libgomp=10.3.1_git20211027-r0 \
        libgphobos=10.3.1_git20211027-r0 \
        libssl1.1=1.1.1l-r8 \
        libstdc++=10.3.1_git20211027-r0 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.2-r7 \
        musl-dev=1.2.2-r7 \
        ncurses-libs=6.3_p20211120-r0 \
        ncurses-terminfo-base=6.3_p20211120-r0 \
        nghttp2-libs=1.46.0-r0 \
        readline=8.1.1-r0 \
        zlib=1.2.11-r3 \
        ; \
    curl -sfLo go.tar.gz $ARCHIVE_URL; \
    echo "$ARCHIVE_SUM *go.tar.gz" | sha256sum -wc -; \
    tar -C /usr/local -xzf go.tar.gz; \
    # Bootstrap and install the standard library
    cd /usr/local/go/src && GOROOT_BOOTSTRAP="$(go env GOROOT)" GOHOSTOS="linux" GOHOSTARCH="amd64" ./make.bash; \
    go install std; \
    # Clean up a little
	apk del --no-network .build-deps; \
    rm -rf \
            /go.tar.gz \
            /root/.cache \
    		/usr/local/go/pkg/*/cmd \
    		/usr/local/go/pkg/bootstrap \
    		/usr/local/go/pkg/obj \
    		/usr/local/go/pkg/tool/*/api \
    		/usr/local/go/pkg/tool/*/go_bootstrap \
    		/usr/local/go/src/cmd/dist/dist; \
    # Dependencies commonly needed for building go apps
    apk add --no-cache \
        binutils=2.37-r3 \
        brotli-libs=1.0.9-r5 \
        busybox=1.34.1-r4 \
        ca-certificates=20211220-r0 \
        expat=2.4.5-r0 \
        gcc=10.3.1_git20211027-r0 \
        git=2.34.1-r0 \
        gmp=6.2.1-r1 \
        isl22=0.22-r0 \
        libatomic=10.3.1_git20211027-r0 \
        libcrypto1.1=1.1.1l-r8 \
        libcurl=7.80.0-r0 \
        libgcc=10.3.1_git20211027-r0 \
        libgomp=10.3.1_git20211027-r0 \
        libgphobos=10.3.1_git20211027-r0 \
        libssl1.1=1.1.1l-r8 \
        libstdc++=10.3.1_git20211027-r0 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.2-r7 \
        musl-dev=1.2.2-r7 \
        nghttp2-libs=1.46.0-r0 \
        pcre2=10.39-r0 \
        zlib=1.2.11-r3 \
        ; \
    # Set up the go path
    mkdir -p $GOPATH/src $GOPATH/bin; \
    chmod -R 777 $GOPATH;
