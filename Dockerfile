# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-taggit"
LABEL REPO="https://github.com/mrlyc/taggit"

ENV PROJPATH=/go/src/github.com/mrlyc/taggit

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/mrlyc/taggit
WORKDIR /go/src/github.com/mrlyc/taggit

RUN make build-alpine

# Final Stage
FROM mrlyc/taggit:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/mrlyc/taggit"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/taggit/bin

WORKDIR /opt/taggit/bin

COPY --from=build-stage /go/src/github.com/mrlyc/taggit/bin/taggit /opt/taggit/bin/
RUN chmod +x /opt/taggit/bin/taggit

# Create appuser
RUN adduser -D -g '' taggit
USER taggit

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/taggit/bin/taggit"]
