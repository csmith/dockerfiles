FROM reg.c5h.io/alpine@sha256:1775bebec23e1f3ce486989bfc9ff3c4e951690df84aa9f926497d82f2ffca9d

ARG ARCHIVE_URL="https:/golang.org/dl/go1.16.6.src.tar.gz"
ARG ARCHIVE_SUM="a3a5d4bc401b51db065e4f93b523347a4d343ae0c0b08a65c3423b05a138037d"

ENV GOPATH=/go \
    PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -eux; \
    # Dependencies and source for bootstrapping go
    apk add --no-cache --virtual .build-deps curl bash go; \
    curl -sfLo go.tar.gz $ARCHIVE_URL; \
    echo "$ARCHIVE_SUM *go.tar.gz" | sha256sum -wc -; \
    tar -C /usr/local -xzf go.tar.gz; \
    # Bootstrap and install the standard library
    cd /usr/local/go/src && GOROOT_BOOTSTRAP="$(go env GOROOT)" GOHOSTOS="linux" GOHOSTARCH="amd64" ./make.bash; \
    go install std; \
    # Clean up a little
	apk del --no-network .build-deps; \
    rm -rf \
            go.tar.gz \
    		/usr/local/go/pkg/*/cmd \
    		/usr/local/go/pkg/bootstrap \
    		/usr/local/go/pkg/obj \
    		/usr/local/go/pkg/tool/*/api \
    		/usr/local/go/pkg/tool/*/go_bootstrap \
    		/usr/local/go/src/cmd/dist/dist; \
    # Dependencies needed for cgo
    apk add gcc musl-dev; \
    # Set up the go path
    mkdir -p $GOPATH/src $GOPATH/bin; \
    chmod -R 777 $GOPATH;
