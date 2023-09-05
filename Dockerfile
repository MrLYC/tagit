# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-tagit"
LABEL REPO="https://github.com/mrlyc/tagit"

ENV PROJPATH=/go/src/github.com/mrlyc/tagit

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/mrlyc/tagit
WORKDIR /go/src/github.com/mrlyc/tagit

RUN make build-alpine

# Final Stage
FROM mrlyc/tagit:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/mrlyc/tagit"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/tagit/bin

WORKDIR /opt/tagit/bin

COPY --from=build-stage /go/src/github.com/mrlyc/tagit/bin/tagit /opt/tagit/bin/
RUN chmod +x /opt/tagit/bin/tagit

# Create appuser
RUN adduser -D -g '' tagit
USER tagit

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/tagit/bin/tagit"]
