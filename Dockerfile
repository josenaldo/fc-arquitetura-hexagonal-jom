# Atualizei a imagem para a versão 1.22 do golang
FROM golang:1.22

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

# Instalação do cobra-cli, para criar aplicativos de linha de comando
RUN go get -u github.com/spf13/cobra@latest & go install github.com/spf13/cobra-cli@latest

# Instalação do mockgen, para gerar mocks de testes
# Mudei o mockgen da versão da google para a da uber, pois a da google foi descontinuada
# Como a versão da uber exige uma versão mais nova do Go, atualizei a imagem para a versão 1.22
RUN go get -u go.uber.org/mock/mockgen@v0.4.0 & go install go.uber.org/mock/mockgen@v0.4.0

# Instalação dpo sqlite3
RUN apt-get update && apt-get install sqlite3 -y

RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache
USER www-data

CMD ["tail", "-f", "/dev/null"]