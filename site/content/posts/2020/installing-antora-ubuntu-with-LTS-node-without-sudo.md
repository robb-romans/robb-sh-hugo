---
title: "Installing Antora on Ubuntu with the latest LTS Node without Sudo"
date: 2020-02-05T11:59:39-06:00
draft: false
toc: false
featuredImg: "/images/posts/2020/02/pexels-photo-169573.jpeg"
slug: "install-antora-node-ubuntu-nosudo"
aliases:
  - "/posts/installing-antora-ubuntu-with-lts-node-without-sudo/"
tags: 
  - hugo
  - asciidoc
---

I imagine you're staring at your Ubuntu Linux [box, laptop, VM, Crostini
container] and you want to publish some sweet [Asciidoc](http://asciidoc.org/)
docs with [Antora](https://antora.org/). So, you hit the Google and land on
<https://docs.antora.org/antora/2.2/install/install-antora/> to find your next
move. Here's the process I followed to get a maintainable solution given some
personal constraints.

## Updating Node on Ubuntu

Aye, I'm on Ubuntu OS 18.04.

```bash
$ cat /etc/lsb-release | tail -n 1
DISTRIB_DESCRIPTION="Ubuntu 18.04.4 LTS"
```

Ubuntu 18.04 LTC has Node.js version 8 as of today. That seems ancient, and
Antora recommends the latest LTS Node. I prefer to use the distribution package
manager if available. In this case I don't want to install and use `nvm` or
whack Node directly into my system `/usr/bin` directory. The next best thing to
distribution packages are official [PPA](https://help.ubuntu.com/community/PPA)
packages.

### Use a PPA to get the latest LTC Node

Let's use a maintained PPA to get a more current LTS version of Node.

```bash
# Remove old versions installed with Ubuntu's package manager
$ sudo apt remove nodejs nodejs-dev npm

# Remove any other node packages
$ sudo apt autoremove

# WARNING: never just pipe something from the internet into sudo.
# Check the setup script carefully before you run the next command.
#
# Install LTS node 
$ curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -

# The big moment
$ sudo apt install nodejs

# Check that the versions are correct
$ node -v
v12.14.1

$ npm -v
6.13.4
```

### Set a local directory for Node modules

Installing node packages into `/usr/lib/node_modules` would require `sudo` and
creates unmaintained cruft in your file system. Antora docs suggest an
[alternative of installing its Node packages in a document
directory](https://docs.antora.org/antora/2.2/install/install-antora/#install-dir),
but this seems like it would have to be redone for each doc project directory,
and then you have to maintain it individually. Also, you have to prefix the path
every time you run antora.

Let's use a directory in our home directory and tell `npm` that this is now the
global directory for this user. We'll update our path to find these packages in
a bit.

First, let's set the package directory and update the `npm` config file.

```bash
# Create a home directory for node packages
$ export NPM_PACKAGES="$HOME/.npm-packages"

$ mkdir -p $NPM_PACKAGES

$ echo "prefix = $NPM_PACKAGES" >> ~/.npmrc
```

### Update your Bash .profile

I adapted some instructions that I found at
<https://makandracards.com/makandra/72209-how-to-install-npm-packages-globally-without-sudo-on-linux>.
Add this code to your `~/.profile`:

```bash
# NPM packages in homedir
NPM_PACKAGES="$HOME/.npm-packages"

# Tell our environment about user-installed node tools
if [ -d "$NPM_PACKAGES/bin" ] ; then
    PATH="$NPM_PACKAGES/bin:$PATH"
fi
# Unset manpath so we can inherit from /etc/manpath via the `manpath`
# command
if [ -d "$NPM_PACKAGES/share/man" ] ; then
    unset MANPATH  # delete this if you modified MANPATH elsewhere
                   # in your environment
    MANPATH="$NPM_PACKAGES/share/man:$(manpath)"
fi
# Tell Node about these packages
if [ -d "$NPM_PACKAGES/lib/node_modules" ] ; then
    NODE_PATH="$NPM_PACKAGES/lib/node_modules:$NODE_PATH"
fi
```

## Install Antora

Let's do it. `-g` is for global, y'all.

```bash
$ npm i -g @antora/cli@2.2 @antora/site-generator-default@2.2
[lots of output]
```

### Source your new Bash profile and run Antora

```bash
# Set environment variables and update your path
$ source ~/.profile

# Did it work? How are you feeling?
$ antora version
2.2.0
```

Enjoy.

---
Photo credit: Negative Space from Pexels
