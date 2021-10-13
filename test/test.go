package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	logger "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var ctx = context.Background()

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "10.10.88.161:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	logger.SetReportCaller(true)
	logger.SetFormatter(&logger.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006/01/02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (fn string, file string) {
			filename := path.Base(f.File)
			funcname := path.Base(f.Function)
			return fmt.Sprintf("%s():", funcname),
				fmt.Sprintf(" %s:%d,", filename, f.Line)
		},
	})

	logger.SetLevel(logger.InfoLevel)
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		panic("eeeee")
	}
	//logger.SetOutput(os.Stdout)
	logger.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logger.Info("info msg")

}

//获取离线下载记录
type ListTaskResponse struct {
	Result          int32       `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Message         string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Time            float64     `protobuf:"fixed64,3,opt,name=time,proto3" json:"time,omitempty"`
	Total           float64     `protobuf:"fixed64,4,opt,name=total,proto3" json:"total,omitempty"`
	CollectionTotal float64     `protobuf:"fixed64,5,opt,name=collection_total,json=collectionTotal,proto3" json:"collection_total,omitempty"`
	TaskInfo        []*TaskInfo `protobuf:"bytes,6,rep,name=task_info,json=taskInfo,proto3" json:"task_info,omitempty"`
	Config          *TaskConfig `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty"`
	RecordTotal     float64     `protobuf:"fixed64,8,opt,name=record_total,json=recordTotal,proto3" json:"record_total,omitempty"`
}

type TaskInfo struct {
	Filetype int32 `protobuf:"varint,1,opt,name=filetype,proto3" json:"filetype,omitempty"`
	//create_time
	CreateTime string `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// url_type
	UrlType int32 `protobuf:"varint,3,opt,name=url_type,json=urlType,proto3" json:"url_type,omitempty"`
	// filesize
	Filesize float64 `protobuf:"fixed64,4,opt,name=filesize,proto3" json:"filesize,omitempty"`
	//filename 文件名称
	Filename string `protobuf:"bytes,5,opt,name=filename,proto3" json:"filename,omitempty"`
	// url
	Url string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	// task_id
	TaskId float64 `protobuf:"fixed64,7,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	//collection_id
	CollectionId float64 `protobuf:"fixed64,8,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	RecordType   int32   `protobuf:"varint,9,opt,name=record_type,json=recordType,proto3" json:"record_type,omitempty"`
	RemarkName   string  `protobuf:"bytes,10,opt,name=remark_name,json=remarkName,proto3" json:"remark_name,omitempty"`
	TaskIdStr    string  `protobuf:"bytes,11,opt,name=task_id_str,json=taskIdStr,proto3" json:"task_id_str,omitempty"`
}

type TaskConfig struct {
	// collection_num_limit 服务器的收藏配置，json格式
	CollectionNumLimit map[string]int32   `protobuf:"bytes,1,rep,name=collection_num_limit,json=collectionNumLimit,proto3" json:"collection_num_limit,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	SaveTime           map[string]float64 `protobuf:"bytes,2,rep,name=save_time,json=saveTime,proto3" json:"save_time,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	ValidTimeLimit     map[string]float64 `protobuf:"bytes,3,rep,name=valid_time_limit,json=validTimeLimit,proto3" json:"valid_time_limit,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

type Replay struct {
	Path   string `json:"path"`
	Accept string `json:"accept"`
	Auth   string `json:"auth"`
	UserID int64  `json:"user_id"`
}

var ch = make(chan Replay, 1000)

func test(replayData Replay, i int) {
	oldRespData, oldRespBody, oldUrl, err := NewReq("http://api-zone-lixian-vip-ssl.xunlei.com/collections", replayData)
	if len(oldRespBody) > 170 {
		log.Println(replayData)
	}
	if err != nil {
		logger.Println(replayData)
	}


	logger.Println(string(oldRespBody))
	logger.Println(oldRespData)
	logger.Println(oldUrl)

	//newRespData, newRespBody, newUrl, err := NewReq("http://xllixian.record.k8s.xunlei.cn/tasks", replayData)
	//if err != nil {
	//	log.Println(replayData)
	//}
	//
	//if newRespData == nil || oldRespData == nil {
	//	return
	//}
	//
	//logger.Println("----------", i, "----------")
	//
	//if newRespData.Total != oldRespData.Total || len(newRespData.TaskInfo) != len(oldRespData.TaskInfo) {
	//	if strings.Index(replayData.Path, "record_type=1") > 0 {
	//		if newRespData.Total == (oldRespData.RecordTotal - oldRespData.Total) {
	//			return
	//		}
	//	}
	//
	//	logger.Println(string(newRespBody))
	//	logger.Println(newRespData)
	//	logger.Println(newUrl)
	//	logger.Println(string(oldRespBody))
	//	logger.Println(oldRespData)
	//	logger.Println(oldUrl)
	//	logger.Printf("%+v \n", replayData)
	//} else {
	//	if len(newRespData.TaskInfo) == 0 {
	//		return
	//	}
	//
	//	for key, val := range newRespData.TaskInfo {
	//		if val.TaskId != oldRespData.TaskInfo[key].TaskId ||
	//			//val.Filesize != oldRespData.TaskInfo[key].Filesize ||
	//			//val.Filename != oldRespData.TaskInfo[key].Filename ||
	//			val.RecordType != oldRespData.TaskInfo[key].RecordType {
	//			logger.Println(string(newRespBody))
	//			logger.Println(newRespData)
	//			logger.Println(newUrl)
	//			logger.Println(string(oldRespBody))
	//			logger.Println(oldRespData)
	//			logger.Println(oldUrl)
	//			logger.Printf("%+v \n", replayData)
	//		}
	//	}
	//}

}

func GetFilenameSuffixCaseSensitive(path string) string {
	name := filepath.Base(path)
	i := strings.LastIndexByte(name, '.')
	if i >= 0 && i+1 < len(name) {
		suf := name[i+1:]
		if isValidSuffix(suf) {
			return suf
		}
	}
	return ""
}

func isValidSuffix(s string) bool {
	if s == "" {
		return false
	}
	for _, v := range strings.ToLower(s) {
		if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'z') {
			continue
		} else {
			return false
		}
	}
	return true
}

func GetFilenameSuffix(path string) string {
	return strings.ToLower(GetFilenameSuffixCaseSensitive(path))
}

func parseSuffixMap(m map[string]string) map[string]string {
	ret := make(map[string]string)
	for k, v := range m {
		k = strings.TrimSpace(k)
		for _, s := range strings.Split(strings.ToLower(v), ",") {
			ret[strings.TrimSpace(s)] = k
		}
	}
	return ret
}

const (
	Unknown  = "unknown"
	Folder   = "folder"
	Video    = "video"
	Text     = "text"
	Image    = "image"
	Audio    = "audio"
	Archive  = "archive"
	Font     = "font"
	Document = "document"
)


var defaultSuffixMap = map[string]string{
	Video:    "3g2,3gp,asf,asx,avi,divx,dv,flv,f4v,m2ts,m4v,mkv,mov,mp4,mpe,mpeg,mpg,qt,rm,rmvb,ts,vob,webm,wmv,xv",
	Text:     "csv,html,htm,css,ini,json,tsv,xml,yaml,yml,md,markdown,cnf,conf,cfg,log,txt",
	Image:    "ai,bmp,cdr,cr2,dwg,dxf,eps,exif,fpx,gif,hdri,heif,ico,jfif,jif,jpe,jpeg,jpeg2000,jpg,jxr,pcd,pcx,png,psd,raw,svg,tga,tif,tiff,ufo,webp,wmf",
	Audio:    "aac,aiff,amr,cda,flac,m4a,mid,mp3,ogg,vqf,wav,wma",
	Archive:  "7z,ar,bz2,cab,crx,dcm,deb,elf,eot,epub,exe,gz,iso,lz,nes,ps,rar,rpm,rtf,sqlite,swf,tar,xz,z,zip",
	Document: "doc,docx,ppt,pptx,xls,xlsx,pdf",
	Font:     "otf,ttf,woff,woff2",
}

func main() {
	//var suffixToType map[string]string
	suffixToType := parseSuffixMap(defaultSuffixMap)
	log.Println(suffixToType)
	////a := "<p><img src=\"https://img-xlppc-zhan.oss-cn-shanghai.aliyuncs.com/0dee70e18b7117b9a4eb15c90bd977c58ffd0c73?Expires=1625720782&amp;OSSAccessKeyId=LTAI5t9YGZ8Q96W7Xn4Am1Yb&amp;Signature=AhBW4kMinj%2BJ%2FGJbm436pif4PgA%3D\" width=\"640\"></p><p><strong>滴滴被查</strong></p><p>滴滴上市上出了大麻烦，但是如果不上市，可能麻烦会更大。</p><p><strong>6月30日，滴滴在纽交所挂牌，既没有撞钟，也没有开新闻发布会，而且选择的时间又是非常的敏感，</strong>刻意为之是显而易见的。</p><p>但是谁也没想到，中国政府的反应更加迅猛。7月2日，国家网信办就宣布对他进行网络安全检查，而且暂停了新客户的发展。到<strong>7月4日，调查结论就基本上已经出来了，指控滴滴严重违法违规收集客户个人信息，做出的处罚决定也是非常严厉的，就是滴滴出行App下架。</strong></p><p><strong>上市是一场赌博？</strong></p><p>这个消息一出，舆论哗然，大家纷纷关注滴滴复牌之后，在美国股市会跌多少？另外，也担心会引发美国投资者的集团诉讼。</p><p>我想这一切，应该都在滴滴的臆想之中。</p><p>事实上，<strong>如果中国网信办的调查在滴滴上市之前展开的话，恐怕滴滴再要想短期上市的可能性就几乎没有了。所以和上市以后股价暴跌，包括被美国投资者诉讼相比，上市不成恐怕影响更大。</strong></p><p><strong>等待滴滴的将是什么？</strong></p><p>谁都知道，滴滴现在是中国乃至世界最大的互联网出租车平台，他现在的活跃用户在全球有4.93亿，中国大陆有3.77亿，司机在全球活跃用户有1500万，在中国有1300万左右。</p><p><strong>2015年，滴滴和快的合并，后来又跟Uber合并。</strong>那么在中国基本上是一股独大，占有垄断地位。</p><p>那么这么巨大的数据，这么巨大的交易额，他手中掌握的这些数据也是非常非常重要的。在中国政府越来越强调网络数据安全的背景之下，滴滴实际上是有义务配合政府做好个人用户的隐私保护的。</p><p>那么<strong>滴滴之所以这一次闯关，我想还是来自于股东的压力。</strong>这一次滴滴在美国上市，发行价是14美元，发行2.88亿股，融了大概40亿美元的资。这点资金对于现在依然处在亏损之中的滴滴来讲，应该讲还是一个巨大的支撑。</p><p>更重要的是，那些境外的投资者，不管是软银也好，还是Uber也好，他们的投资份额有可能在纽交所变现，我想这才是更为重要的。</p><p><strong>滴滴现在的态势已经到了大而不能倒的地步，也是一个系统重要性的社会基础设施，所以等待着滴滴的，应该讲是一个规范整顿，规范发展。</strong>但是，上不上市就不一样了。</p>"
	//a := "<p>（原标题：坚决防止未成年人沉迷网络游戏 ——国家新闻出版署有关负责人就《关于进一步严格管理 切实防止未成年人沉迷网络游戏的通知》答记者问）</p><p>国家新闻出版署近日印发《关于进一步严格管理 切实防止未成年人沉迷网络游戏的通知》。国家新闻出版署有关负责人就通知有关背景和要求，接受了记者专访。</p><p>问：请介绍一下通知出台的背景。</p><p>答：近年来，我国网络游戏产业在快速发展的同时，也出现一些突出问题，特别是未成年人沉迷网络游戏问题引起社会广泛关切，广大家长反映强烈。国家新闻出版署高度重视防止未成年人沉迷网络游戏工作，2019年印发了《关于防止未成年人沉迷网络游戏的通知》"
	//
	//
	//b := []rune(a)
	//if len(b) < 300 {
	//	log.Println()
	//}
	//log.Println(string(b[0:32]))
	//t1 := "----"
	//if _, err := strconv.Atoi(t1); err != nil {
	//	log.Println(1111)
	//}
	//
	//timestamp, _ := time.Parse("20060102", t1)
	//println(timestamp.Unix())

	//urla := fmt.Sprintf("http://cache4.vip.xunlei.com:8001/cache?userid=715982997&protocol_version=121")
	//req, err := http.NewRequest("GET", urla, nil)
	//if err != nil {
	//	log.Println(err)
	//
	//}
	//clt := http.Client{}
	//resp, err := clt.Do(req)
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//
	//log.Println(body)
	//
	//str := "ret=0&userid=715982997&isvip=1&uservas=2&level=3&grow=2970&payid=5&payname=年费支付方式&daily=18&expire=20230215&autodeduct=0&remind=0&isyear=1&month_expire=20210815&vas_type=3&register=20210201"
	//
	//u, err := url.ParseQuery(str)
	//log.Println(u.Get("isvip11"))


	//go redisInit()
	//
	//i := 0
	//for {
	//	select {
	//	case msg := <- ch:
	//		i++
	//		go test(msg, i)
	//	}
	//}
}

func redisInit() {
	for {
		if rdb.LLen(ctx, "test").Val() < 5 {
			time.Sleep(2 * time.Second)
		}

		redisJson, err := rdb.RPop(ctx, "collection_list").Result()
		replayData := Replay{}
		if err = json.Unmarshal([]byte(redisJson), &replayData); err != nil {
			log.Println(err)
		}

		ch <- replayData
	}
}

func NewReq(reqUrl string, params Replay) (*ListTaskResponse, []byte, string, error) {
	url := fmt.Sprintf("%s/%d?%s", reqUrl, params.UserID, params.Path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, nil, url, err
	}

	req.Header.Set("accept", params.Accept)
	req.Header.Set("authorization", params.Auth)

	clt := http.Client{}
	resp, err := clt.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Println(err)
		return nil, nil, url, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, nil, url, err
	}

	data := new(ListTaskResponse)
	if err = json.Unmarshal(body, data); err != nil {
		log.Println(err)
		return nil, nil, url, err
	}

	return data, body, url, nil
}
