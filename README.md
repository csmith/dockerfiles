# Dockerfiles from the ground up

## What? Why?

This is a collection of Dockerfiles for various software projects I want to run, built from the ground up.

Most projects have either official docker images or third-party contributions, but they're always a bit
hit-or-miss on how they work, what base images are used, etc. Third-party images often lag behind releases
or just die off without warning, too. Building everything out from scratch ensures the images are standard,
and can follow upstream updates as quickly or slowly as required.

I'm intending to use these for production services, but won't vouch for their stability or usability for
anyone else's purposes. Feel free to use them, and report any issues you do find, but at your own risk!

## Structure

Each image has its own folder, containing:

* `Dockerfile` - the actual dockerfile the latest version should be built from
* `Dockerfile.gotpl` - a template file for generating the Dockerfile

The template files using go's [text/template syntax](https://pkg.go.dev/text/template), with functions
defined to retrieve the latest images, releases, packages, etc.

To update all the Dockerfiles in the repo, run the `cmd` package (`go run ./cmd`). This will commit
each Dockerfile individually and build and push any changed images by default; for local development
you can pass `--commit=false` and `--build=false`.

## Images

All images are available at `reg.c5h.io/<name>`. Only the latest tag is built.

| Name               | Upstream                                              | Reproducible? | Non-root? | Minimal? |
|--------------------|-------------------------------------------------------|:-------------:|:---------:|:--------:|
| alpine             | https://alpinelinux.org/                              |       ✅      |    N/A    |    ✅    |
| arch               | https://archlinux.org/                                |       ✅      |    N/A    |    ❌    |
| base               | N/A                                                   |       ✅      |    ✅     |    ✅    |
| distribution       | https://github.com/distribution/distribution          |       ✅      |    ✅     |    ✅    |
| irc-bot            | https://github.com/greboid/irc-bot                    |       ✅      |    ✅     |    ✅    |
| ↳ irc-distribution | https://github.com/csmith/irc-distribution            |       ✅      |    ✅     |    ✅    |
| ↳ irc-github       | https://github.com/greboid/irc-github                 |       ✅      |    ✅     |    ✅    |
| ↳ irc-goplum       | https://github.com/greboid/irc-goplum                 |       ✅      |    ✅     |    ✅    |
| ↳ irc-news         | https://github.com/csmith/ircplugins                  |       ✅      |    ✅     |    ✅    |
| golang             | https://golang.org/                                   |       ✅      |    N/A    |    ✅    |
| linx-server        | https://github.com/csmith/linx-server                 |       ✅      |    ✅     |    ✅    |
| postgres-13        | https://www.postgresql.org/                           |       ❌      |    ❌     |    ❌    |
| vault              | https://github.com/hashicorp/vault                    |       ✅      |    ✅     |    ✅    |

Meaning of the status columns:

**Reproducible** - if the same Dockerfile is rebuilt at any time on any machine it will produce the same
image. This is nice to have, but it is quite challenging and makes little difference in day-to-day
operations. (It also requires the use of a tool like `buildah` that can set layer timestamps; `docker build`
cannot make reproducible images.)

**Non-root** - the entrypoint for the image is invoked as a non-root user. Base images are marked as N/A.
Other images probably drop root later either via a script or as part of the process itself, but it's
preferable for it to happen in the image.

**Minimal** - the image contains only the bare essentials required. No leftovers, nothing irrelevant,
no bloat. For base images this definition is a bit hazy as they'll contain things that might be used
in downstream images. For applications this generally means they're statically compiled and run in the
"base" image.
