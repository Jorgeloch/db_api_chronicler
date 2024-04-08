# Chronicler
Este é um projeto implementado e utilizado como atividade para a matéria de banco de dados
## Instruções de uso
### Tecnologias e Versões
Neste projeto estamos utilizando as seguintes tecnologias:
-  Golang, na sua versão `1.21.8`
-  Postgres, na sua major `16`
### executando o projeto
Antes de executar o projeto, é necessário configurar as variáveis de ambiente utilizadas, como mostrado no arquivo `.env.exemple`.
Com o ambiente configurado, para executar o projeto na sua máquina, você ira precisar executar os seguintes comandos
```sh
go mod tidy
```
Este comando irá instalar todas as dependências do projeto
```sh
make migrate_up
```
este comadno irá executar as migrations no seu banco de dados
```sh
make build
```
este comando ira realizar a compilacao do projeto
```sh
./bin/main
```
este comando irá executar o projeto em sua máquina
Ao encerrar o projeto, caso queira desfazer tudo o que foi realizado no banco de dados, basta utilizar o seguinte comando:
```sh
make migrate_down
```
