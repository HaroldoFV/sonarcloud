# ci-sonarcloud

Este é um projeto simples para configurar a integração contínua com o SonarCloud utilizando Golang. O projeto está configurado para rodar verificações de qualidade de código e testes automatizados em pull requests para a branch `develop`.

## Requisitos

- Conta no GitHub
- Conta no SonarCloud
- Projeto Golang configurado

## Configuração

### 1. Adicione os segredos no GitHub

Você precisará adicionar os seguintes segredos no seu repositório GitHub:
- `GITHUB_TOKEN`: Token padrão do GitHub para acessar informações do PR.
- `SONAR_TOKEN`: Token gerado no SonarCloud para autenticação.

### 2. Crie o arquivo de workflow do GitHub Actions

Crie um arquivo chamado `.github/workflows/ci-sonarcloud.yml` no seu repositório com o seguinte conteúdo:

```yaml
name: ci-sonarcloud

on: 
  pull_request:
    branches:
      - develop

jobs:
  run-ci:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: 1.22.2
      - run: go test -coverprofile=coverage.out
      - name: SonarCloud Scan
        uses: SonarSource/sonarcloud-github-action@master
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}  # Needed to get PR information, if any
          SONAR_TOKEN: ${{ secrets.SONAR_TOKEN }}
