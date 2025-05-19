FROM golang:1.24-alpine AS build

WORKDIR /src
COPY . .
RUN go build -o /app .

FROM scratch

COPY --from=build /app /app

CMD [ "/app" ]
