FROM debian:bullseye

SHELL ["/bin/bash", "-c"]

ARG ASDF_VERSION=v0.10.2
ARG GO_VERSION=1.19.3
ARG NODE_VERSION=18.12.0
ARG SRC_MODULE=github.com/mkaiho/go-aws-sandbox

RUN apt-get update -y && apt-get install -y \
  curl gcc make git zip unzip jq openssl ca-certificates gnupg lsb-release groff

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
  && unzip awscliv2.zip \
  && rm awscliv2.zip \
  && ./aws/install \
  && rm -r ./aws \
  && curl -L -o direnv https://github.com/direnv/direnv/releases/download/v2.29.0/direnv.linux-amd64 \
  && chmod +x direnv \
  && mv direnv /usr/local/bin/ \
  && echo 'eval "$(direnv hook bash)"' >> ~/.bashrc
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg \
  && echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null \
  && apt-get update -y && apt-get -y install docker-ce docker-ce-cli containerd.io docker-compose-plugin

RUN mkdir -p /go/src/$SRC_MODULE && chmod -R 777 /go
RUN chsh -s /bin/bash

WORKDIR /go/src/$SRC_MODULE

RUN git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch ${ASDF_VERSION} \
  && echo '. ~/.asdf/asdf.sh' >> ~/.bash_profile \
  && echo '. ~/.asdf/completions/asdf.bash' >> ~/.bash_profile \
  && . ~/.asdf/asdf.sh \
  && . ~/.asdf/completions/asdf.bash \
  && asdf plugin-add golang https://github.com/kennyp/asdf-golang.git \
  && asdf plugin add nodejs \
  && asdf install golang ${GO_VERSION} \
  && asdf install nodejs ${NODE_VERSION}
