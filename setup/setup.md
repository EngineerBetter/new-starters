# Setup

We use various tools/technologies for client engagements and beach work. As such we generally use the following applications/tools:

| Utilities | Formatting Tools | Database Tools | Languages | Git Tools | Container Tools | k8s | CF/BOSH/Concourse/Credhub | Terraform | Cloud IAAS | Security Tools | Apps |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| ripgrep | jq | postgresql | node | git-secrets | docker | kubernetes-cli | cf-cli | terraform-docs | awscli | openssl | brave-browser |
| tree | yj | sqlite | yarn | tig | skaffold | kustomize | bosh-cli | tflint | awslogs | gnupg | visual-studio-code |
| fd | yq | mysql-client | python | gh | - | kubectx | ironbird | tfsec | aws-iam-authenticator | pinentry-mac | slack |
| asdf | xmlstarlet | - | pyenv | - | - | k9s | control-tower | tfenv | aws-vault | ykman | zoom |
| bat | ytt | - | pipenv | - | - | krew | credhub-cli | - | assume-role | gpgme | discord |
| bash | - | - | ruby | - | - | ksd | opa | - | azure-cli | 1password-cli | - |
| bash-completion@2 | - | - | rbenv | - | - | kubectx | vault | - | google-cloud-sdk | - | - |
| zsh | - | - | rbspy | - | - | stern | safe | - | - | - | - |
| zsh-lovers | - | - | - | - | - | kops | - | - | - | - | - |
| zsh-syntax-highlighting | - | - | - | - | - | kind | - | - | - | - | - |
| gnutls | - | - | - | - | - | - | - | - | - | - | - |
| readline | - | - | - | - | - | - | - | - | - | - | - |
| wget | - | - | - | - | - | - | - | - | - | - | - |
| nmap | - | - | - | - | - | - | - | - | - | - | - |
| dog | - | - | - | - | - | - | - | - | - | - | - |
| telnet | - | - | - | - | - | - | - | - | - | - | - |
| coreutils | - | - | - | - | - | - | - | - | - | - | - |
| shellcheck | - | - | - | - | - | - | - | - | - | - | - |
| graphviz | - | - | - | - | - | - | - | - | - | - | - |
| p7zip | - | - | - | - | - | - | - | - | - | - | - |
| watch | - | - | - | - | - | - | - | - | - | - | - |
| tldr | - | - | - | - | - | - | - | - | - | - | - |
| vim | - | - | - | - | - | - | - | - | - | - | - |
| neofetch | - | - | - | - | - | - | - | - | - | - | - |
| neovim | - | - | - | - | - | - | - | - | - | - | - |
| gawk | - | - | - | - | - | - | - | - | - | - | - |
| gnu-sed | - | - | - | - | - | - | - | - | - | - | - |
| z | - | - | - | - | - | - | - | - | - | - | - |

You can always install the following using the [Brewfile](Brewfile) provided. For instructions on how to configure your system see below:

To Install using [Homebrew](https://brew.sh) see below (_although it's best to check if this command is up to date_):

```sh
$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
...
$ cd <path to workspace>/landing-zone/setup
$ brew update
$ brew bundle install --file=Brewfile
```
