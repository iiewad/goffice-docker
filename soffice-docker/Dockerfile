#Set the sys version
FROM ubuntu:16.04

ENV DEBIAN_FRONTEND noninteractive

## Create a '$UNAME' user
## --build-arg UID=$(id -u) --build-arg GID=$(id -g)
ARG UNAME=docker
ARG UID=1000
ARG GID=1000

RUN addgroup --gid $GID $UNAME \
 && adduser --uid $UID --gid $GID --disabled-password --gecos "docker User" $UNAME \
 && usermod -L $UNAME


RUN apt-get clean           \
 && rm -rf /var/lib/apt/lists/* \
 && apt-get update          \
 && apt-get install -y      \
    ghostscript             \
    libxinerama1            \
    libdbus-glib-1-2        \
    libcairo2               \
    libcups2                \
    libgl1-mesa-dri         \
    libgl1-mesa-glx         \
    libsm6                  \
    fonts-opensymbol        \
    hyphen-fr               \
    hyphen-de               \
    hyphen-en-us            \
    hyphen-it               \
    hyphen-ru               \
    fonts-dejavu            \
    fonts-dejavu-core       \
    fonts-dejavu-extra      \
    fonts-dustin            \
    fonts-f500              \
    fonts-fanwood           \
    fonts-freefont-ttf      \
    fonts-liberation        \
    fonts-lmodern           \
    fonts-lyx               \
    fonts-sil-gentium       \
    fonts-texgyre           \
    fonts-tlwg-purisa       \
    curl                    \
    default-jre		    \
    --no-install-recommends \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN mkdir /tmp/lo \
  && cd /tmp/lo \
  && curl -L -k http://mirrors.ustc.edu.cn/tdf/libreoffice/stable/6.1.5/deb/x86_64/LibreOffice_6.1.5_Linux_x86-64_deb.tar.gz | tar xz --strip-components=1 \
  && cd DEBS \
  && dpkg -i *.deb \
  && rm -fr /tmp/lo

# for compat with canvadocs converter, which hardcodes path to soffice
RUN ln -s /opt/libreoffice6.1/program/soffice /usr/bin/soffice
ENV PATH $PATH:/opt/libreoffice6.1/program

USER $UNAME
RUN cd /home/$UNAME && mkdir ocr_storage

# install windows fonts
RUN soffice --headless --convert-to pdf *.docx
ADD fonts.tar.gz /home/$UNAME/.config/libreoffice/4/user/
