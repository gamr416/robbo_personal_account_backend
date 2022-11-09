FROM golang:latest

WORKDIR /robbo_student_personal_account

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o robbo_student_personal_account

EXPOSE 8080

CMD [ "/robbo_student_personal_account" ]