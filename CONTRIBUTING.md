# [＜](README.md) Contribuindo

Antes de tudo, muito obrigado 🎉! É extremamente contagiante saber que podemos contar com novas features e pull-requests da comunidade! Caso sua mudança não seja algo trivial, abra uma [issue](https://github.com/lpmatos/loli/issues) podermos discutir a sua ideia e sua estratégia de implementação. Será incrível poder interagir e se conectar com novos contribuidores!

## ➤ Overview

Nosso principal objetivo aqui é gerar valor para o processo de desenvolvimento. Pensando nisso implementamos um arquivo de contribuição para abordarmos as melhores práticas a serem seguidas durante o desenvolvimento nosso desenvolvimento. O modelo é uma recomendação importante e ainda está em processo de construção. Esperamos que você possa contribuir para essa evolução 🤗!

## ➤ Menu

1. [Código de conduta](#-código-de-conduta)
1. [Regras de codificação](#-regras-de-codificação)
1. [Submission Guidelines](#-submission-guidelines)
1. [Commit Message Guidelines](#-commit-message-guidelines)

## ➤ Código de conduta

Aqui estão todas as diretrizes que gostaríamos que você seguisse caso queira nos ajudar de alguma forma 😄!

## ➤ Regras de codificação


Para garantir a consistência do nosso código fonte, lembre-se de seguir essas regras enquanto trabalhar:

- Todos as **features** ou **fixs** (correções de bugs) **devem ser testados**!
- A **pipeline** deve possui um status de sucesso antes de qualquer **pull-request** ser aprovado, caso contrário, o código não é mesclado em seu branch alvo.
- Todas as **features** e **fixs** são criadas a partir da branch **main** e são mescladas novamente na banch **main** assim que o **pull-request** for aprovado.
- Certifique-se de dar um `git rebase` antes de mesclar sua **feature** ou **fix** na branch **main**, isso vai evitar possíveis conflitos, além de gerar um histórico linear das modificações.
- Sempre que o código vai para a branch **main** uma nova stable tag deve ser criada.
- Nesse projeto utilizamos a convenção do [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) como boa prática de criação de mensagens de commit, que é totalmente ligado a convenção [SemVer](https://semver.org/), que é responsável por ditar as regras de versionamentodo/release do código. A partir dessas convenções conseguimos utilizar plugins **npm** para automatizar nosso processo de geração de **tag/release**, tudo de forma automática e com base em regras pré-configuradas, podendo ser customizáveis de acordo com o cenário.
- Utilize a estratégia de commits atômicos onde a cada pequena alteração você cria um commit. Evite criar commits com muitas alterações.

## ➤ Teste suas mudanças

Assim que você criar um **pull-request** da sua **features** ou **fix** para a branch **main**, uma **pipeline** será criada para validar as mudanças. Nenhum **pull-request** será aprovada caso a **pipeline** não seja bem sucedida. É de vital importância que a branch `main` passe nos testes o tempo todo, caso contrário nada irá para produção. Sempre que possível, adicione novos testes para garantir que seu código fique o melhor possível.

## ➤ Submission Guidelines

>
> 1. Faça um **fork** ou crie uma **branch** **feature** ou **fix**. 
> 1. Leia as [regras de contribução](CONTRIBUTING.md).
> 1. Siga a organização do repositório sempre que você for alterar ou adicionar arquivos ou diretórios.
> 1. Faça o **commit** com suas alterações.
> 1. Abra um **pull-request** assim que perceber que suas alterações estão prontas para serem promovidas.
> 1. Espere até que seu **pull-request** seja aprovado... 🚀
>

**Lembre-se**: Não existe código ruim, temos diferentes formas de resolver um mesmo problema. 😊

### Add to git and push

📝 Você precisa mandar suas modificações para o servidor do **Git** assim que terminá-las. Para isso, faça o seguinte:

```bash
git add -f .
git commit -m "chore(initial): include config files"
git push -u origin <branch>
```

### Submitting a Merge Request (MR)

Antes de aceitar um **merge-request**, preferimos que você esmague seus commits em um único commit utilizando o `git stash`. Essa ação visa garantir um histórico de commits mais limpo. A maioria das ferramentas Git já possuem integração pela própria UI para já acionar esse comando durante a abertura dessa mesclagem.

Assumindo que o **code-review** foi concluído e os teste foram passados (a pipeline foi bem sucedida), sua mudança deve ser mesclada o mais rápido possível para a branch alvo.

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
