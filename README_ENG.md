<p align="center">
  <img alt="logo" src="https://i.pinimg.com/280x280_RS/d0/13/35/d01335f147c586e56829415e611f0ae7.jpg" width="350px" float="center"/>
</p>

# Welcome to Loli repository

<p align="center">
  <a href="https://spdx.org/licenses/Apache-2.0.html" target="_blank">
    <img alt="License: Apache 2.0" src="https://img.shields.io/badge/License-Apache 2.0-yellow.svg" />
  </a>
  <a href="https://github.com/semantic-release/semantic-release">
    <img alt="semantic-release" src="https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg">
  </a>
  <a href="http://commitizen.github.io/cz-cli/">
    <img alt="Commitizen friendly" src="https://img.shields.io/badge/commitizen-friendly-brightgreen.svg">
  </a>
  <a href="https://www.codacy.com/gh/lpmatos/loli/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=lpmatos/loli&amp;utm_campaign=Badge_Grade">
    <img alt="Codacy Badge" src="https://app.codacy.com/project/badge/Grade/7d69caff8e2646a783681cc765948187">
  </a>
  <a href="https://github.com/lpmatos/loli/actions?query=workflow%3ALinter">
    <img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/lpmatos/loli/Linter">
  </a>
  <a href="https://github.com/lpmatos/loli/actions">
    <img alt="Latest Release" src="https://img.shields.io/github/v/release/lpmatos/Loli">
  </a>
</p>

>
> Loli is a GoLang CLI tool find a anime by images
>

## ➤ Menu

<p align="left">
  <a href="#-description">Description</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-contribute">How to contribute</a>
</p>

## ➤ Getting Started

If you want use this repo first you need to make a **git clone**:

>
> 1. git clone --depth 1 <https://github.com/lpmatos/loli.git> -b master
>

This will give you access to the code on your **local machine**.

## ➤ Description

Some description were.

## ➤ Development with Docker

Steps to build the Docker Image.

<details><summary>🐋 Build</summary>
<p>

Docker commands to build your image:

```bash
docker image build -t <IMAGE_NAME> -f <PATH_DOCKERFILE> <PATH_CONTEXT_DOCKERFILE>
docker image build -t <IMAGE_NAME> . (This context)
```
</p>
</details>

<details><summary>🐋 Run</summary>
<p>
Docker commands to run a container with yout image:

* **Linux** running:

```bash
docker container run -d -p <LOCAL_PORT:CONTAINER_PORT> <IMAGE_NAME> <COMMAND>
docker container run -it --rm --name <CONTAINER_NAME> -p <LOCAL_PORT:CONTAINER_PORT> <IMAGE_NAME> <COMMAND>
```

* **Windows** running:

```bash
winpty docker.exe container run -it --rm <IMAGE_NAME> <COMMAND>
```
</p>
</details>

## ➤ Project organization features

- Default gitignore and editorconfig.
- [Semantic Versioning](https://semver.org/)
- [GitLeaks](https://github.com/zricethezav/gitleaks) file - Scan for secrets using regex and entropy.
- [Semantic Release](https://github.com/semantic-release/semantic-release) + Plugins configuration
- NPM modules automation.
  - [Commitlint](https://github.com/conventional-changelog/commitlint) using [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).
  - Git Hooks with [Husky](https://github.com/typicode/husky).

### Semantic Release Plugins

This shareable configuration use the following plugins:

- [`@semantic-release/git`](https://github.com/semantic-release/git)
- [`@semantic-release/gitlab`](https://github.com/semantic-release/gitlab)
- [`@semantic-release/changelog`](https://github.com/semantic-release/changelog)
- [`@semantic-release/commit-analyzer`](https://github.com/semantic-release/commit-analyzer)
- [`@semantic-release/release-notes-generator`](https://github.com/semantic-release/release-notes-generator)

## ➤ How to contribute

>
> 1. Make a **Fork** or Create a **Feature Branch**.
> 2. Follow the project organization.
> 3. Add the file to the appropriate level folder - If the folder does not exist, create according to the standard.
> 4. Make the **Commit**.
> 5. Open a **Merge Request**.
> 6. Wait for your merge request to be accepted.. 🚀
>

**Remember**: There is no bad code, there are different views/versions of solving the same problem. 😊

## ➤ Add to git and push

📝 You must send your work to [**GitHub**](https://github.com/lpmatos/loli) repo after your changes:

```bash
git add -f .
git commit -m "chore(initial): include config files"
git push -u origin master
```

## ➤ Versioning

🚨 We currently do not have a [**CHANGELOG.md**](CHANGELOG.md) generated 🚨

## ➤ Author

👤 **Lucca Pessoa**

Hey!! If you like this project or if you find some bugs feel free to contact me in my channels:

>
> * Email: luccapsm@gmail.com
> * Website: [lpmatos](https://github.com/lpmatos)
> * Github: [@lpmatos](https://github.com/lpmatos)
> * GitLab: [@lpmatos](https://gitlab.com/lpmatos)
> * LinkedIn: [@luccapessoa](https://www.linkedin.com/in/luccapessoa/)
>

## ➤ License 📑

🔖 Distributed under the **Apache License**. See [LICENSE](LICENSE) for more information.

## ➤ Troubleshooting

🚨 This is just a personal project created for study/demonstration purposes and to simplify my working life, it may or may not be a good fit for your project!

## ➤ Show your support

Give a ⭐️ if this project helped you!

---

This [README.md]() was generated with ❤️ by [me]()
