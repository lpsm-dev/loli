<!-- BEGIN_DOCS -->

<a name="readme-top"></a>

[◀ Voltar](README.md)

<div align="center">

<img alt="contributing" src="https://github.com/lpsm-dev/lpsm-dev/blob/98272299ea611ba50254b132490ea385149dc5cf/.github/assets/contributing.png" width="225"/>

**Diretrizes para o processo de contribuição**

</div>

Obrigado por considerar contribuir com este projeto! Seguir nossas diretrizes facilitará seu onboarding e tornará sua colaboração mais eficaz.

# Sumário

<details>
  <summary><strong>Expandir</strong></summary>

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Práticas](#pr%C3%A1ticas)
  - [Geral](#geral)
  - [Comunicação](#comunica%C3%A7%C3%A3o)
- [Setup](#setup)
  - [DevBox](#devbox)
  - [Direnv](#direnv)
  - [Task](#task)
- [Commit Messages](#commit-messages)
  - [Type](#type)
  - [Scope](#scope)
  - [Description](#description)
- [MR Process](#mr-process)
  - [Steps](#steps)
  - [Reviewing](#reviewing)
- [Versioning Process](#versioning-process)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

<p align="right">(<a href="#readme-top">back to top</a>)</p>

</details>

# Práticas

## Geral

- Sempre crie branchs descritivas, exemplo: `feature/nova-funcionalidade` ou `fix/corrige-bug`.
- Evite commits muito grandes. Prefira commits pequenos e atômicos.
- Não reinvente a roda. Se você pesquisou e viu que já existe uma solução bem estabelecida para o seu problema, use-a. Isso economiza tempo e recurso.

## Comunicação

- Se você não conseguir continuar uma tarefa, informe imediatamente sua equipe. A comunicação rápida evita atrasos e permite que outras pessoas te ajudem a resolver os problemas com mais rapidez.
- Minimize o uso de IA na comunicação diária com a equipe. Valorizamos interações reais e genuínas.
- Seja objetivo na sua comunição quando precisa de ajuda (isso não significa ser rude).
- A comunicação assíncrona é uma grande aliada para equipes remotas. Para mais detalhes, clique [aqui](https://nohello.net/en/).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Setup

Para você começar a contribuir, é essencial configurar seu ambiente de trabalho local:

- **Instalar ferramentas**: Certifique-se de ter instalado todas as ferramentas de linha de comando que o projeto requer.
- **Configurar variáveis de ambiente**: Ajuste as variáveis de ambiente para garantir que seu sistema esteja preparado para rodar o projeto.
- **Executar scripts de automação**: Rode os scripts fornecidos para configurar dependências, inicializar bancos de dados e outras tarefas automatizadas.

> [!NOTE]
> Lembre-se: cada projeto tem seu próprio contexto e necessidades!

Pensando nisso, elaboramos as etapas abaixo para te guiar.

## DevBox

O **DevBox** é uma ferramenta CLI que cria ambientes de desenvolvimento isolados e reproduzíveis, sem precisar usar containers Docker ou a linguagem Nix de forma nativa.

> [!NOTE]
> Use essa opção se você não quiser instalar muitas ferramentas CLI diretamente em seu ambiente de trabalho.

Siga essas etapas para configurar seu ambiente:

- Instale o [devbox](https://www.jetify.com/devbox/docs/installing_devbox/):

```bash
curl -fsSL https://get.jetpack.io/devbox | bash
```

- Inicialize seu projeto:

```bash
devbox init
```

- Adicione os pacotes que deseja (vai mudar de projeto para projeto). Ex:

```json
{
  "$schema": "https://raw.githubusercontent.com/jetify-com/devbox/0.10.7/.schema/devbox.schema.json",
  "packages": ["awscli2@latest", "kubectl@1.29.3", "kubernetes-helm@3.14.3"],
  "shell": {
    "init_hook": ["echo 'Welcome to devbox!' > /dev/null"],
    "scripts": {
      "test": ["echo \"Error: no test specified\" && exit 1"]
    }
  }
}
```

- Execute o seguinte comando para inicializar o shell temporário:

```bash
devbox shell
```

Com isso, podemos garantir que todos no projeto tenham as mesmas ferramentas nas mesmas versões, necessárias para o processo de desenvolvimento.

> [!NOTE]
> Se você precisar de mais detalhes sobre essa configuração, verifique o arquivo [devbox.json](devbox.json) do seu projeto. Caso não exista, crie ele seguindo o passo a passo descrito acima.

## Direnv

Para você automatizar certas ações em seu terminal sempre que você for trabalhar nesse projeto, configure o **Direnv**. Essa ferramenta vai ajustar o seu shell conforme o seu diretório atual. Assim, sempre que você entrar na pasta do projeto, o **Direnv** fará algo, como: carregará as variáveis definidas no `.env` ou disparar o shell do **DevBox**.

Siga essas etapas para configurar seu ambiente:

- Acesse a documentação do [direnv](https://direnv.net/docs/installation.html) e siga as instruções para instalá-lo.

- Após a instalação, crie um arquivo `.env` na raiz do seu projeto para armazenar as variáveis de ambiente utilizadas.

- Crie o arquivo `.envrc` com o seguinte conteúdo:

```bash
# Dotenv Support
[[ ! -f .env ]] || dotenv .env

# Devbox Support
has devbox && eval "$(devbox generate direnv --print-envrc)" && exit 0
```

- A primeira vez que você criar ou modificar um arquivo `.envrc`, você precisará autorizá-lo com o comando:

```bash
direnv allow
```

Seguindo essas etapas, quando você navegar para a pasta do seu projeto, as variáveis de ambiente serão carregadas automaticamente e o **DevBox** será inicializado.

> [!NOTE]
> Se você precisar de mais detalhes sobre esse configuração, verifique o arquivo [.envrc](.envrc) do seu projeto.

## Task

A ferramenta **task** oferece uma maneira conveniente de definir e gerenciar tarefas específicas do projeto, facilitando a automatização de scripts comuns e simplificando os fluxos de trabalho de desenvolvimento.

> [!NOTE]
> É semelhante à ferramenta `make`, que é utilizada principalmente para automatizar tarefas.

Siga essas etapas para configurar seu ambiente:

- Certifique-se de que você instalou o comando `task` seguindo as etapas de configuração do **DevBox**.
  - Caso não tenha seguido, acesse a documentação do [task](https://taskfile.dev/installation/) e siga as instruções para instalá-lo.
- Execute o comando `task` no diretório raiz do projeto para ver todos os comandos disponíveis.

> [!NOTE]
> Se você precisar de mais detalhes sobre cada tarefa definida, verifique o arquivo [Taskfile.yaml](Taskfile.yaml) do seu projeto. Caso não exista, crie ele seguindo o passo a passo descrito acima.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Commit Messages

Nesse projeto, exigimos que **todos os commits** sigam um formato específico de mensagem, o [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/). Com isso, conseguimos:

- **Clareza e Consistência**: As mensagens de commit seguem um formato padrão, facilitando a leitura e a compreensão das mudanças.
- **Automatização**: Permite a automação de processos, como geração de changelogs, versionamento semântico e lançamentos automáticos.
- **Rastreamento de Alterações**: Facilita o rastreamento de mudanças ao longo do tempo, tornando mais fácil identificar o que foi modificado e por quê.
- **Melhoria na Revisão de Código**: Proporciona uma melhor experiência de revisão de código, já que as mudanças são descritas de forma clara e padronizada.
- **Comunicação Eficaz**: Ajuda a todos os membros da equipe a entenderem rapidamente o contexto e o propósito de cada alteração no código.

Veja como é organizado esse formato de commits:

```txt
<type>(optional scope): <description>

[optional body]
```

## Type

Descreve o tipo de alteração do commit. Temos as seguintes opções:

| Tipo         | Descrição                                                                                                                                                                                                                                                             |
| ------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **feat**     | Um novo recurso (adição de um novo componente, fornecimento de novas variantes para um componente existente, etc.).                                                                                                                                                   |
| **fix**      | Uma correção de bug (correção de um problema de estilo, resolução de um bug na API de um componente etc.). Ao atualizar dependências que não sejam de desenvolvimento, marque suas alterações como `fix`.                                                             |
| **docs**     | Alterações somente na documentação.                                                                                                                                                                                                                                   |
| **style**    | Alterações que não afetam o significado do código (espaços em branco, formatação, falta de ponto e vírgula etc.). Não deve ser usado para alterações na interface do usuário, pois essas são alterações significativas; em vez disso, considere usar `feat` ou `fix`. |
| **refactor** | Uma alteração de código que não corrige um bug nem adiciona um recurso.                                                                                                                                                                                               |
| **perf**     | Uma alteração de código que melhora o desempenho.                                                                                                                                                                                                                     |
| **test**     | Adição de testes ausentes ou correção de testes existentes.                                                                                                                                                                                                           |
| **build**    | Alterações que afetam o sistema de build.                                                                                                                                                                                                                             |
| **ci**       | Alterações em arquivos e scripts de configuração de CI/CD.                                                                                                                                                                                                            |
| **chore**    | Outras alterações que não modificam arquivos de origem ou de teste. Use esse tipo ao adicionar ou atualizar dependências de desenvolvimento.                                                                                                                          |
| **revert**   | Reverte um commit anterior.                                                                                                                                                                                                                                           |

## Scope

É qualquer coisa que forneça informações adicionais ou que especifique o local de alteração do seu código. Por exemplo `events`, `kafka`, `dockerfile`, `authorization` e etc. Cada tipo (`type`) de commit pode ter um escopo (`scope`) opcional, cabendo a você adicionar ou omitir essa informação. Por exemplo:

```
feat(login): add route
```

> [!NOTE]
> Use a convenção [PascalCase](https://www.dio.me/articles/camel-case-vs-pascal-case) na hora de definir seu escopo (`scope`).

## Description

É o campo onde você diz o que foi feito no commit, porém de forma breve. Para isso, recomendamos que:

- Priorize descrições em inglês.
- Use o imperativo, tempo presente: "change", não "changed" ou "changing".
- Não coloque a primeira letra em maiúscula.
- Não coloque ponto (.) no final.

> [!NOTE]
> Cada tipo de commit tem um efeito sobre a próxima release que você for lançar.

Existem mais opções, porém essas são as mais comuns. Para mais detalhes, consulte a [documentação oficial](https://www.conventionalcommits.org/en/v1.0.0/).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# MR Process

Ao criar um Merge Request (MR), recomenda-se definir o título conforme a mesma convenção adotada para mensagens de commit. Isso garante que, em caso de **squash merge**, o título do MR possa ser diretamente utilizado como a mensagem final do commit, preservando a padronização do histórico. Essa prática contribui para um log de alterações mais estruturado, conciso e linear, facilitando a rastreabilidade e automação de releases em fluxos de desenvolvimento que seguem padrões como **Conventional Commits**.

## Steps

- Crie uma branch a partir da branch `main`:

```bash
git checkout main
git pull origin main
git checkout -b sua-nova-branch
```

- Trabalhe na branch criada:
  - Realize as alterações necessárias e faça os commits das mudanças.
  - Certifique-se de que seu código atenda os padrões de qualidade estabelecidos.
  - Garanta que seus commits sigam a convenção de commits definida acima.

```bash
git add .
git commit -m "fix: change the commit"
```

- Faça o push da sua branch:

```bash
git push origin sua-nova-branch
```

- Abra uma solicitação de MR:

  - Navegue até o repositório e abra uma nova Merge Request da sua branch para a branch de produção `main`.
  - Adicione uma descrição clara do que foi feito e qualquer informação relevante para o processo de review.
  - Defina o título usando commits convencionais.
  - Marque a opção de remover a branch de origem após a mesclagem.
  - Marque a opção para squash dos commits.

- Revisão e Aprovação:

  - Um mantenedor revisará seu código.
  - Se o código atender aos requisitos e padrões, ele será aprovado.
  - Após a aprovação, seu código será mesclado na branch `main`.

- Finalização:
  - Após a mesclagem, sua branch pode ser deletada se não for mais necessária.
  - Certifique-se de executar o `git pull` na branch `main`.

Seguir este processo garante que as alterações sejam revisadas adequadamente e que o código de produção permaneça estável e com qualidade.

> [!NOTE]
> Se você tiver vários commits em seu PR, é recomendável usar a opção de squash para manter um histórico limpo e organizado.

## Reviewing

Durante o processo de revisão do MR, siga essas políticas:

- Seja respeitoso e construtivo.
- Busque realizar um code review completo, verificando a lógica, a qualidade do código e a conformidade com os padrões estabelecidos.
- Sugira alterações em vez de simplesmente comentar os problemas encontrados.
- Exigimos pelo menos um aprovador no MR, que não seja o autor.
- Se não tiver certeza sobre algo, pergunte ao autor do MR.
- Se você estiver satisfeito com as alterações, aprove o MR.
- Preze por pair programming, sempre que possível.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Versioning Process

Esse projeto segue a especificação [SemVer](https://semver.org/). Ou seja, utilizamos um sistema de versionamento semântico para controlar as versões do projeto. Isso nos ajuda a comunicar melhor as mudanças feitas no código e a garantir a compatibilidade entre as versões. Esse sistema é composto por três números (`x.x.x`), que representam respectivamente:

1. **Major**: Quando fazemos alterações incompatíveis com a versão atual. Ex: `1.2.0 → 2.0.0`
2. **Minor**: Quando adicionamos funcionalidades de forma compatível com a versão atual. Ex: `1.2.0 → 1.3.0`
3. **Patch**: Quando corrigimos bugs de forma compatível com a versão atual. Ex: `1.2.0 → 1.2.1`

Nós usamos o [Semantic Release](https://semantic-release.gitbook.io/semantic-release/) para automatizar esse processo de versionamento. Ele analisa os commits desde a última versão e determina automaticamente o próximo número da versão com base nas alterações feitas. Isso garante que as versões sejam consistentes e sigam as regras do SemVer.

Vale ressaltar que o Semantic Release é acionado automaticamente quando um MR é mesclado na branch `main` ou quando existe um commit direto na branch `main`. Portanto, é importante seguir as convenções de commit para garantir que as versões sejam geradas corretamente.

> [!NOTE]
> O Semantic Release é um ferramental poderoso, e possui um sistema de plugins que podem ser configurados para atender as necessidades específicas do projeto. Para mais detalhes, consulte a [documentação oficial](https://semantic-release.gitbook.io/semantic-release/).

<p align="right">(<a href="#readme-top">back to top</a>)</p>
<!-- END_DOCS -->
