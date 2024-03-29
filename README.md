<h1 align="center" class="header">Chaker</h1>
<p align="center" class="desc">
  The Hacker News 'client' for the terminal. Written in Golang.
</p>

## Table of Content
- [Introduction](#introduction)
- [Features](#features)
- [Contributions and Issues](#contributions-and-issues)
- [Installation](#installation)
- [Usage](#usage)
- [Note](#note)
- [Credits](#credits)

## Introduction
<img src="https://files.catbox.moe/vba2h7.jpg" alt="Chaker screenshot">

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)

![GitHub Repo stars](https://img.shields.io/github/stars/HoangTuan110/chaker) [![CodeFactor](https://www.codefactor.io/repository/github/hoangtuan110/chaker/badge)](https://www.codefactor.io/repository/github/hoangtuan110/chaker) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](http://makeapullrequest.com) ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/HoangTuan110/chaker/Go)

Chaker (formerly Hecker) is a Hacker News 'client' for the terminal written in Golang.

(The client is in quote because technially this is more of a web scraper with a UI rather than an actual client)

## Features

* Easy to use - You just need to learn a few keybinds to use Chaker.
* Move between submissions easily with Up and Down arrow key.
* Open the submission's URL using `Enter`.
* View the comment section (or more correctly the submission itself) using `c`.
* Submission time similiar to HN. (Pleasse see [Note](#note))
* Move between different pages of submissions.
* Shows up submission's data (like upvote and so on) only when they are being pointed and fainted.
* Shows up the time on the top, and the page number and help at the bottom!

## Contributions

Contributions are welcome. If you found any bugs or want to requrest a feature, please open an issue.

## Installation

* First, install the Go compiler at [here](https://golang.org/dl/) or if you are using a Linux distro, install it using your distro's package manager.
* Next, if you have `$GOPATH` set up:
```sh
go get github.com/HoangTuan110/chaker
```
  And if you don't:
```sh
git clone https://github.com/HoangTuan110/chaker.git # Clone the repo
cd chaker
go mod init chaker # Initialize the package
go mod tidy
go build
```
* Finally, you will see the `chaker` binary in the cloned directory.

## Usage

Just run:

```sh
./chaker
```

## Note
* When start the program, you will need to wait for a few seconds for the program to scrape data.
* Depends on your Internet, it can take a pretty long time for the program to scrape data, and sometimes you will
  get `TLS handshake timeout` error. If you got one, make sure to wait for a few minutes and then run the program again.
  If that still doesn't fix the problem, then check your Internet.
* The submission time in Chaker, for most of the time, is correct, but not always (`45 hours ago`).
  If you have one, then unfortunately, I can't fix it. It is likely due to Go's Unix time thing is somewhat wrong.
  If you know how to fix this, then please open a new issue.

## Credits

Thanks [Charm](https://charm.sh) for their amazing CLI library `bubbletea` and `lipgloss`.

Thanks [README Templates](https://www.readme-templates.com) and [GitPoint's README](https://github.com/gitpoint/git-point#readme) for the template. This project uses GitPoint's README template.

Thanks [this TOC generator](https://ecotrust-canada.github.io/markdown-toc/) for the TOC (Table Of Content).
