FROM ubuntu

RUN apt update -y
RUN apt upgrade -y

RUN apt-get install software-properties-common -y
RUN add-apt-repository ppa:neovim-ppa/unstable -y
RUN apt-get update -y
RUN apt-get install -y --no-install-recommends neovim fish

# CMD /usr/bin/fish
CMD ["tail", "-f", "/dev/null"]
