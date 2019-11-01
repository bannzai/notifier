FROM golang:1.13.3-stretch AS build
WORKDIR /go/src/repo
COPY . .
ENV GO111MODULE=on
RUN go build -o app

FROM gcr.io/distroless/base
COPY --from=build /go/src/repo/app /app
ENV PORT=${PORT}
ENTRYPOINT [ "/app" ]
CMD [ ".app" ]
