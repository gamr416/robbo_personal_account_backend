FROM golang:latest

WORKDIR /robbo_student_personal_account

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./robbo_control_acces

RUN go build -o /robbo_student_personal_account

EXPOSE 8080

CMD [ "/app" ]