FROM golang:1.19-alpine3.15
RUN mkdir /app
ADD . /app
WORKDIR /app
ARG EnvironmentVariable
RUN go build -o main .
EXPOSE 8080

CMD /app/main