## FFmpeg 的转封装

### 转 MP4 格式

#### 3.1.2 MP4 分析工具

1. Elecard StreamEye
2. mp4box
3. mp4info

#### MP4 在 FFmpeg 中的 Demuxer

`ffmpeg -h demuxer=mp4` 可以看到`MP4`和`mov、3gp、m4a、3g2、mj2`的`Demuxer`相同。

|参数| 类型|说明|
|---|---|---|
|use_absolute_path|布尔|可以通过绝对路径加载外部的`tra`，可能会有安全因素的影响，默认不开启|
|seek_streams_individually|布尔|根据单独流进行`seek`，默认开启|
|ignore_editlist|布尔|忽略`EditList Atom`信息，默认不开启|
|ignore_chapters|布尔|忽略`hapters`信息，默认不开启|
|enable_drefs|布尔|外部`track`支持，默认不开启|

#### MP4 在 FFmpeg 中的 Muxer

`MP4` 与 `mov、3gp、m4a 3g2 mj2`的`Demuxer`相同，它们的`Muxer`也差别不大，但是是不同的`Muxer`，

|参数| 类型|说明|
|---|---|---|
|||`MP4 Muxer` 标记|
||rtphint|增加`RTP`的`hint track`|
||empty_moov| 初始化空的`moov box`|
||frag_keyframe|在视频关键帧处切片|
||separate_moof|每一个`Track`写独立的`moof/mdat box` |
||frag_custom|每一个`caller`请求时`Flush`个片段|
||isml|创建实时流媒体（创建一个直播流发布点）|
||fasts_tart|将 `moov box` 移动到文件的头部|
|movflags|omit_tfhd_offset| 忽略 `tfhd` 容器中的基础数据偏移|
||disable_chpl| 关闭 `Nero Chapter` 容器|
||default_base_moof|在 `tfhd` 容器中设置 `default-base-is-moof` 标记|
||dash| 兼容 `DASH` 格式的 `mp4` 分片|
||frag_discont|分片不连续式设置 `discontinuous` 信号|
||delay_moov|延迟写入 `moov` 信息，直到第一个分片切出来，或者第一片被刷|
||global_sidx|在文件的开头设置公共的 `sidx` 索引|
||write_colr|写人 `colr` 容器|
||wnte_gama|写被弃用的 `gama` 容器|
|mioov_size|正整数|设置 `moov` 容器大小的最大值|
|||设置 `rtp` 传输相关的标记|
||latm|使用 `MP4A-LATM` 方式传输 `AAC` 音频|
||rfc2190|使用 `RFC2190` 传输 `H.264`, `H.263` |
|rtpflags|skip_rtcp|忽略使用 `RTCP`|
||h264_modeO|使用 `RTP` 传输 `modeO` 的 `H264`|
||send_bye|当传输结束时发送 `RTCP` 的 `BYE` 包|
|skip_iods|布尔型|不写入 `iods` 容器|
|iods_audio_profile|0～255|设置 `iods` 的音频 `profile` 容器|
|iods_video_profile|0～255|设置 `iods` 的视频 `profile` 容器|
|frag_duration|正整数|切片最大的 `duration`|
|min_frag_duration|正整数|切片最小的 `duration`|
|frag_size|正整数|切片最大的大小|
|ism_lookahead|正整数|预读取 `ISM` 文件的数量|
|video_track_timescale|正整数|设置所有视频的时间计算方式|
|brand|字符串|写 `major brand` |
|use_editlist|布尔型|使用 `edit list` |
|fragment_index|正整数|下一个分片编号|
|mov_gamma|0~10|`Gama` 容器的 `gama` 值|
|frag_interleave|正整数|交错分片样本|
|encryption_scheme|字符串|配置加密的方案|
|encryption_key|二进制|秘钥|
|encryption_kid|二进制|秘钥标识符|

* 示例1: faststart 
在互联网的视频点播中，如果希望 MP4 文件被快速打开，需要将 `moov` 存放 `mdat` 的前面；
如果放后面，需要将 `MP4` 文件下载完成后才可以进行播放。可以通过 `faststart`实现:
`./ffmpeg -i input.flv -c copy -f mp4 -movflags faststart output.mp4`

* 示例2: dash
`./ffmpeg -i input.flv -c copy -f mp4 -movflags dash output.mp4`

* 示例3: isml
`ISMV`是微软发布的一个流媒体格式，通过参数`isml`可以发布`ISML`直播流，将`ISMV`推流到`IIS`服务器。
`./ffmpeg -re -i input.mp4 -c copy -movflags isml+frag_keyframe -f ismv Stream`
生成的文件格式原理类似于`HLS`,使用`XML`格式进行索引。

### 转 FLV

#### 3.2.2 转 FLV 参数

`ffmpeg -h muxer=flv`

|参数| 类型|说明|
|---|---|---|
||flag|设置生成 `FLV` 时使用的 `flag`|
||aac_seq_header_detect|添加 `AAC` 音频的 `Sequence Header`|
|flvflags|no_sequence_end|生成 `FLV` 结束时不写入 `Sequence End` |
||no_metadata|生成 `FLV` 时不写人 `metadata`|
||no_duration_filesize|用于直播时不在 `metadata` 中写入 `duration` 与 `filesize`|
||add_ keyframe _index| 生成 `FLV` 时自动写入关键帧索引信息到 `metadata` 头|

根据表中的参数可以看出，在生成 `FLV` 文件时，写人视频、音频数据时均需要 `Sequence Header` 数据，
如果 `FLV` 的视频流中没有 `Sequence Header`，那么视频很有可能不会显示出来；
如果 `FLV` 的音频流中没有 `Sequence Header`，那么音频很有可能不会被播放出来。
所以需要将 `ffmpeg` 中的参数 `flvflags` 的值设置为 `aac_seq_header_detect`，其将会写入音频 `AAC Sequence Header`

#### 3.2.3 FFmpeg 文件转 FLV 举例

1. `FLV` 封装中支持的视频编码主要包含如下内容
    * Sorenson H.263 
    * Screen Video 
    * On2 VP6 
    * 带 Alpha 通道的 On2 VP6 
    * Screen Video 2 
    * H.264 (AVC) 

2. `FLV` 封装中支持的音频主要包含如下内容
    * 线性`PCM`，大小端取决于平台
    * ADPCM 音频格式
    * 线性`PCM`，小端
    * Nellymoser 16kHz Mono 
    * Nellymoser 8kHz Mono 
    * Nellymoser 
    * G.711 A-law 
    * G.711 mu-law 
    * Speex 
    * MP3 
    * AAC 
    * MP3 8kHz

如果封装 `FLV`时候，内部的音频或视频不符合标准时候，会报错；
将`AC3`封装进`FLV`的报错示例：
```
ffmpeg -i intput_ac3.mp4 -c copy -f flv output.flv
```

会出现 `[flv @ Ox7fe624809200] Audio codec ac3 not compatible with flv `的错误。

将 `AC3` 转换为 `AAC` 或者 `MP3`则可以避免出错: 
```
ffmpeg -i input ac3.mp4 -vcodec copy -acodec aac -f flv output.flv
```

#### 3.2.4 FFmpeg 生成带关键字索引的 FLV

使用 `add_keyframe_index` 将`FLV`文件中的关键帧建立一个索引，并写入`Metadata`头中

```
ffmpeg -i input.mp4 -c copy -f flv -flvflags add_keyframe_index output.flv
```

#### 3.2.5 FLV 文件格式分析工具
* `flvparse`
* `FlvAnalyzer`
* `ffprobe`: `ffprobe -v trace -i output.flv`

### 转 M3U8

M3U8文件的标签

|标签|说明|
|---|---|
|EXTM3U|必须且必须文件的第一行|
|EXT-X-VERSION |常见的是`3`|
|EXT-X-TARGETDURATION|最大分片浮点数四舍五入的整数值|
|EXT-X-MEDIA-SEQUENCE|`M3U8`直播切片序列，当打开`M3U8` 时，以这个标签的值为参考，播放对应的序列号的切片。分片必须是动态的，序列不能相同、必须是增序的。当`M3U8`列表中没有出现`EXT-X-ENDLIST`，播放分片都是从倒数第三片开始，不满足三片则不应该播放(定制可以不遵守)；`EXT-X-DISCONTINUITY`解决前后2个分片不连续的错误|
|EXTINF|每一分片的duration，`EXTINF`下面的信息为具体分片信息，分片路径可以是`相对路径`、`绝对路径`、`网络的URL链接地址`|
|EXT-X-ENDLIST|表明该 `M3U8` 文件不会再产生更多的切片|
|EXT-X-STREAM-INF|主要是出现在多级 `M3U8` 文件中：`M3U8` 中包含子 `M3U8`列表|

#### 3.3.2 FFmpeg 转 HLS 参数

|参数| 类型|说明|
|---|---|---|
|start_number|整数|设置 `M3U8` 列表中的第一片的序列数|
|hls_time|浮点数|设置每一片时长|
|hls_list_size|整数|设置 `M3U8` 中分片的个数|
|hls_ts_options|字符串|设置 `TS` 切片的参数|
|hls_wrap|整数|设置切片索引回滚的边界值|
|hls_allow_cache|整数|设置 `M3U8` 中 `EXT-X-ALLOW-CACHE` 的标签|
|hls_base_url|字符串|设置 `M3U8` 中每一片的前置路径|
|hls_segment_filename|字符串|设置切片名模板|
|hls_key_info_file|字符串|设置 `M3U8` 加密的 `key` 文件路径|
|hls_subtitle_path|字符串|设置 `M3U8` 字幕路径|
|hls_flags|标签(整数)|设置 `M3U8` 文件列表的操作。具体如下：`single_file`：生成一个媒体文件索引与字节范围; `delete_segments`：删除 M3U8 文件中不包含的过期的`TS`切片文件; `round_durations`: 生成的`M3U8`切片信息的`duration`为整数; `discont_start`：生成`M3U8`的时候在列表前边加上`discontinuity`标签; `omit_endlist`：在`M3U8`末尾不追加`endlist 标签|
|use_localtime|布尔| 设置`M3U8`文件序号为本地时间戳|
|use_localtime_mkdir|布尔|根据本地时间戳生成目录|
|hls_playlist_type|整数|设置`M3U8`列表为事件或者点播列表|
|method|字符串|设置 `HTTP` 属性|

#### 3.3.3 示例

文件转换 `HLS` 直播
```
ffmpeg -re -i input.mp4 -c copy -f hls -bsf:v h264_mp4toannexb output.m3u8
```

参数`-bsf:v h264_mp4toannexb`将 `MP4` 中的 `H.264` 转换为 `H.264 AnnexB`标准编码，`AnnexB`常用于实时传输流中。
如果源文件为`FLV`、`TS`等可作为直播传输流的视频，则不需要这个参数。

1. start_number  参数
设置`M3U8`列表中的`第一片的序列数`，下面命令将第一片的序列数设置为`300`：
```
ffmpeg -re -i input.mp4 -c copy -f hls -bsf:v h264_mp4toannexb -start_number 300 output.m3u8
```

2. hls_time 参数
设置`M3U8`列表中切片的`duration`，下面命令控制切片为10秒钟一个(从关键帧处开始切片，所以时间并不是很均匀，如果先转码在切片则会比较均匀)
```
ffmpeg -re -i input.mp4 -c copy -f hls -bsf:v h264_mp4toannexb -hls_time 10 output.m3u8
```

3. hls_list_size 参数
设置`M3U8`列表中切片的`个数`，下面命令控制切片为3个
```
ffmpeg -re -i input.mp4 -c copy -f hls -bsf:v h264_mp4toannexb -hls_list_size 3 output.m3u8
```

4. hls_wrap 参数
为`M3U8`列表中`TS序号`设置刷新回滚参数，当`TS`分片序号等于`hls_wrap`参数设置的数值时候回滚;
下面的命令，当 `TS`序号等与`3`时候，将其设置为`0`；

*对`CDN`缓存节点的支持并不友好，未来可能会被弃用*

```
ffmpeg -re -i input.mp4 -c copy -f hls -bsf:v h264_mp4toannexb -hls_wrap 3 output.m3u8

输出：
#EXTMJU 
#EXT-X-VERSION:J 
#EXT-X-TARGETDURATION:7 
#EXT-X-MEDIA-SEQUENCE:62 
#EXTINF:S.000000, 
output2.ts 
#EXTINF:6.960000, 
outputO.ts 
#EXTINF:J.200000, 
output1.ts 
#EXTINF:J.840000, 
output2.ts 
#EXTINF:0.960000, 
outputO.ts
```

5. hls_base_url 参数
为`M3U8`列表中的分片设置前置基本路径，下面命令为其设置网络路径，也可以是本地绝对路径或者相对路径
```
ffmpeg -re -i input.mp4 -c copy -f hls -hls_base_url http://192.168.0.1/live/ -bsf:v h264_mp4toannexb output.m3u8
```

6. hls_segment_filename 参数
为`M3U8`列表设置切片文件名的规则模板参数。下面命令设置切片名称为`test_output-1.ts`, `test_output-2.ts` ....
```
ffmpeg -re -i input.mp4 -c copy -f hls -hls_segment_filename test_output-%d.ts -bsf:v h264_mp4toannexb output.m3u8
```

7. hls_flags 参数
    包含一些子参数，如下：

    * delete_segments: 删除已经不在`M3U8`列表中的旧文件(`hls_list_size`大小2倍作为删除依据)
    ```
    ffmpeg -re -i input.mp4 -c copy -f hls -hls_flags delete_segments -hls_list_size 4 -bsf:v h264_mp4toannexb output.m3u8
    ```

    * round_duration: 切片信息的`duration`为整数
    ```
    ffmpeg -re -i input.mp4 -c copy -f hls -hls_flags round_duration -bsf:v h264_mp4toannexb output.m3u8
    ```

    * discont_start: 在第一片切片前，插入 `discontinuity`标签
    ```
    ffmpeg -re -i input.mp4 -c copy -f hls -hls_flags discont_start -bsf:v h264_mp4toannexb output.m3u8
    ```

    * omit_endlist: 在生成`M3U8`结束的时候，若不在文件末尾则不追加`endlist`标签。常规生成的`M3U8`文件结束时，`FFmpeg`会默认写入`endlist`标签
    ```
    ffmpeg -re -i input.mp4 -c copy -f hls -hls_flags omit_endlist -bsf:v h264_mp4toannexb output.m3u8
    ```

    * split_by_time: 必须配合`hls_time`参数，根据`hls_time`参数设置的值作为秒数参考对`TS`进行切片，不一定遇到关键帧;所以有可能会影响首画面体检，例如花屏或者首画面显示慢的问题，因为第一帧不一定是关键帧
    ```
    ffmpeg -re -i input.ts -c copy -f hls -hls_time 2 -hls_flags split_by_time output.m3u8
    ```

8. use_localtime 参数
以本地系统时间为切片文件名
```
ffmpeg -re -i input.mp4 -c copy -f hls -use_localtime 1 -bsf:v h264_mp4toannexb output.m3u8
```

9. method 参数
设置`HLS`将`M3U8`以及`TS`文件上传到`HTTP`服务器。需要一台`HTTP`服务器，支持上传的相关方法，比如`PUT、POST`等；
可以尝试使用`Nginx`的`webdav`模块来完成这个功能；
* Nginx配置如下:
```
location / { 
    client_max_body_size lOM; 
    dav_access      group:rw all:rw; 
    dav methods PUT DELETE MKCOL COPY MOVE; 
    root html/; 
}
```
* FFmpeg命令如下: 上传到`Nginx`配置目录下的`/test/`目录下
```
ffmpeg -re -i input.mp4 -c copy -f hls -hls_time 3 -hls_list_size 0 -method PUT -t 30 http://127.0.0.1/test/output_test.m3u8
```

### 3.4 视频文件切片

视频文件切片与`HLS`基本类似，但是`HLS`切片在标准中只支持`TS`格式的切片，并且是直播与点播切片。

#### 3.4.1 FFmpeg 切片 segment 参数

|参数|类型|说明|
|---|---|---|
||||
||||
||||
||||
||||
||||
||||
||||