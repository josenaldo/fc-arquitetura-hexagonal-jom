<h1 align="center">Welcome to fc-arquitetura-hexagonal-jom 👋</h1>
<p>
  <a href="LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://twitter.com/josenaldomatos" target="_blank">
    <img alt="Twitter: josenaldomatos" src="https://img.shields.io/twitter/follow/josenaldomatos.svg?style=social" />
  </a>
</p>

> Hexagonal Architecture course project from Full Cycle.

### 🏠 [Homepage](https://plataforma.fullcycle.com.br/courses/3b8c4f2c-aff9-4399-a72a-ad879e5689a2/315/168/123/conteudos?capitulo=123&conteudo=6963)

## Install

```sh
docekr compose up -d --build
```

## Usage

```sh
docker exec -it appproduct bash
```

## 📚 References

- [Arquitetura Hexagonal](https://blog.octo.com/en/hexagonal-architecture-three-principles-and-an-implementation-example/)
- [Docker](https://www.docker.com/)
- [Go Lang](https://golang.org/)
- [Google UUID](https://github.com/google/uuid)
- [Go Validator](https://github.com/asaskevich/govalidator)
- [Testify](https://github.com/stretchr/testify)
- [Mockgen](https://github.com/uber-go/mock)

## Comands used

Command used to generate the mocks

```sh
mockgen -destination=application/mocks/application.go -source=application/product.go application
```

```sh
## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2024 [Josenaldo de Oliveira Matos Filho](https://github.com/josenaldo).

This project is [MIT](LICENSE) licensed.

## Author

👤 **Josenaldo de Oliveira Matos Filho**

- Website: <https://josenaldo.github.io>
- Twitter: [@josenaldomatos](https://twitter.com/josenaldomatos)
- Github: [@josenaldo](https://github.com/josenaldo)
- LinkedIn: [@josenaldo](https://linkedin.com/in/josenaldo)
-
***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
