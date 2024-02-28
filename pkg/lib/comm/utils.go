package _commUtils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/emirpasic/gods/maps"
	"math/rand"
	"net"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func Base(pathStr string) string {
	pathStr = filepath.ToSlash(pathStr)
	return path.Base(pathStr)
}

func RemoveBlankLine(str string) string {
	myExp := regexp.MustCompile(`\n{3,}`) // 连续换行
	ret := myExp.ReplaceAllString(str, "\n\n")

	ret = strings.Trim(ret, "\n")
	ret = strings.TrimSpace(ret)

	return ret
}

func GetOs() string {
	osName := runtime.GOOS

	if osName == "darwin" {
		return "mac"
	} else {
		return osName
	}
}
func IsWin() bool {
	return GetOs() == "windows"
}
func IsLinux() bool {
	return GetOs() == "linux"
}
func IsMac() bool {
	return GetOs() == "mac"
}

func UpdateUrl(url string) string {
	if strings.LastIndex(url, "/") < len(url)-1 {
		url += "/"
	}

	return url
}

func IntToStrArr(intArr []int) (strArr []string) {
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}

	return
}
func UintToStrArr(intArr []uint) (strArr []string) {
	for _, i := range intArr {
		strArr = append(strArr, fmt.Sprintf("%d", i))
	}

	return
}

func LinkedMapToMap(mp maps.Map) map[string]string {
	ret := make(map[string]string, 0)

	for _, keyIfs := range mp.Keys() {
		valueIfs, _ := mp.Get(keyIfs)

		key := strings.TrimSpace(keyIfs.(string))
		value := strings.TrimSpace(valueIfs.(string))

		ret[key] = value
	}

	return ret
}

func IsRelease() bool {
	arg1 := strings.ToLower(os.Args[0])
	name := filepath.Base(arg1)

	return strings.Index(name, "___") != 0 && strings.Index(arg1, "go-build") < 0
}

func GetUserHome() string {
	userProfile, _ := user.Current()
	home := userProfile.HomeDir
	return home
}

func IsPortInUse(port int) bool {
	if conn, err := net.DialTimeout("tcp", net.JoinHostPort("", fmt.Sprintf("%d", port)), 3*time.Second); err == nil {
		conn.Close()
		return true
	}
	return false
}

func IsDisable(enable string) bool {
	if enable == "1" {
		return false
	} else {
		return true
	}
}

func JsonEncode(data interface{}) (res string) {

	if resByte, err := json.Marshal(data); err == nil {
		res = string(resByte)
	} else {
		panic(err)
	}

	return

}

func JsonDecode(str string, res interface{}) (err error) {
	if str == "" {
		return nil
	}
	if err = json.Unmarshal([]byte(str), res); err != nil {
		//panic(err)
		logUtils.Error(err.Error())
	}

	return

}

func ArrayRemoveDuplication(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}

func ArrayRemoveUintDuplication(arr []uint) []uint {
	set := make(map[uint]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}

func Map2Struct(m interface{}, s interface{}) {
	JsonDecode(JsonEncode(m), s)
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesCBCEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesCBCDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

// InSlice 判断字符串是否在切片中
func InSlice(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

// IntInSlice 判断uint是否在切片中
func IntInSlice(target int, intArray []int) bool {
	sort.Ints(intArray)
	index := sort.SearchInts(intArray, target)
	if index < len(intArray) && intArray[index] == target {
		return true
	}
	return false
}

// Intersect 求交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// Difference 差集
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

// IntersectUint 求交集
func IntersectUint(slice1, slice2 []uint) []uint {
	m := make(map[uint]int)
	nn := make([]uint, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// DifferenceUint 差集
func DifferenceUint(slice1, slice2 []uint) []uint {
	m := make(map[uint]int)
	nn := make([]uint, 0)
	inter := IntersectUint(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

func InArray(target string, array []string) bool {
	for _, item := range array {
		if item == target {
			return true
		}
	}
	return false
}

func GetEnvVar(name, def string) (ret string) {
	ret = os.Getenv(name)

	if ret == "" {
		ret = def
	}

	return
}

// 驼峰式写法转为下划线写法
func Camel2Case(name string) (ret string) {
	var buffer []rune
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer = append(buffer, '_')
			}
			buffer = append(buffer, unicode.ToLower(r))
		} else {
			buffer = append(buffer, r)
		}
	}
	return string(buffer)

}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// Sha256 Sha256加密
func Sha256(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

func RandStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func EncryptHmacMd5(key, data string) string {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

func GetSign(appKey, appSecret, nonce, timestamp, body string) (sign string) {
	preSignStr := strings.Join([]string{appKey, timestamp, nonce, body}, "")
	sign = EncryptHmacMd5(appSecret, preSignStr)
	return
}

func ArrayUnique(arr []string) (ret []string) {
	temp := map[string]bool{}
	for _, x := range arr {
		if _, ok := temp[x]; !ok {
			ret = append(ret, x)
		}
	}
	return
}

func UintArrToStr(arr []uint) (res string) {
	for _, item := range arr {
		if res != "" {
			res += fmt.Sprintf(",%v", item)
		} else {
			res += fmt.Sprintf("%v", item)
		}
	}
	return
}
