package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"cms-front/model"
	"github.com/jinzhu/gorm"
	"cms-front/app"
	"strconv"
	"strings"
	"fmt"
)

type FrontStruct struct {}

type collect struct {
	Collection		string
}

var Front FrontStruct


/**
* 栏目列表查询
*/
func (a *FrontStruct) SearchByCon(c *gin.Context) {
	var (
		page 			Pagination
		err				error
		itemsTop		[]*model.Categorys
		itemsLeft		[]*model.Categorys
		cateTitle		string
		articles		[]*model.Articles
		db				*gorm.DB = app.LoadDB()
		tags			[]*model.Tags

	)

	parentId,_ := strconv.Atoi(c.DefaultQuery("cate","2"))

	//导航栏目
	dbTop := db.Where("parent_id =  ?", 0)
	if err = dbTop.Limit(7).Offset(0).Find(&itemsTop).Error; err != nil {
		log.Printf("查询数据失败 %s", err.Error())
		return
	}

	//左边栏目
	dbLeft := db.Where("parent_id =  ?", parentId)
	if err = dbLeft.Find(&itemsLeft).Error; err != nil {
		log.Printf("查询数据失败 %s", err.Error())
		return
	}

	//栏目词
	for _,v := range itemsTop {
		if v.Id == uint(parentId) {
			cateTitle = v.Title
		}
	}

	type collectionTag struct {//tag 信息
		Id			uint
		Collection	string
	}

	type collectionArticle struct {
		Collection	string
	}


	var col collectionTag
	var colArticle collectionArticle
	//栏目相关词
	if err = db.Table("tags").Select("tags.id,relate_tags.collection").Joins("left join relate_tags on relate_tags.tag_id = tags.id").Where("tags.tag = ?",cateTitle).Scan(&col).Error;err != nil {
		log.Println(err)
		c.Abort()
	}

	//相关词文章
	if err = db.Table("relate_tag_article").Where("tag_id = ?",col.Id).Scan(&colArticle).Error;err != nil {
		log.Println("获得相关文章ID",err.Error())
		c.Abort()
	}

	//文章
	if err = c.BindQuery(&page); err != nil {
		log.Printf("请求参数错误 %s", err.Error())
		c.Abort()
	}

	if err = db.Limit(page.GetPageSize()).Offset((page.GetPageNo()-1)*page.GetPageSize()).Where("id in (?)",strings.Split(colArticle.Collection,",")).Find(&articles).Error; err != nil {
		log.Printf("查询数据失败 %s", err.Error())
		c.Abort()
	}

	//栏目词
	for _,v := range itemsTop {
		if v.Id == uint(parentId) {
			cateTitle = v.Title
		}
	}

	//相关词相关词
	if err = db.Where("id in (?)",strings.Split(col.Collection,",")).Find(&tags).Error;err != nil {
		log.Println(err)
		c.Abort()
	}

	var nums []int
	if page.PageNo > 9 {
		for i := 1; i<10;i++ {
			nums = append(nums,i)
		}
	} else {
		for i := 1; i<page.PageNo+1;i++ {
			nums = append(nums,i)
		}
	}
	fmt.Println()
	c.HTML(http.StatusOK,"list.html",gin.H{
		"items_left":       itemsLeft,
		"items_top":		itemsTop,
		"articles":			articles,
		"cate_title":		cateTitle,
		"tags":				tags,
		"cate_type":		parentId,
		"first":       		1,
		"current":     		page.PageNo,
		"before":      		page.GetBeforePageNo(),
		"next":        		page.GetNextPageNo(),
		"total_items": 		page.Total,
		"total_pages": 		page.TotalPage,
		"display_nums":     nums,
	})
	return
}

func (a *FrontStruct) SearchDetail(c *gin.Context) {
	articleId,_ := strconv.Atoi(c.DefaultQuery("id","0"))

	var (
		article		model.Articles
		articles	[]*model.Articles
		db			*gorm.DB = app.LoadDB()
		err			error
		col			collect
	)
	//获得文章内容
	if err = db.Where("id = ?",articleId).Find(&article).Error;err != nil {
		log.Println("获取文章失败：",err.Error())
		c.Abort()
	}

	//获得相关文章
	if err = db.Table("relate_article").Where("article_id = ?",articleId).Scan(&col).Error;err != nil {
		log.Println("获取相关文章ID失败：",err.Error())
		c.Abort()
	}

	if err = db.Where("id in (?)",strings.Split(col.Collection,",")).Find(&articles).Error;err != nil {
		log.Println("获取相关文章失败：",err.Error())
		c.Abort()
	}

	c.HTML(http.StatusOK,"list-con.html",gin.H{
		"article":	article,
		"article_tag": strings.Split(article.Tag,","),
	})
}


func (a *FrontStruct) SearchByTags(c *gin.Context) {
	tagContent := c.Query("tag")
	var (
		db = app.LoadDB()
		err error
		articles []*model.Articles
		tags 		[]*model.Tags
	)

	type collectionTag struct {//tag 信息
		Id			uint
		Collection	string
	}


	var col collectionTag
	var colTag collectionTag
	//栏目相关文章
	if err = db.Table("tags").Select("tags.id,relate_tag_article.collection").Joins("left join relate_tag_article on relate_tag_article.tag_id = tags.id").Where("tags.tag = ?",tagContent).Scan(&col).Error;err != nil {
		log.Println(err)
		c.Abort()
	}

	//栏目相关词
	if err = db.Table("tags").Select("tags.id,relate_tags.collection").Joins("left join relate_tags on relate_tags.tag_id = tags.id").Where("tags.tag = ?",tagContent).Scan(&colTag).Error;err != nil {
		log.Println(err)
		c.Abort()
	}

	//获得相关词文章
	if err = db.Where("id in (?)",strings.Split(col.Collection,",")).Find(&articles).Error;err != nil {
		log.Println("获取相关文章失败：",err.Error())
		c.Abort()
	}

	//相关词相关词
	if err = db.Where("id in (?)",strings.Split(col.Collection,",")).Find(&tags).Error;err != nil {
		log.Println(err)
		c.Abort()
	}

	c.HTML(http.StatusOK,"search.html",gin.H{
		"articles":	articles,
		"total":len(articles),
		"tag":tagContent,
		"tags":tags,
	})
}
