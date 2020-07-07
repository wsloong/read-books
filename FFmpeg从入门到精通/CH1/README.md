## FFmpeg简介

### 1.3 FFmpeg 的墓本组成

`FFmpeg` 框架的基本组成包含 `AVFormat`、`AVCodec`、`AVFilter`、`AVDevice`、`AVUtil`模块库

* `AVFormat`: 封装和解封装，`MP4`,`FLY`,`KV`,`TS`等文件封装格式，`RTMP`,`RTSP`, `MMS`,`HLS`等网络协议封装格式
* `AVCodec`: 编码和解码，支持`MPEG4`, `AAC`, `MJPEG`等，还支持第三方的编码，如`H.264`(需要x264编码器)， `H.265`(需要x265编码器)，`MP3(mp3lame)`(需要libmp3lame编码器)
* `AVFilter`: 通用的音频、视频、字幕等滤镜处理框架，滤镜框架可以有多个输入和多个输出。

```
./ffmpeg -i INPUT -vf "split[main][tmp];[tmp]crop=iw:ih/2:0:0,vflip[flip];[main][flip]overlay=O:H/2" OUTPUT

命令将输入视频切割成两部分流， 一部分交给 crop 和 vflip 滤镜处理模块进行操作，另一部分保持原样，然后合并流到原来的 overlay 图层中，并显示在最上面一层，输出新的视频。

split 路径将分割后的视频流的第二部分打上标签[tmp],通过 crop 滤镜对该部分进行处理，然后进行纵坐标调整操作，打上新标签[flip]。然后将 [main] 和 [flip]标签进行合并

规则：
    1. 相同的 Filter 线性链之间用逗号隔开
    2. 不同的 Filter 线性链之间用分号隔开
    3. 线性链汇合时通过 [] 括起来的标签进行标示
```

* `ffplay`: `FFmpeg`的播放器
* `ffprobe`: `FFmpeg` 的多媒体分析器

### FFmpeg 安装

```
内容摘抄自 [Ubuntu-FFmpeg](http://trac.ffmpeg.org/wiki/CompilationGuide/Ubuntu);
有部分修改，比如一些库由于无法访问，替换了 github 上的仓库
```

1，在家目录下创建如下三个目录
    * `ffmpeg/ffmpeg_sources`: 存放源文件
    * `ffmpeg/ffmpeg_build`: 存放构建文件和安装库
    * `ffmpeg/bin`: 存放编译的二进制文件(ffmpeg,ffplay,ffprobe,x264,x265)
2， 安装依赖
```
sudo apt-get update -qq && sudo apt-get -y install \
  autoconf \
  automake \
  build-essential \
  cmake \
  git-core \
  libass-dev \
  libfreetype6-dev \
  libgnutls28-dev \
  libsdl2-dev \
  libtool \
  libva-dev \
  libvdpau-dev \
  libvorbis-dev \
  libxcb1-dev \
  libxcb-shm0-dev \
  libxcb-xfixes0-dev \
  pkg-config \
  texinfo \
  wget \
  yasm \
  zlib1g-dev

NOTE: 不使用 ffplay和x11grab； 可以省略依赖：libsdl2-dev libva-dev libvdpau-dev libxcb1-dev libxcb-shm0-dev libxcb-xfixes0-dev
```

3，编译安装

* NASM
```
如果仓库中 nasm 版本 >= 2.13: sudo apt-get install nasm

否则：
cd ~/ffmpeg/ffmpeg_sources && \
wget https://www.nasm.us/pub/nasm/releasebuilds/2.14.02/nasm-2.14.02.tar.bz2 && \
tar xjvf nasm-2.14.02.tar.bz2 && \
cd nasm-2.14.02 && \
./autogen.sh && \
PATH="$HOME/ffmpeg/bin:$PATH" ./configure --prefix="$HOME/ffmpeg/ffmpeg_build" --bindir="$HOME/ffmpeg/bin" && \
make && \
make install
```

* libx264

```
关联： --enable-gpl --enable-libx264

如果仓库中 libx264-dev 版本 >= 118: sudo apt-get install libx264-dev

否则：
cd ~/ffmpeg/ffmpeg_sources && \
git -C x264 pull 2> /dev/null || git clone --depth 1 https://code.videolan.org/videolan/x264.git && \
cd x264 && \
PATH="$HOME/ffmpeg/bin:$PATH" PKG_CONFIG_PATH="$HOME/ffmpeg/ffmpeg_build/lib/pkgconfig" ./configure --prefix="$HOME/ffmpeg/ffmpeg_build" --bindir="$HOME/ffmpeg/bin" --enable-static --enable-pic && \
PATH="$HOME/bin:$PATH" make && \
sudo make install
```

* lib265

```
关联： --enable-gpl --enable-libx265

如果仓库中 libx265-dev 版本 >= 68: sudo apt-get install libx265-dev libnuma-dev

否则：
sudo apt-get install mercurial libnuma-dev && \
cd ~/ffmpeg/ffmpeg_sources && \
hg clone https://bitbucket.org/multicoreware/x265 && \
cd x265/build/linux && \
PATH="$HOME/ffmpeg/bin:$PATH" cmake -G "Unix Makefiles" -DCMAKE_INSTALL_PREFIX="$HOME/ffmpeg/ffmpeg_build" -DENABLE_SHARED=off ../../source && \
PATH="$HOME/ffmpeg/bin:$PATH" make && \
sudo make install
```

* libvpx

```
关联： --enable-libvpx

如果仓库中 libvpx-dev 版本 >= 1.4.0: sudo apt-get install libvpx-dev

否则：
cd ~/ffmpeg/ffmpeg_sources && \
git -C libvpx pull 2> /dev/null || git clone --depth 1 git@github.com:webmproject/libvpx.git && \
cd libvpx && \
PATH="$HOME/ffmpeg/bin:$PATH" ./configure --prefix="$HOME/ffmpeg/ffmpeg_build" --disable-examples --disable-unit-tests --enable-vp9-highbitdepth --as=yasm && \
PATH="$HOME/ffmpeg/bin:$PATH" make && \
sudo make install
```

* libfdk-aac

```
关联： --enable-libfdk-aac(如果有`--enable-gpl`,也会关系到`--enable-nonfree`)

如果仓库中有 libvpx-dev: sudo apt-get install libfdk-aac-dev

否则：
cd ~/ffmpeg/ffmpeg_sources && \
git -C fdk-aac pull 2> /dev/null || git clone --depth 1 https://github.com/mstorsjo/fdk-aac && \
cd fdk-aac && \
autoreconf -fiv && \
./configure --prefix="$HOME/ffmpeg/ffmpeg_build" --disable-shared && \
make && \
sudo make install
```

* libmp3lame

```
关联： --enable-libmp3lame

如果仓库中 libmp3lame-dev 版本 >= 3.98.3: sudo apt-get install libmp3lame-dev

否则：
cd ~/ffmpeg/ffmpeg_sources && \
wget -O lame-3.100.tar.gz https://downloads.sourceforge.net/project/lame/lame/3.100/lame-3.100.tar.gz && \
tar xzvf lame-3.100.tar.gz && \
cd lame-3.100 && \
PATH="$HOME/ffmpeg/bin:$PATH" ./configure --prefix="$HOME/ffmpeg/ffmpeg_build" --bindir="$HOME/ffmpeg/bin" --disable-shared --enable-nasm && \
PATH="$HOME/ffmpeg/bin:$PATH" make && \
sudo make install
```

* libopus

```
关联： --enable-libopus

如果仓库中 libopus-dev 版本 >= 1.1: sudo apt-get install libopus-dev

否则：
cd ~/ffmpeg/ffmpeg_sources && \
git -C opus pull 2> /dev/null || git clone --depth 1 https://github.com/xiph/opus.git && \
cd opus && \
./autogen.sh && \
./configure --prefix="$HOME/ffmpeg/ffmpeg_build" --disable-shared && \
make && \
sudo make install
```

* libaom

```
cd ~/ffmpeg/ffmpeg_sources && \
git -C aom pull 2> /dev/null || git clone --depth 1 https://aomedia.googlesource.com/aom && \
mkdir -p aom_build && \
cd aom_build && \
PATH="$HOME/ffmpeg/bin:$PATH" cmake -G "Unix Makefiles" -DCMAKE_INSTALL_PREFIX="$HOME/ffmpeg/ffmpeg_build" -DENABLE_SHARED=off -DENABLE_NASM=on ../aom && \
PATH="$HOME/ffmpeg/bin:$PATH" make && \
sudo make install
```

* FFmpeg

```
cd ~/ffmpeg/ffmpeg_sources && \
wget -O ffmpeg-snapshot.tar.bz2 https://ffmpeg.org/releases/ffmpeg-snapshot.tar.bz2 && \
tar xjvf ffmpeg-snapshot.tar.bz2 && \
cd ffmpeg && \
PATH="$HOME/ffmpeg/bin:$PATH" PKG_CONFIG_PATH="$HOME/ffmpeg/ffmpeg_build/lib/pkgconfig" ./configure \
  --prefix="$HOME/ffmpeg/ffmpeg_build" \
  --pkg-config-flags="--static" \
  --extra-cflags="-I$HOME/ffmpeg/ffmpeg_build/include" \
  --extra-ldflags="-L$HOME/ffmpeg/ffmpeg_build/lib" \
  --extra-libs="-lpthread -lm" \
  --bindir="$HOME/ffmpeg/bin" \
  --enable-gpl \
  --enable-gnutls \
  --enable-libaom \
  --enable-libass \
  --enable-libfdk-aac \
  --enable-libfreetype \
  --enable-libmp3lame \
  --enable-libopus \
  --enable-libvorbis \
  --enable-libvpx \
  --enable-libx264 \
  --enable-libx265 \
  --enable-nonfree && \
PATH="$HOME/ffmpeg/bin:$PATH" make && \
sudo make install && \
hash -r
```

* 应用环境变量
    `source ~/.profile`
* 更新
    * rm -rf ~/ffmpeg_build ~/ffmpeg_sources ~/bin/{ffmpeg,ffprobe,ffplay,x264,x265,nasm}
    * hash -r
    * 删除之前的安装包`sudo apt-get autoremove autoconf automake build-essential cmake git-core libass-dev libfreetype6-dev libgnutls28-dev libmp3lame-dev libnuma-dev libopus-dev libsdl2-dev libtool libva-dev libvdpau-dev libvorbis-dev libvpx-dev libx264-dev libx265-dev libxcb1-dev libxcb-shm0-dev ibxcb-xfixes0-dev mercurial texinfo wget yasm zlib1g-dev`

执行以上过程之后的问题:

    1. `gnutls not found using pkg-config`： `gnuts`是一个`C语言`加密库，不需要去掉即可，也可以单独安装