# go-migrations-sqlc

* Projeto baseado no projeto desenvolvido durante o módulo SQLC do Curso GoExpert, com a adição do Docker para rodar a aplicação, banco de dados, migrations e SQLC.

## Pré-requisitos

- Docker
   - A aplicação não depende de nenhuma tecnologia além do Docker, o que facilita a execução, independente da máquina.

## Configuração inicial

1. Clone este repositório:

   ```bash
   git clone https://github.com/Solon97/go-migrations-sqlc.git
   ```
2. Acesse o diretório do projeto:

   ```bash
   cd go-migrations-sqlc
   ```
3. Crie o arquivo .env:

   ```bash
   cp .env.example .env
   ```
4. Inicie o ambiente Docker:

   ```bash
   docker-compose up -d
   ```

## Uso
* Ao executar a aplicação, as migrations serão executadas e serão inseridos dados no banco de dados.
* Estão inclusos no arquivo Makefile, os comando para criação e revert das migrations, assim como o comando de geração dos arquivos SQLC.