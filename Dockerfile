FROM golang:1.13.3-stretch AS build
WORKDIR /go/src/repo
COPY . .
ENV GO111MODULE=on
RUN go build -o app

FROM gcr.io/distroless/base
COPY --from=build /go/src/repo/app /app
COPY --from=build /go/src/repo/test.yml /test.yml
ENV PORT=${PORT}
ENV YAML_FILE_PATH=${YAML_FILE_PATH}
ENV NOTIFIER_SLACK_TOKEN=${NOTIFIER_SLACK_TOKEN}
ENTRYPOINT [ "/app" ]
CMD [ ".app" ]
