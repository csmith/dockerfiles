FROM reg.c5h.io/golang@sha256:69c5302b51efc9e74ffc330a53d151fb95a65dd76dbb4d6f888579b91fb84c28 AS build

ARG TAG="v1.7.3"

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
    # Allow the vault process to prevent its memory regions being swapped out to disk
    /usr/sbin/setcap CAP_IPC_LOCK=+ep /vault;

FROM reg.c5h.io/alpine@sha256:1775bebec23e1f3ce486989bfc9ff3c4e951690df84aa9f926497d82f2ffca9d

COPY --from=build /vault /bin/
COPY --from=build /out/ /vault/
EXPOSE 8080
VOLUME /vault/logs /vault/file
ENTRYPOINT ["/bin/vault"]
ENV VAULT_ADDR http://127.0.0.1:8080
CMD ["server", "-config=/vault/config/"]
