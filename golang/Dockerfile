# Generated from https://github.com/csmith/dockerfiles/blob/master/golang/Dockerfile.gotpl
# BOM: {"apk:bash":"5.2.15-r0","apk:binutils":"2.39-r2","apk:brotli-libs":"1.0.9-r9","apk:busybox":"1.35.0-r29","apk:busybox-binsh":"1.35.0-r29","apk:ca-certificates":"20220614-r4","apk:curl":"7.88.1-r1","apk:gcc":"12.2.1_git20220924-r4","apk:git":"2.38.4-r1","apk:gmp":"6.2.1-r2","apk:go":"1.19.7-r0","apk:isl24":"0.24-r1","apk:libatomic":"12.2.1_git20220924-r4","apk:libcrypto3":"3.0.8-r3","apk:libcurl":"7.88.1-r1","apk:libexpat":"2.5.0-r0","apk:libgcc":"12.2.1_git20220924-r4","apk:libgomp":"12.2.1_git20220924-r4","apk:libssl3":"3.0.8-r3","apk:libstdc++":"12.2.1_git20220924-r4","apk:mpc1":"1.2.1-r1","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.3-r4","apk:musl-dev":"1.2.3-r4","apk:ncurses-libs":"6.3_p20221119-r0","apk:ncurses-terminfo-base":"6.3_p20221119-r0","apk:nghttp2-libs":"1.51.0-r0","apk:pcre2":"10.42-r0","apk:readline":"8.2.0-r0","apk:zlib":"1.2.13-r0","github:google/go-licenses":"v1.6.0","golang":"go1.20.3","image:alpine":"e467edf34e56f1ea4884c2512b2250463f3379ffa4524d8f4e5e31e46c8b63b4"}

FROM reg.c5h.io/alpine@sha256:e467edf34e56f1ea4884c2512b2250463f3379ffa4524d8f4e5e31e46c8b63b4

ARG GOLICENSETAG="v1.6.0"
ARG ARCHIVE_URL="https:/golang.org/dl/go1.20.3.src.tar.gz"
ARG ARCHIVE_SUM="e447b498cde50215c4f7619e5124b0fc4e25fb5d16ea47271c47f278e7aa763a"

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -eux; \
    # Dependencies and source for bootstrapping go
    apk add --no-cache --virtual .build-deps \
        bash=5.2.15-r0 \
        binutils=2.39-r2 \
        brotli-libs=1.0.9-r9 \
        busybox=1.35.0-r29 \
        busybox-binsh=1.35.0-r29 \
        ca-certificates=20220614-r4 \
        curl=7.88.1-r1 \
        gcc=12.2.1_git20220924-r4 \
        gmp=6.2.1-r2 \
        go=1.19.7-r0 \
        isl24=0.24-r1 \
        libatomic=12.2.1_git20220924-r4 \
        libcrypto3=3.0.8-r3 \
        libcurl=7.88.1-r1 \
        libgcc=12.2.1_git20220924-r4 \
        libgomp=12.2.1_git20220924-r4 \
        libssl3=3.0.8-r3 \
        libstdc++=12.2.1_git20220924-r4 \
        mpc1=1.2.1-r1 \
        mpfr4=4.1.0-r0 \
        musl=1.2.3-r4 \
        musl-dev=1.2.3-r4 \
        ncurses-libs=6.3_p20221119-r0 \
        ncurses-terminfo-base=6.3_p20221119-r0 \
        nghttp2-libs=1.51.0-r0 \
        readline=8.2.0-r0 \
        zlib=1.2.13-r0 \
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
        binutils=2.39-r2 \
        brotli-libs=1.0.9-r9 \
        busybox=1.35.0-r29 \
        busybox-binsh=1.35.0-r29 \
        ca-certificates=20220614-r4 \
        gcc=12.2.1_git20220924-r4 \
        git=2.38.4-r1 \
        gmp=6.2.1-r2 \
        isl24=0.24-r1 \
        libatomic=12.2.1_git20220924-r4 \
        libcrypto3=3.0.8-r3 \
        libcurl=7.88.1-r1 \
        libexpat=2.5.0-r0 \
        libgcc=12.2.1_git20220924-r4 \
        libgomp=12.2.1_git20220924-r4 \
        libssl3=3.0.8-r3 \
        libstdc++=12.2.1_git20220924-r4 \
        mpc1=1.2.1-r1 \
        mpfr4=4.1.0-r0 \
        musl=1.2.3-r4 \
        musl-dev=1.2.3-r4 \
        nghttp2-libs=1.51.0-r0 \
        pcre2=10.42-r0 \
        zlib=1.2.13-r0 \
        ; \
    # Set up the go path
    mkdir -p $GOPATH/src $GOPATH/bin; \
    chmod -R 777 $GOPATH; \
    go install github.com/google/go-licenses@$GOLICENSETAG
