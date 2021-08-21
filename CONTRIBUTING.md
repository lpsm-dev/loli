Primeiramente, obrigado 🎉! É muito gratificante saber que podemos contar com novas features e pull-requests do nosso time! Caso sua mudança não seja trivial, abra uma issue nesse projeto para podermos discutir a sua ideia e estratégia de implementação. Será incrível poder interagir e se conectar com novos contribuidores! Esperamos que todos possam ajudar nessa evolução 🤗!

## ➤ Overview

Nosso objetivo aqui é gerar valor para projetos de desenvolvimento com múltiplas participações. Pensando nisso elaboramos um **CONTRIBUTING.md** para abordarmos as melhores práticas a serem seguidas durante o processo de desenvolvimento, definindo o passo a passo de contribução para qualquer novo contribuidor, desde de o que ele precisa ter em sua máquina, até o que ele precisa fazer para gerar uma nova release.

Além dessa padronização é imprescindível uma boa comunicação interna com os envolvidos no projeto, uma vez que de nada adianta ter um processo de trabalho padronizado e ninguém do time se comunicar. Por isso formentamos a criação de comunidades em alguma ferramenta de bate-papo, como: Slack, Teams, Telegram e etc, mas isso fica a cargo do time do projeto.

## ➤ Conventions, Tools and Packages

Antes de qualquer commit para esse repositório é de extrema importância que o contribuidor saiba o que ele precisa refletir em seu ambiente local para que a sua contribução esteja seguindo as nossas diretrizes de trabalho. Portanto, nesse tópico listamos todas as ferramentas, pacotes e padrões utilizados pelos membros que colaboram aqui.

### Conventions

- ⮚ [Semantic Versioning](https://semver.org/)
- ⮚ [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

### Tools

- ⮚ Gerenciador de pacotes [NPM](https://www.npmjs.com/get-npm) ou [Yarn](https://yarnpkg.com/getting-started/install) para instação das dependências de suporte.
  - Este projeto não é um projeto **node.js**. O arquivo `package.json` é usado apenas para definir alguns metadados e dependências que darão suporte para o nosso workflow no git.
- ⮚ Automatizador de tarefas locais [make](https://www.gnu.org/software/make/manual/make.html).
- ⮚ Scan de segredos [gitLeaks](https://github.com/zricethezav/gitleaks).
- ⮚ Mecanismo de controle de git hooks [pre-commit](https://pre-commit.com/#intro).

### Packages

#### Semantic Release

- ⮚ [Semantic Release](https://github.com/semantic-release) + Plugins de configuração
  - [`semantic-release`](https://github.com/semantic-release/semantic-release)
  - [`@semantic-release/git`](https://github.com/semantic-release/git)
  - [`@semantic-release/github`](https://github.com/semantic-release/github)
  - [`@semantic-release/changelog`](https://github.com/semantic-release/changelog)
  - [`@semantic-release/commit-analyzer`](https://github.com/semantic-release/commit-analyzer)
  - [`@semantic-release/release-notes-generator`](https://github.com/semantic-release/release-notes-generator)

#### Conventional Commits

- ⮚ [Commit Lint](https://github.com/conventional-changelog/commitlint) usando o [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).
  - [`commitizen`](https://github.com/commitizen/cz-cli)
  - [`@commitlint/cli`](https://github.com/conventional-changelog/commitlint)
  - [`@commitlint/config-conventional`](https://github.com/conventional-changelog/commitlint)

## ➤ Regras de codificação

Nesse projeto utilizamos a convenção do [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) como boa prática para criação de mensagens de commit, que é totalmente ligada a convenção [SemVer](https://semver.org/), responsável por ditar as regras de versionamentodo/release do código. A partir dessas convenções conseguimos utilizar plugins **npm**, no caso o ecossistema [Semantic Release](https://github.com/semantic-release), para automatizar todo nosso processo de geração de **tag/release**, tudo de forma automática e com base em regras pré-configuradas, podendo ser customizáveis de acordo com o cenário. Essas regras estão definidas no arquivo `.releaserc.json` e `package.json`. Para garantir a consistência do nosso código fonte, lembre-se de seguir essas regras enquanto trabalhar:

- ⮚ Antes de iniciar o desenvolvimento certifique-se de instalar todas as ferramentas abordadas na seção [Conventions, Tools and Packages](#-conventions-tools-and-packages).
- ⮚ Ao realizar um `git commit` opte por usar uma mensagem seguindo o padrão do [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).
- ⮚ Nunca realize uma operação de commit em uma branch estável.
- ⮚ Verifique previamente seus arquivos commitados para nunca enviar algo com conteúdo sensível para o repositório remoto.
- ⮚ Certifique-se de dar um `git rebase` antes de mesclar sua branch, isso vai evitar possíveis conflitos além de gerar um histórico linear das modificações.
- ⮚ Faça um `git squash` dos commits, isso é uma boa prática para manter um histórico de commits mais limpo.
- ⮚ Ao abrir um merge-request, certifique-se de comunicar os `maintainers-owners` do projeto.

## ➤ Submission Guidelines

> 1. Crie uma nova **branch** a partir de outra que faça sentido para o seu fluxo de trabalho.
> 1. Leia as regras de contribução.
> 1. Siga a organização do repositório sempre que você for alterar ou adicionar coisas.
> 1. Faça um **commit** com suas alterações.
> 1. Abra um **merge-request** assim que perceber que suas alterações estão prontas para serem promovidas.
> 1. Espere até que seu **merge-request** seja aprovada... 🚀

**Lembre-se**: Não existe código ruim, temos diferentes formas de resolver um mesmo problema. 😊

### Add to git and push

📝 Você precisa mandar suas modificações para o servidor do Git assim que terminá-las. Para isso, faça o seguinte:

```bash
git add -f .
git commit -m "chore(initial): include config files"
git push -u origin <branch-alvo>
```

### Submitting a Merge Request (MR)

Antes de aceitar um **merge-request**, preferimos que você esmague seus commits em um único commit utilizando o `git stash`. Essa ação visa garantir um histórico de commits mais limpo. A maioria das ferramentas Git já possuem integração pela própria UI para já acionar esse comando durante a abertura dessa mesclagem e no GitLab não é diferente.

Assumindo que o **code-review** foi concluído e os teste foram passados (a pipeline foi bem sucedida), sua mudança deve ser mesclada o mais rápido possível para a branch alvo.

Lembre-se de sempe marcar para delete a branch caso ela não seja a default. Não recomendamos guardar branches de curta durança em seu histórico de branchs (pode causar conflitos e confusão no futuro).

## ➤ Commit Message Guidelines

Como dito nas regras de codificação, temos regras muito precisas sobre como nossas mensagens `git commit` devem ser formatadas. Isso leva a mensagens mais legíveis e fáceis de seguir ao examinar o histórico de contribução do projeto. Também usamos essa convenção de mensagens git commit para estruturar e gerar um arquivo de log com as alterações do nosso projeto, o famoso **CHANGELOG.md**.

### Overview

A especificação do [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) é uma convenção para suas mensagens de commit. Ele fornece um conjunto fácil de regras para criar um histórico de commit explícito; o que torna mais fácil escrever ferramentas automatizadas. Essa convenção se encaixa com o [SemVer](https://semver.org/), descrevendo os recursos, correções e alterações importantes feitas nas mensagens de commit.

### Commit Message Format

A mensagem de commit deve ser estruturada da seguinte forma:

```text
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

Cada mensagem de commit consiste na estrutura acima. Temos um `header`, um `body`e um `footer`. O `header` tem um formato especial que inclui um `type`, um `scope` e uma `description`, sendo o `header` um campo obrigatório, porém seu `scope` opcional.

Para poder comuniar a intenção da sua alteração, a mensagem de commit contém os seguintes elementos estruturais:

1. **fix**: um commit do tipo `fix` corrige um bug em seu código (isso se correlaciona com o **PATCH** no **semantic versioning**).
1. **feat**: um commit do tipo `feat` introduz uma nova feature (recurso) em seu código (isso se correlaciona com o **MINOR** no **semantic versioning**)
1. **BREAKING CHANGE**: um commit que tem seu `footer` com a mensagem `BREAKING CHANGE` ou possui um `!` após o type ou scope, introduz uma mudança significativa em seu código (isso se correlaciona com o **MAJOR** no **semantic versioning**). Uma `BREAKING CHANGE` pode fazer parte de qualquer tipo de commit.
1. Existem outros tipos de commit além do `fix` e `feat`. São permitidos vários tipos, como - `build`, `chore`, `ci`, `docs`, `style`, `refactor`, `perf`, `test` e outros. A lista completa é baseada na [convenção do angular](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#-commit-message-guidelines) e pode ser vista em [@commitlint/config-conventional](https://github.com/conventional-changelog/commitlint/tree/master/%40commitlint/config-conventional).

Qualquer mensagem de commit não pode ter mais que 100 caracteres! Isso quebraria a nossa convenção. Esse limite permite que a mensagem seja mais fácil de ler em várias ferramentas Git.

```bash
C:\>  git add .
C:\>  git commit -m "commit"


husky > commit-msg (node v12.14.0)
⧗   input: commit
✖   subject may not be empty [subject-empty]
✖   type may not be empty [type-empty]

✖   found 2 problems, 0 warnings
ⓘ   Get help: https://github.com/conventional-changelog/commitlint/#what-is-commitlint

husky > commit-msg hook failed (add --no-verify to bypass)
```

Usando o commitzen para te ajudar a construir a mensagem de commit:

```
C:\>  git add .
C:\>  npm run cm


cz-cli@4.0.3, cz-conventional-changelog@3.2.0

? Select the type of change that you're committing: (Use arrow keys)
> feat:     A new feature
  fix:      A bug fix
  docs:     Documentation only changes
  style:    Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
  refactor: A code change that neither fixes a bug nor adds a feature
  perf:     A code change that improves performance
  test:     Adding missing tests or correcting existing tests
  build:    Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
  ci:       Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)
  chore:    Other changes that don't modify src or test files
  revert:   Reverts a previous commit
```

### TL;DR

- Commit messages starting with `fix: ` trigger a patch version bump
- Commit messages starting with `feat: ` trigger a minor version bump
- Commit messages starting with `BREAKING CHANGE: ` trigger a major version bump.

## ➤Automatic versioning

A cada novo commit que você realiza é uma nova intenção que você gera. No final do dia você poderá ou não querer gerar uma nova release do seu código, mas uma coisa é certa: não será você que determinará a próxima versão de tag. Calma, isso não quer dizer que vamos rodar um monte de scripts para efetuar isso, muito pelo contrário, estaremos nos apoiando em convenções consolidadas no mercado de software, convenções essas que ao usar as devidas ferramentas, você consegue automaticamente determinar a sua próxima versão depois de ter realizado uma série de commits. As convenções nos debatemos acima e também qual é o conjunto de ferramenta/plugins que estamos utilizando, mas ainda temos alguns pontos para retratar:

- Sempre que o seu código estiver na branch `release` você será capaz de gerar uma tag `release candite` ao executar o [`semantic-release`](https://semantic-release.gitbook.io/semantic-release/).
- Sempre que o seu código estiver na branch `main` você será capaz de gerar uma tag `stable` ao executar o [`semantic-release`](https://semantic-release.gitbook.io/semantic-release/).

A cada novo commit para uma dessas branches, em um determinado momento poderemos rodar o \[` semantic-release`\] (https://semantic-release.gitbook.io/semantic-release/) que determina e envia uma nova tag de versão (se houver) com base no última versão marcada e os novos commits enviados. Observe que isso significa que se um **merge-request** contém, por exemplo, vários commits de `feat:`, apenas um salto de versão **minor** ocorrerá na mesclagem. Se seu **merge-request** inclui vários commits, você pode preferir ignorar o prefixo em cada commit individual e, em vez disso, adicionar um commit vazio sumarizando suas alterações como:

```
git commit --allow-empty -m '[BREAKING CHANGE|feat|fix]: <changelog summary message
```

### CI

A sua pipeline poderá ou não incluir um template que possui esse job. Caso inclua, todo esse processo ocorre via pipeline e não irá necessitar de nenhuma ação local por parte do contribuidor.

## ➤ Release Steps

```bash
export GITLAB_TOKEN=""
make release-debug
git add . && git commit -am "chore: bump version file"
make release
git push --all
git pull --all
```
