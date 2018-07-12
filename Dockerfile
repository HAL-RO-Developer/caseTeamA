FROM makki0205/deploy

WORKDIR /go/src/github.com/HAL-RO-Developer/caseTeamA


ADD ./ ./

RUN npm i
RUN npm run build

RUN go get github.com/satori/go.uuid
RUN cp config.yml.template config.yml

RUN ls -la

EXPOSE 8000

ENTRYPOINT ["go","run","main.go"]

