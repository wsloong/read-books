## FFmpeg 基本使用

### 2.1 ffmpeg 常用命令
* `ffmpeg --help [long/full]`: 查看帮助
* `ffmpeg -formats`: 查看支持视频文件格式
* `ffmpeg -codecs`: 查看支持编解码格式
    * `ffmpeg -encoders`: 查看编码格式
    * `ffmpeg -decoders`: 查看解码格式
```
输出信息包含三列
1. 第一列包含6个字段
    a. 音频、视频、字幕
    b. 帧级别的多线程支持
    c. 分片级别的多线程
    d. 是否为测试版本
    e. draw horiz hand模式支持
    f. 直接渲染模式支持
2. 编解码格式
3.说明
```
* `ffmpeg -muxers`: 查看封装格式
* `ffmpeg -demuxers`: 查看解封装格式
* `ffmpeg -filters`: 查看支持滤镜
```
输出信息包含四列
1. 第一列有3个字段
    a. 时间轴支持
    b. 分片线程处理支持
    c. 命令支持
2. 滤镜名
3. 转换方式，如音频转音频，视频转视频，创建音频，创建视频等
4. 说明
```
* `ffmpeg -h muxer=flv`: 查看 `FLV` 封装器的参数支持
* `ffmpeg -h demuxer=flv`: 查看 `FLV` 解封装参数支持
* `ffmpeg -h encoder=h264`: 查看 `H.264(AVC)` 编码参数支持
* `ffmpeg -h decoder=h264`: 查看 `H.264(AVC)` 解码参数支持
* `ffmpeg -h filter=colorkey`: 查看 `colorkey` 滤镜的参数支持

#### 2.1.1 ffmpeg 的封装转换参数

转封装功能通过 `AVFormat` 来完成，通过 `libavformat` 库进行 `Mux` 和 `Demux` 操作

|参数|类型|说明|
|---|---|---|
|avioflags|标记|`format`的缓冲设置，默认0，有缓冲|视频id|
||direct|无缓冲状态|
|probesize|整数|在进行媒体数据处理前获得文件内容的大小，可用在预读取文件头时提高速度，也可以设置足够大的值来读取到足够多的音视频数据信息|
||标记||
||flush_packets|立即将`packets`数据刷新写入文件中|
||genpts|输出时按照正常规则产生`pts`|
||nofillin|不填写可以精确计算缺失的值|
||igndts| 忽略`dts`|
||discardcorrupt|丢弃损坏的帧|
|fflags|sortdts|尝试以`dts`的顺序为准输出|
||keepside|不合并数据|
||fastseek|快速`seek`（定位）操作，但是不够精确|
||latm|设置`RTP MP4 LATM`生效|
||nobuffer|直接读取或写出，不存入`buffer`，用于在直播采集时降低延迟|
||bitexact|不写入随机或者不稳定的数据|
|seek2any|整数|支持随意位置`seek`，这个`seek`不以`keyframe`为参考|
|analyzeduration|整数|指定解析媒体所需要花销的时间，这里设置的值越高，解析越准确，如果在直播中为了降低延迟，这个值可以设置得低一些|
|codec_whitelist|列表|设置可以解析的`codec`白名单|
|format_whitelist|列表|设置可以解析的`format`自名单|
|output_ts_offset|整数|设置输出文件的起始时间|

#### 2.1.2 ffmpeg 的转码参数

编解码功能通过 `AVCodec` 来完成，通过 `libavcodec` 库进行 `Encode` 和 `Decode` 操作

|参数|类型|说明|
|---|---|---|
|b|整数|设置音频与视频码率，可以认为是音视频加起来的码率，默认为`200kbit/s`；使用这个参数可以根据`b:v`设置视频码率，`b:a`设置音频码|
|ab|整数|设置音频的码率，默认是`128kbit/s`|
|g|整数|设置视频`GOP`（可以理解为关键帧间隔）大小，默认是`12`帧一个`GOP`|
|ar|整数|设置音频采样率，默认为`0`|
|ac|整数|设置音频通道数，默认为`0`|
|bf|整数|设置连续编码为`B`帧的个数，默认为`0`|
|maxrate|整数|最大码率设置，与`bufsize`一同使用即可，默认为`0`|
|minrate|整数|最小码率设置，配合`maxrate`和`bufsize`可以设置为`CBR`模式，平时很少用，默认为`0`|
|bufsize|整数|设置控制码率的`buffer`的大小，默认为`0`|
|keyint_min|整数|设置关键帧最小间隔，默认为`25`|
|sc_threshold|整数|设置场景切换支持，默认为`0`|
|me_threshold|整数|设置运动估计阂值，默认为``0|
|mb_threshold|整数|设置宏块阂值，默认为`0`|
|profile|整数|设置音视频的`profile`，默认为`－99`|
|level|整数|设置音视频的`level`，默认为`99`|
|timecode_frame_start|整数|设置`GOP`帧的开始时间，需要在`non-drop-frame`默认情况下使用|
|channel_layout|整数|设置音频通道的布局格式|
|threads|整数|设置编解码工作的线程数|

#### 2.1.3 ffmpeg 的基本转码原理
示例命令： `./ffmpeg -i ~/Movies/inputl.rmvb -vcodec mpeg4 -b:v 200k -r 15 -an output.mp4`
```
从输出信息看:
1. 封装格式从 RMVB 转换为 MP4
2. 视频编码从 RV40 转换为 MPEG4
3. 视频码率从 277kbit/s 转换为 200kbit/s
4. 视频帧率从 23.98fps 转换为 15fps
5. 转码后的文件不包含音频(-an 参数)

转码流程 解封装(解RMVB) => 解码(视频编码为RV40,音频为COOK) => 编码(编为MP4) => 封装(没有音频的MP4)
```

### 2.2 ffprobe 常用命令

* `ffprobe --help`: 查看帮助
* `ffprobe -show_packets input.flv`: 查看数据包信息
* `ffprobe -show_data -show_packets input.flv`: 组合参数查看包中具体数据
* `ffprobe -show_format input.mp4`: 查看封装格式
* `ffprobe -show_frames input.flv`: 查看帧信息
* `ffprobe -show_streams input.flv`: 查看流信息
* `ffprobe -print_format` 和 `ffprobe -of`: 修改输出格式，格式可以是`xml/ini/json/csv/flat`等

### 2.3 ffplay 常用命令

略，不常用