FROM alpine:3.21

EXPOSE 9101

ARG CI_COMMIT_TAG
ENV VERSION=$CI_COMMIT_TAG

RUN apk add gcompat

COPY prometheus-dispatcher-exporter_$VERSION /app/prometheus-dispatcher-exporter

CMD ["/app/prometheus-dispatcher-exporter"]
