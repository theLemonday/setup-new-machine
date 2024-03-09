#!/bin/bash

sudo apt-get install fd-find

ln -s $(which fdfind) /home/$1/.local/bin/fd
