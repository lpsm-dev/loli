# Releasing

O [Loli CLI](https://github.com/lpmatos/loli) usa o [Semantic Release](https://github.com/semantic-release/semantic-release), um pacote de gerenciamento de versão e publicação, em conjunto com alguns plugins do seu ecossistema, para acelear e automatizar o processo de versionamento e release do código. Por ser bastante [configurável](.releaserc.json), conseguimos torná-lo parte do workflow de desenvolvimento, sendo totalmente plugável em diversos tipos de projeto, independente da linguagem de programação ou propósito.

Nossa ideia com esse setup é fazer com que o desenvolvedor foque no seu desenvolvimento, sem perder muito tempo executando scripts manuais ou descobrir qual será a próxima versão de publicação, gerando menos risco de erro humano.

## ➤ Overview

Aqui usamos algumas convenções para introduzir boas práticas no workflow de desenvolvimento. O [Semantic Versioning](https://semver.org/spec/v2.0.0.html) é usado para dar um significado padronizado para seu controle de versão e o [Convetional Commits](https://www.conventionalcommits.org/en/v1.0.0) para tornar cada mensagem de commit legível por um ser humano (inclusive essa convenção se encaixa perfeitamente no SemVer, descrevendo os recursos, correções e alterações importantes feitas nas mensagens de commit).

Usando essas duas convenções conseguimos usar algumas ferramentas para descobrir automaticamente qual será a próxima versão do teu software com base nas suas mensagens de commit, não necessitando de nenhuma interação humano, a não ser a criação da mensagem de commit seguindos os padrões e refletindo a real intenção daquela alteração.

## ➤ Plugins

Aqui estão os plugins utilizados:

### ⚡ Suporte ao Workflow

- [`husky`](https://github.com/semantic-release/git)
- [`commitizen`](https://github.com/semantic-release/git)

### ⚡ Conventional Commits

- [`@commitlint/cli`](https://github.com/semantic-release/git)
- [`@commitlint/config-conventional`](https://github.com/semantic-release/git)

#### ⚡ Semantic Release

- [`semantic-release`](https://github.com/semantic-release/git)
- [`@semantic-release/git`](https://github.com/semantic-release/git)
- [`@semantic-release/github`](https://github.com/semantic-release/github)
- [`@semantic-release/changelog`](https://github.com/semantic-release/changelog)
- [`@semantic-release/commit-analyzer`](https://github.com/semantic-release/commit-analyzer)
- [`@semantic-release/release-notes-generator`](https://github.com/semantic-release/release-notes-generator)

## ➤ Configuração

A configuração desses plugins são realizadas nos arquivos [package.json](package.json) e [.releaserc.json](.releaserc.json).

## ➤ Requisitos

Antes de saber como gerar uma nova release, o desenvolvedo precisa obrigatoriamente confirmar a realização das seguintes tarefas:

1. 📝 [Instalar NPM](https://www.npmjs.com/get-npm) ou [Yarn](https://classic.yarnpkg.com/en/docs/install/#windows-stable).
1. 📝 [Gerar um GitHub token](https://goreleaser.com/environment/#github-token) e introduzí-lo no seu bash por meio de uma variável chamada `GH_TOKEN` ou `GITHUB_TOKEN`.
1. 📝 Rodar o comando `npm install` ou `yarn install` na raiz do repositório para instalar todas as dependências mapeadas no [package.json](package.json).

🚨 Vale a pena ressaltar que o arquivO [.gitignore](.gitignore) inclui alguns arquivos que são gerados nesse processo e que são ignorados quando é realizado um novo commit (eles não são considerados durante a confirmação) 🚨

## ➤ Gerando uma nova release

1. Merge all PRs intended for the release into the `main` branch
1. Checkout and update the main branch and ensure all tests are passing:
    * `git checkout main`
    * `git pull`
    * `make all`
1. Go to GitHub and check that the release was successful:
    * Check the release CI job status via the [Actions](https://github.com/lpmatos/loli/actions?query=workflow%3ARelease) tab
    * Check the release exists with valid assets and changelog: https://github.com/lpmatos/loli/releases
1. Announce release internally via Slack
1. Celebrate :tada:
