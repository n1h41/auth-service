FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o main .
ENV PORT 8080
ENV DB_HOST rain.db.elephantsql.com
ENV DB_PORT 5432
ENV DB_USER acleugpi
ENV DB_PASS yoEgYBhX61OYLwH4Cg8f23XE_j4nX3_j
ENV DB_NAME acleugpi
ENV JWT_SECRET n2ks9x64
ENV SMPT_HOST smtp.mailgun.org
ENV SMPT_PORT 587
ENV SMPT_USER postmaster@sandbox89641ce6d0764fc8834b50755a77d322.mailgun.org
ENV SMPT_PASS ad05d48bd930141af22820d6c52a0c9f-77316142-f3c27377
EXPOSE 8080

CMD [ "./main" ]
