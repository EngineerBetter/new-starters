# Workstation Setup

## Workspace

Most folks clone repositories into `$HOME/workspace/`. You don't _have_ to follow this convention, but it's worth being aware of.

## Tools

We use various tools/technologies for client engagements and beach work. You may find it useful to install everything in advance; alternatively you could instead install them in a just-in-time fashion as-and-when they are needed.

[This Brewfile](Brewfile) lists commonly-used tools and gives a brief explanation of each, and can also be used with [Homebrew](https://brew.sh) to install all of them:

```sh
# Install Homebrew package manager first
$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
# Update brew definitions
$ brew update
# Install all the things in the Brewfile
$ brew bundle install --file <(https://raw.githubusercontent.com/EngineerBetter/new-starters/main/setup/Brewfile)
```
