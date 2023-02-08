# Basalt

Basalt is a command-line tool for managing user access to AWS infrastructure through bastion, distributing and rotating
SSH keys.

## Features

- Create IAM users and associate them with SSH keys for access to bastion host and AWS infrastructure through it
- Secure access to infrastructure hosts through bastion and distributing SSH keys to them
- Rotate SSH keys in AWS Secrets Manager and on infrastructure hosts
- Remove IAM users and associated SSH keys from AWS Secrets Manager and infrastructure hosts

## Installation

Download the latest binary archive suitable to your operating system
from [GitHub releases](https://github.com/andrewmolyuk/basalt/releases).

Extract the suitable to your operating system archive to any preferable place, and you are ready to run the binary. For
more comfortable usage you can move binary to `/usr/local/bin` on your MacOS or `/usr/bin` on your Linux system. For
Windows, you can move it to `C:\Program Files\Basalt` and add the folder to the `%PATH%` environment variable.

## Usage

Get the list of available options with `basalt --help` command.

```shell
basalt --help
```
