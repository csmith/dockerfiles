# Generated from https://github.com/csmith/dockerfiles/blob/master/golang/Dockerfile.gotpl
# BOM: {"apk:bash":"5.1.4-r0","apk:binutils":"2.35.2-r2","apk:brotli-libs":"1.0.9-r5","apk:busybox":"1.33.1-r3","apk:ca-certificates":"20191127-r5","apk:curl":"7.79.1-r0","apk:expat":"2.4.1-r0","apk:gcc":"10.3.1_git20210424-r2","apk:git":"2.32.0-r0","apk:gmp":"6.2.1-r0","apk:go":"1.16.8-r0","apk:isl22":"0.22-r0","apk:libatomic":"10.3.1_git20210424-r2","apk:libcrypto1.1":"1.1.1l-r0","apk:libcurl":"7.79.1-r0","apk:libgcc":"10.3.1_git20210424-r2","apk:libgomp":"10.3.1_git20210424-r2","apk:libgphobos":"10.3.1_git20210424-r2","apk:libssl1.1":"1.1.1l-r0","apk:libstdc++":"10.3.1_git20210424-r2","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.2-r3","apk:musl-dev":"1.2.2-r3","apk:ncurses-libs":"6.2_p20210612-r0","apk:ncurses-terminfo-base":"6.2_p20210612-r0","apk:nghttp2-libs":"1.43.0-r0","apk:pcre2":"10.36-r0","apk:readline":"8.1.0-r0","apk:zlib":"1.2.11-r3","golang":"go1.17.1","image:alpine":"sha256:f32b56d2f268e9a4dce68b95adac9f81b8ed8f9bc7fd33ba0d804e7426f603fd"}

FROM reg.c5h.io/alpine@sha256:f32b56d2f268e9a4dce68b95adac9f81b8ed8f9bc7fd33ba0d804e7426f603fd

ARG ARCHIVE_URL="https:/golang.org/dl/go1.17.1.src.tar.gz"
ARG ARCHIVE_SUM="49dc08339770acd5613312db8c141eaf61779995577b89d93b541ef83067e5b1"

ENV GOPATH=/go \
    PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -eux; \
    # Dependencies and source for bootstrapping go
    apk add --no-cache --virtual .build-deps \
        bash=5.1.4-r0 \
        binutils=2.35.2-r2 \
        brotli-libs=1.0.9-r5 \
        busybox=1.33.1-r3 \
        ca-certificates=20191127-r5 \
        curl=7.79.1-r0 \
        gcc=10.3.1_git20210424-r2 \
        gmp=6.2.1-r0 \
        go=1.16.8-r0 \
        isl22=0.22-r0 \
        libatomic=10.3.1_git20210424-r2 \
        libcrypto1.1=1.1.1l-r0 \
        libcurl=7.79.1-r0 \
        libgcc=10.3.1_git20210424-r2 \
        libgomp=10.3.1_git20210424-r2 \
        libgphobos=10.3.1_git20210424-r2 \
        libssl1.1=1.1.1l-r0 \
        libstdc++=10.3.1_git20210424-r2 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.2-r3 \
        musl-dev=1.2.2-r3 \
        ncurses-libs=6.2_p20210612-r0 \
        ncurses-terminfo-base=6.2_p20210612-r0 \
        nghttp2-libs=1.43.0-r0 \
        readline=8.1.0-r0 \
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
        binutils=2.35.2-r2 \
        brotli-libs=1.0.9-r5 \
        busybox=1.33.1-r3 \
        ca-certificates=20191127-r5 \
        expat=2.4.1-r0 \
        gcc=10.3.1_git20210424-r2 \
        git=2.32.0-r0 \
        gmp=6.2.1-r0 \
        isl22=0.22-r0 \
        libatomic=10.3.1_git20210424-r2 \
        libcrypto1.1=1.1.1l-r0 \
        libcurl=7.79.1-r0 \
        libgcc=10.3.1_git20210424-r2 \
        libgomp=10.3.1_git20210424-r2 \
        libgphobos=10.3.1_git20210424-r2 \
        libssl1.1=1.1.1l-r0 \
        libstdc++=10.3.1_git20210424-r2 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.2-r3 \
        musl-dev=1.2.2-r3 \
        nghttp2-libs=1.43.0-r0 \
        pcre2=10.36-r0 \
        zlib=1.2.11-r3 \
        ; \
    # Set up the go path
    mkdir -p $GOPATH/src $GOPATH/bin; \
    chmod -R 777 $GOPATH;
