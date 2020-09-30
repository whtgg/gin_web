package api

import (
	"dms/app"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"errors"
	"github.com/axgle/mahonia"
)

type Pagination struct {
	PageSize  int `form:"page_size" json:"page_size"`
	PageNo    int `form:"page_no" json:"page_no"`
	Total     int
	TotalPage int
	First     int
	Before    int
	Next      int
}

func (p *Pagination) GetPageNo() int {
	if p.PageNo <= 0 {
		p.PageNo = 1
	}
	return p.PageNo
}

func (p *Pagination) GetPageSize() int {
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	return p.PageSize
}

func (p *Pagination) GetTotalPage() int {
	if p.TotalPage == 0 && p.Total > 0 {
		p.TotalPage = int(math.Ceil(float64(p.Total) / float64(p.GetPageSize())))
	}
	return p.TotalPage
}

func (p *Pagination) GetBeforePageNo() int {
	if p.GetPageNo() > 1 {
		//if p.GetPageNo() >= p.GetTotalPage() {
		//	return p.GetTotalPage() - 1
		//}
		return p.PageNo - 1
	}
	return p.PageNo
}

func (p *Pagination) GetNextPageNo() int {
	// 初始化 TotalPage
	if p.GetPageNo() >= p.GetTotalPage() {
		return p.GetPageNo()
	}
	return p.GetPageNo() + 1
}

/**
* 图片上传
*/
func UploadImage(c *gin.Context) {
	var err error
	f, err := c.FormFile("file")
	if err == nil {
		rand.Seed(time.Now().UnixNano())
		uploadDir := app.Config.GetString("upload_path") + time.Now().Format("2006-01-02") + string(os.PathSeparator)
		if _, err = os.Stat(uploadDir); err != nil {
			os.Mkdir(uploadDir, os.ModePerm)
		}
		filePath := uploadDir + fmt.Sprintf("%v", time.Now().Unix()) + strconv.Itoa(rand.Intn(100)) + ".png"
		c.SaveUploadedFile(f, filePath)
		c.JSON(http.StatusOK, gin.H{"data": "/" + filePath})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"msg": err})
	return
}

/**
*防止中文乱码
*/
func convertEncoding(srt, eno string) string{
	enc := mahonia.NewDecoder(eno)
	return enc.ConvertString(srt)
}

/**
* 切换数据库
*/
func ChangeMysql(c *gin.Context)  (*gorm.DB,error){
	var (
		db *gorm.DB
		err error
		rev interface{}
	)

	key,bo := c.Get("token")
	if bo {
		rev, err = app.RedisConn.Do("GET",key.(string)+"-site-db")
		if err != nil {
			log.Printf("%s",err.Error())
			return db,errors.New("链接失败")
		}
	}
	db,err = gorm.Open("mysql",fmt.Sprintf("%s",rev)+"&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("打开数据库 %s",err.Error())
		return nil,err
	}
	return db,nil
}


func HideStar(str string) (result string) {
	if str == "" {
		return "***"
	}
	if strings.Contains(str, "@") {
		res := strings.Split(str, "@")
		if len(res[0]) < 3 {
			resString := "***"
			result = resString + "@" + res[1]
		} else {
			res2 := Substr2(str, 0, 3)
			resString := res2 + "***"
			result = resString + "@" + res[1]
		}
		return result
	} else {
		reg := `^1[0-9]\d{9}$`
		rgx := regexp.MustCompile(reg)
		mobileMatch := rgx.MatchString(str)
		if mobileMatch {
			result = Substr2(str, 0, 3) + "****" + Substr2(str, 7, 11)
		} else {
			nameRune := []rune(str)
			lens := len(nameRune)

			if lens <= 1 {
				result = "***"
			} else if lens == 2 {
				result = string(nameRune[:1]) + "*"
			} else if lens == 3 {
				result = string(nameRune[:1]) + "*" + string(nameRune[2:3])
			} else if lens == 4 {
				result = string(nameRune[:1]) + "**" + string(nameRune[lens-1:lens])
			} else if lens > 4 {
				result = string(nameRune[:2]) + "***" + string(nameRune[lens-2:lens])
			}
		}
		return
	}
}

// https://www.iphpt.com/detail/114
// 各种字符串加星
// 匹配: 手机号,邮箱,中文,身份证等等 用于数据脱敏
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}


func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,User-Token,Authorization,Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}


