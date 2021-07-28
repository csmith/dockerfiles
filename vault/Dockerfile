FROM reg.c5h.io/golang@sha256:dd7e126ba854294376c61b442d961839eee6e72e5ad36211155a90c70dc3a1b9 AS build

ARG TAG="v1.8.0"

RUN set -eux; \
    # Dependencies
    apk add --no-cache git libcap; \
    # Check out and build
    git clone --depth=1 -b $TAG --single-branch https://github.com/hashicorp/vault /src; \
    cd /src; \
    go build -tags release -ldflags "-X github.com/hashicorp/vault/sdk/version.GitCommit=$(git rev-parse HEAD)" -o /vault  github.com/hashicorp/vault; \
    # Make some directories with the correct ownership to copy into the final image, and a listener config snippet
    mkdir -p /out/logs /out/file /out/config; \
    echo -e 'listener "tcp" {\n  address = "0.0.0.0:8080"\n  tls_disable = 1\n}' > /out/config/listener.hcl; \
    chown -R 65532:65532 /out; \
    # Clobber all the timestamps to make the build more reproducible
    find /out -exec touch --date=@0 {} \;; \
    touch --date=@0 /vault;

FROM reg.c5h.io/alpine@sha256:1a78f21d1807da0423a4b835cf05bd85f665a875b5f5030750f25fecd15e41fc AS base
# This is basically the same thing as reg.c5h.io/base, but we add libcap and don't drop root until after we've
# had a chance to call setcap on the binary.

RUN set -eux; \
    # Sort out our dependencies (and sources)
    apk add --no-cache ca-certificates musl tzdata rsync libcap; \
    # Set up directories and password file
    mkdir -p /rootfs/etc/ssl/certs /rootfs/tmp /rootfs/home/nonroot /rootfs/root; \
    echo "root:x:0:0:root:/root:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nobody:x:65534:65534:nobody:/nonexistent:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nonroot:x:65532:65532:nonroot:/home/nonroot:/sbin/nologin" >> /rootfs/etc/passwd; \
    chown 65532:65532 /rootfs/home/nonroot; \
    chmod 01777 /tmp; \
    # Copy package contents
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    apk info -qL libcap | rsync -aq --files-from=- / /rootfs/; \
    # Generate thw CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Clobber all the timestamps to make the build more reproducible
    find /rootfs -exec touch --date=@0 {} \;;

FROM scratch

COPY --from=base /rootfs/ /
COPY --from=build /vault /bin/
COPY --from=build /out/ /vault/

RUN ["/usr/sbin/setcap","CAP_IPC_LOCK=+ep","/bin/vault"]

ENV VAULT_ADDR=http://127.0.0.1:8080 \
    SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt

USER nonroot
EXPOSE 8080
VOLUME /vault/logs /vault/file
ENTRYPOINT ["/bin/vault"]

CMD ["server", "-config=/vault/config/"]
