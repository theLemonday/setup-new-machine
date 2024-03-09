#!/bin/bash

# $1: user to change shell
usermod --shell $(which zsh) $1
