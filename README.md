# Checkpoint 1 - Funcao Serverless na Nuvem

Este projeto contem uma funcao serverless simples que responde a requisicoes HTTP e foi implantada em ambiente de nuvem.

## Provedor Utilizado
* GCP

## Como rodar localmente

### Pre-requisitos
* Goland

### Passo a passo
1. Clone o repositorio para sua maquina:
git clone https://github.com/Doebber/Checkpoint-1.git

2. Entre na pasta do projeto:
```bash
cd Checkpoint-1
```

3. Instale as dependencias do projeto:
```bash
go mod tidy
```
4. Rode o servidor de testes local:
```bash
go run main.go
curl http://localhost:8080/
```