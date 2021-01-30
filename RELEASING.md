# [＜](README.md) Releasing

Nesse repositório usamos o [Semantic Release](https://github.com/semantic-release/semantic-release), um pacote **npm** que permite automatizar a geração de release para projetos versionados com o Git. Sua utilização vem em conjunto com alguns plugins do seu ecossistema, que permitem integrar outros serviços com esse processo, como: Docker, GitLab, GitHub... Por ser bastante [configurável](.releaserc.json), conseguimos torná-lo parte do workflow de desenvolvimento, sendo totalmente plugável em diversos tipos de projetos, independente da linguagem de programação ou propósito.

Nossa ideia com esse setup é fazer com que o desenvolvedor foque no seu desenvolvimento, sem perder muito tempo executando uma série de scripts manuais e descobrindo qual vai a próxima versão da sua release, gerando menos risco de erro humano e faverendo automações a nível de CI.

## ➤ Overview

Algumas convenções conhecidas são utilizadas para introduzirmos boas práticas no fluxo de desenvolvimento desse projeto. Elas são:

- [🗲 Semantic Versioning](https://semver.org/spec/v2.0.0.html)
- [🗲 Convetional Commits](https://www.conventionalcommits.org/en/v1.0.0)

O **Semantic Versioning** é usado para dar um significado padronizado para nosso controle de versão, já o **Convetional Commits** para tornar cada mensagem de commit legível (inclusive essa convenção se encaixa perfeitamente o **Semantic Versioning**, descrevendo os recursos, correções e alterações importantes feitas nas mensagens de commit).

Usando essas convenções postas, conseguimos criar um processo que automaticamente descobre qual vai ser a próxima versão do teu software com base nas mensagens de commit que você cria, não necessitando de nenhuma interação humano a não ser a interação com a criação da mensagem de commit seguindos os padrões e refletindo a real intenção daquela alteração, que já é algo do dia a dia para alguem que utilizar o Git como sistema de controle de versão.

## ➤ Plugins

Aqui está a lista de todos os plugins do ecossistema do **Semantic Release** e de outras automatizações, que utilizamos nesse projeto:

### ⚡ Conventional Commits

Lista de plugins que nos ajudam a configurar o **Convetional Commits**.

- [`@commitlint/cli`](https://github.com/semantic-release/git)
- [`@commitlint/config-conventional`](https://github.com/semantic-release/git)

### ⚡ Suporte ao Workflow

Lista de plugins que nos ajudam com o processo de validação e construção das nossas mensagens de commit.

- [`husky`](https://github.com/semantic-release/git)
- [`commitizen`](https://github.com/semantic-release/git)

#### ⚡ Semantic Release

Lista de plugins do **Semantic Release** para configurarmos no arquivo [.releaserc.json](.releaserc.json) e gerar nossa tag no git, nossa release e nosso [CHANGELOG.md](CHANGELOG.md).

- [`semantic-release`](https://github.com/semantic-release/git)
- [`@semantic-release/git`](https://github.com/semantic-release/git)
- [`@semantic-release/github`](https://github.com/semantic-release/github)
- [`@semantic-release/changelog`](https://github.com/semantic-release/changelog)
- [`@semantic-release/commit-analyzer`](https://github.com/semantic-release/commit-analyzer)
- [`@semantic-release/release-notes-generator`](https://github.com/semantic-release/release-notes-generator)

## ➤ Configuração

A configuração desses plugins é realizadas nos arquivos:

- [package.json](package.json): Contém todos os pacotes **npm**, além da configurações de alguns pacotes instalados e alias para scripts.
- [.releaserc.json](.releaserc.json): Contém a configuração do [`semantic-release`](https://github.com/semantic-release/git) e seus plugins.

## ➤ Requisitos

Antes de seguir o processo de geração de uma nova release, o desenvolvedor precisa 💥 obrigatoriamente 💥 confirmar a realização das seguintes tarefas:

1. 📝 [Instalar o NPM](https://www.npmjs.com/get-npm) ou [Yarn](https://classic.yarnpkg.com/en/docs/install/#windows-stable).
1. 📝 [Gerar um GitLab token](https://goreleaser.com/environment/#github-token) e introduzí-lo no seu bash por meio de uma variável chamada `GL_TOKEN` ou `GITLAB_TOKEN`.
1. 📝 Rodar o comando `npm install` ou `yarn install` na raiz do repositório para instalar todas os pacotes npm mapeadas no [package.json](package.json).

🚨 Vale a pena ressaltar que o [.gitignore](.gitignore) inclui alguns arquivos que são gerados nesse processo e que são ignorados quando é realizado um novo commit (eles não são considerados durante a confirmação) 🚨

## ➤ Gerando uma nova release

As releases poderão ser geradas automaticamente pelo CI ou de forma manual pelo desenvolvedor rodando os scripts definidos no [package.json](package.json).

### Automática pelo CI

1. 🔖 Realize o merge de todos os PRs pretendidos para a release dentro da branch `main`
1. ✨ Faça o checkout e update da branch `main` e valide se todos os testes passarão
    * `git checkout main`
    * `git pull`
    * `make all`
1. 🐎 Faça seu commit seguindo o **Conventional Commits** e depois o push para o GitHub
1. 👌 Vá para o GitHub e valide se a release foi gerada com sucesso:
    * Valide o status do job de release via [Actions](https://github.com/lpmatos/loli/actions?query=workflow%3ARelease)
    * Valide se a release existe com os assests e [CHANGELOG](CHANGELOG.md): https://github.com/lpmatos/loli/releases
1. 📦 Anuncie para seu time o geração da nova release
1. 🎉 Comemore

### Manual

1. 🔖 Realize o merge de todos os PRs pretendidos para a release dentro da branch `main`
1. ⚡ Execute `npm run release` ou `yarn run release` para acionar a ferramenta do **Semantic Release**
1. 📦 Anuncie para seu time o geração da nova release
1. 🎉 Comemore
