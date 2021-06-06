FROM golang:1.15 as build
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o courses-api cmd/main.go
 
FROM ubuntu
COPY --from=build /app/courses-api /courses-api
EXPOSE 9500
CMD ["./courses-api"]
