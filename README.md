# GINCRUD
# Index
- [English](#en)
- [Português](#pt-br)

# *EN*

A CLI tool written in Go that automatically generates the basic boilerplate (controller, model, repository, and service) for a resource using Gin Gonic.

### Purpose

The goal of this project is to speed up the creation of RESTful resources in applications that use Gin by generating initial files with enough structure to start implementing routes, models, and persistence logic.

### Prerequisites

- Go installed (Go 1.16+ recommended)

### Installation

To compile the binary locally:
```bash
go build -o gincrud .
```

Or run it directly with:

```bash
go run main.go <args>
```
### Usage

Basic syntax:

```bash
gincrud [command] [resource-name] [output-directory]
``` 
Supported commands:

- create — creates the set of files for the specified resource.

- help — displays help information.

### Example:

```bash
# run without compiling
go run main.go create book ./output/
```
```bash
# run compiled binary
./gincrud create user ./src/
``` 
Notes about the output parameter:

### What is generated

When running `create` for the resource book, the following files will be created (in the directory `./output/book`):

- `book.controller.go` — route registration and handlers (empty functions)

- `book.model.go` — initial model definition (BookModel)

- `book.repository.go` — repository skeleton

- `book.service.go` — service skeleton

The generated package uses the resource name in lowercase as the package name (e.g., `package book`).

### Best Practices

- Review and adapt the generated files before using them in production.

- Add validation, error handling, and database integration as needed.

### Contribution

Pull requests are welcome. Open an issue to discuss larger changes.

### License

MIT — feel free to use and adapt this project.


# *PT-BR*


Ferramenta CLI escrita em Go que gera automaticamente o boilerplate básico (controller, model, repository e service) para um recurso utilizando Gin Gonic.

### Objetivo

O objetivo do projeto é acelerar a criação de recursos RESTful em aplicações que usam Gin, gerando arquivos iniciais com estrutura suficiente para começar a implementar as rotas, modelos e lógica de persistência.

### Pré-requisitos

- Go instalado (1.16+ recomendado)

### Instalação

Para compilar o binário localmente:

```bash
go build -o gincrud .
```

Ou executar diretamente com:

```bash
go run main.go <args>
```

### Uso

Sintaxe básica:

```bash
gincrud [command] [resource-name] [output-directory]
```

Comandos suportados:

- `create` — cria o conjunto de arquivos para o recurso informado.
- `help` — mostra a ajuda.

Exemplo:

```bash
# executar sem compilar
go run main.go create book ./output/

# executar binário compilado
./gincrud create user ./src/
```

Observações sobre o parâmetro de saída:

### O que é gerado

Ao executar `create` para o recurso `book` serão criados (no diretório `./output/book/`):

- `book.controller.go` — registro de rotas e handlers (funções vazias)
- `book.model.go` — definição inicial do model (`BookModel`)
- `book.repository.go` — esqueleto do repositório
- `book.service.go` — esqueleto do service

O pacote gerado usa o nome do recurso em minúsculas como package (ex.: `package book`).

### Boas práticas

- Revise e adapte os arquivos gerados antes de usá-los em produção.
- Adicione validação, tratamento de erros e integração com banco conforme necessário.

### Contribuição

Pull requests são bem-vindos. Abra uma issue para discutir alterações maiores.

### Licença

MIT — sinta-se à vontade para usar e adaptar o projeto.
