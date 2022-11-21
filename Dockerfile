FROM golang:alpine as build

ARG COMPONENT

RUN apk --no-cache add tzdata

WORKDIR /app

ADD ./cmd/$COMPONENT/*.go ./cmd/$COMPONENT/
ADD ./cmd/$COMPONENT/config/*.yaml ./cmd/$COMPONENT/config/
ADD ./internal/ ./internal/

COPY go.mod go.sum ./
RUN go mod download && go mod verify

WORKDIR /app/cmd/$COMPONENT

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X app.buildVersion=0.0.1 -X 'app.buildDate=$(date +'%Y/%m/%d %H:%M:%S')' -X app.buildCommit=0.0.1 " -a -installsuffix cgo -o app .


FROM scratch as final

ARG COMPONENT

COPY --from=build /app/cmd/$COMPONENT/app /
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ADD ./cmd/$COMPONENT/config/*.yaml /config/
ADD ./cmd/$COMPONENT/config/*.pem /config/
ADD ./cmd/$COMPONENT/*.crt /
ADD ./cmd/$COMPONENT/*.key /

ENV TZ=Europe/Moscow

CMD ["/app"]
