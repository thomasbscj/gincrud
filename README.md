# gincrud

Ferramenta CLI escrita em Go que gera automaticamente o boilerplate básico (controller, model, repository e service) para um recurso utilizando Gin Gonic.

## Objetivo

O objetivo do projeto é acelerar a criação de recursos RESTful em aplicações que usam Gin, gerando arquivos iniciais com estrutura suficiente para começar a implementar as rotas, modelos e lógica de persistência.

## Pré-requisitos

- Go instalado (1.16+ recomendado)

## Instalação

Para compilar o binário localmente:

```bash
go build -o gincrud .
```

Ou executar diretamente com:

```bash
go run main.go <args>
```

## Uso

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
- Recomenda-se incluir a barra final no caminho de saída (`./output/`) para que o diretório do recurso seja criado como `./output/<resource>`.

## O que é gerado

Ao executar `create` para o recurso `book` serão criados (no diretório `./output/book/`):

- `book.controller.go` — registro de rotas e handlers (funções vazias)
- `book.model.go` — definição inicial do model (`BookModel`)
- `book.repository.go` — esqueleto do repositório
- `book.service.go` — esqueleto do service

O pacote gerado usa o nome do recurso em minúsculas como package (ex.: `package book`).

## Boas práticas

- Revise e adapte os arquivos gerados antes de usá-los em produção.
- Adicione validação, tratamento de erros e integração com banco conforme necessário.

## Contribuição

Pull requests são bem-vindos. Abra uma issue para discutir alterações maiores.

## Licença

MIT — sinta-se à vontade para usar e adaptar o projeto.
