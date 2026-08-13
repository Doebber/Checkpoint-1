# checkpoint1-go

API em Go (Fiber) que responde em `GET /` com a data/hora atual em
`America/Sao_Paulo`. Deploy via Cloud Run (aceita container, ao contrário do
Cloud Functions, que exigiria outro formato de handler).

## Rodar localmente

```bash
go mod tidy
go run main.go
curl http://localhost:8080/
```

## Deploy no Cloud Run

```bash
gcloud config set project project-ab4fa986-e441-43f4-9fe
gcloud services enable run.googleapis.com cloudbuild.googleapis.com

gcloud run deploy checkpoint1-go \
  --source . \
  --region southamerica-east1 \
  --allow-unauthenticated \
  --port 8080
```

O comando imprime a URL pública ao final. Teste com `curl <url>/`.

Para atualizar, rode o mesmo comando de deploy de novo. 

Para remover:

```bash
gcloud run services delete checkpoint1-go --region southamerica-east1
```