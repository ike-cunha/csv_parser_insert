<h1 align="center">Golang - Parse and Persist an CSV 🛠️</h1>
<p align="center">Serviço de manipulação de dados e persistência em base de dados relacional.</p>

<p align="center">
 <a href="#objetivo">Objetivo</a> •
 <a href="#funcionalidades">Funcionalidades</a> •
 <a href="#pré-requisitos">Pré-requisitos</a> •
 <a href="#tecnologias">Tecnologias</a> •
 <a href="#going Further">Tecnologias</a> • 
 <a href="#autor">Autor</a>
</p>

<h4 align="center"> 
	🧪 Em beta ⚗️
</h4>

### Objetivo

Manipular e persistir dados em base de dados relacional.

### Funcionalidades 🛠️

- [x] API p/ envio de arquivo
- [x] Parse dos dados
- [X] Persistência dos dados no DB
- [X] Higienização da base

### Pré-requisitos

Antes de começar, você vai precisar ter instalado em sua máquina as seguintes ferramentas:
[Docker](https://www.docker.com), [Golang](https://golang.org).

### 🎲 Para Rodar a aplicação

```bash
# Clone este repositório
$ git clone <https://github.com/ike-cunha/csv_parser_insert>

# Acesse a pasta do projeto no terminal/cmd
$ cd csv_parser_insert

# Inicie o Docker-Compose
$ docker-compose run --service-ports web bash

# Quando o container estiver iniciado, execute o comando
$ go build

# Um arquivo com o nome csv_parser_insert.exe surgirá, execute
$ ./csv_parser_insert

# O servidor inciará na porta:8080 - acesse <http://localhost:8080>
```

### Tecnologias 💻

As seguintes ferramentas foram usadas na construção do projeto:

- [Golang](https://golang.org/)
- [Docker](https://www.docker.com)
- [PostgreSQL](https://www.postgresql.org)

### Autor
---

 <img style="border-radius: 50%;" src="https://avatars.githubusercontent.com/u/23556713?s=400&u=6464c4e6297b42a9761f0964bc3bc3dd18bda537&v=4" width="100px;" alt=""/>
 <sub><b>Henrique Cunha</b></sub>

[![Linkedin Badge](https://img.shields.io/badge/-Henrique-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/henriquecunha/)](hhttps://www.linkedin.com/in/henriquecunha/) 
[![Gmail Badge](https://img.shields.io/badge/-henrique.eccunha@gmail.com-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:henrique.eccunha@gmail.com)](mailto:henrique.eccunha@gmail.com)