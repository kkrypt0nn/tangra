# tangra

[![Discord Server Badge](https://img.shields.io/discord/1358456011316396295?logo=discord)](https://discord.gg/xj6y5ZaTMr) [![Go Reference](https://pkg.go.dev/badge/github.com/kkrypt0nn/tangra.svg)](https://pkg.go.dev/github.com/kkrypt0nn/tangra) ![Repository License](https://img.shields.io/github/license/kkrypt0nn/tangra?style=flat-square) ![Code Size](https://img.shields.io/github/languages/code-size/kkrypt0nn/tangra?style=flat-square) ![Last Commit](https://img.shields.io/github/last-commit/kkrypt0nn/tangra?style=flat-square)

A lightweight and easy to use Go logging library that includes logging functionalities with different levels and custom
formatting. It can also be used as a library to simply format the various messages you print in the terminal.

## Showcase

### macOS

![macOS Showcase](assets/macos.png)

### Linux

![Linux Showcase](assets/linux.png)

### Windows

![Windows Showcase](assets/windows.png)

## Installation

If you want to use this library for one of your projects, you can install it like any other Go library

```shell
go get github.com/kkrypt0nn/tangra
```

## Customizing

### Prefix

The prefix, what comes before the message, can be changed with the `SetPrefix` method on a `Logger` structure.

> The default prefix is `${datetime} ${level:color}${level:name}${reset}: `

### Placeholders

There are formatting placeholders that will be replaced in both the message and the prefix that can be
seen [here](PLACEHOLDERS.md). All the placeholders in effect can be seen in
the [`placeholders example`](examples/placeholders/main.go).

[For example](examples), logging the following message

```
${fg:red}${effect:blink}${effect:bold}${sys:username} says hello!
```

Will print a red blinking message in bold that says `<username> says hello!`, where `<username>` is the username on your
system.

### Styling

You can choose whether you want to style your messages or not with the `SetStyling` method on a `Logger` structure.
Styling includes foreground colors, background colors and special effects such as bold, and others - see
the [terminal package](terminal).

> **Note**: The styling will **not** apply to the message if it is not supported by the terminal.

### Log File

Logs can also be written inside a log file with styling removed. See the [example here](examples/file/main.go).

## License

This library was made with 💜 by Krypton and is under the [MIT](LICENSE.md) license.
