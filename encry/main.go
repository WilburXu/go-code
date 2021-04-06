package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"math/rand"
	"os"
	"time"
)

var (
	nonce = "9I0fMA(GC#qG"
)

type LogSt struct {
	RequestId string
	RespBody  string
	Path      string
}

type LogListSt []LogSt
const soh = string(1)

func main() {
	//ReadTestFileContentOfLine("./log/cleaning.log")
	//ReadTestFileSort("./log/test/split/")


	//ReadReleaseFileContentOfLine("./log/team.proxy.access.1617307201.log")
}


func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// 获取此 slice 的长度
func (p LogListSt) Len() int { return len(p) }

// 根据元素的requestId降序排序 （此处按照自己的业务逻辑写）
func (p LogListSt) Less(i, j int) bool {
	return p[i].RequestId > p[j].RequestId
}

// 交换数据
func (p LogListSt) Swap(i, j int) { p[i], p[j] = p[j], p[i] }


//func main() {
//	// Engin
//	router := gin.Default()
//	//router := gin.New()
//
//	router.POST("/hello", func(c *gin.Context) {
//		log.Println(">>>> hello gin start <<<<")
//
//		body, _ := c.GetRawData()
//		log.Printf("--11 %v", string(body))
//		log.Printf("--12 %v", c.Request.Header.Get("X-Xc-Proto-Req"))
//
//		decodeBody, _ := AESGCMDecode(c.Request.Header.Get("X-Xc-Proto-Req"), body)
//		log.Printf("--22 %v", string(decodeBody))
//		if len(decodeBody) > 0 {
//
//		}
//
//		data := map[string]interface{}{
//			"token": "c162e6bb9435208bc87d18a9b599e9338bfd93070cc6e14a04af8486d6906923",
//			"h_m": 10,
//		}
//		dataJson, _ :=json.Marshal(data)
//		log.Printf("--33 %v", string(dataJson))
//
//		res, decodeKey, _ := AESGCMEncode(dataJson)
//		log.Printf("--44 %x", res)
//		log.Printf("--45 %s", decodeKey)
//
//		c.Writer.Header().Set("X-Xc-Proto-Req", decodeKey)
//
//		log.Printf("--50 %x", res)
//
//		//c.Data(http.StatusOK, "application/json", res)
//		c.String(http.StatusOK, fmt.Sprintln(res))
//	})
//
//	// 指定地址和端口号
//	router.Run("localhost:9988")
//}

func AESGCMDecode(encodeKey string, body []byte) (decodeBody []byte, err error) {
	key, _ := hex.DecodeString(encodeKey)
	nonce := []byte(nonce)
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return
	}

	cipherText, _ := hex.DecodeString(string(body))
	//cipherText := body
	decodeBody, err = aesgcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return
	}

	return decodeBody, nil
}

func AESGCMEncode(body []byte) (encodeBody []byte, decodeKey string, err error) {
	decodeKey = GetRandomHexString(32)
	key, _ := hex.DecodeString(decodeKey)
	//key, _ := hex.DecodeString("04ab9d9b8b7f923472b5259448312dc5")

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	//nonce := make([]byte, 12)
	nonce := []byte(nonce)

	aesgcm, err := cipher.NewGCMWithNonceSize(block, 12)
	if err != nil {
		return
	}

	plainText := body
	encodeBody = aesgcm.Seal(nil, nonce, plainText, nil)

	return
}

func GetRandomHexString(lens int) string {
	str := "0123456789abcde"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
