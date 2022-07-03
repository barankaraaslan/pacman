FROM golang

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 go build -o pacman .

FROM scratch
COPY --from=0 /app/pacman /pacman
ENTRYPOINT [ "/pacman" ]