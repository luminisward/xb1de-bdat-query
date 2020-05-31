package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"xb1de-bdat-query/db"
	"xb1de-bdat-query/models"
)

var languages = []string{
	"cn", "fr", "gb", "ge", "it", "jp", "kr", "sp", "tw",
}

func Translations(c *gin.Context) {
	var q models.Query
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(400, gin.H{"errmsg": err.Error()})
		return
	}

	q.QueryString = strings.TrimSpace(q.QueryString)
	if q.QueryString == "" {
		c.JSON(400, gin.H{"errmsg": "query_string can't be empty"})
		return
	}

	if !contains(languages, q.QueryLanguage) {
		c.JSON(400, gin.H{"errmsg": "query_language is not exist"})
		return
	}

	for _, resultLanguage := range q.ResultLanguages {
		if !contains(languages, resultLanguage) {
			c.JSON(400, gin.H{"errmsg": "result_languages: " + resultLanguage + " is not exist"})
			return
		}
	}

	if q.Limit <= 0 {
		q.Limit = 250
	}
	if q.Limit > 500 {
		q.Limit = 500
	}

	debug := c.Query("monado")
	if debug != "" {
		q.Limit = 0
	}

	xb1db := db.GetPgsql()

	var tablenames []string
	err := xb1db.Select(&tablenames, fmt.Sprintf("SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = '%s'", q.QueryLanguage))
	if err != nil {
		log.Fatal(err)
	}

	results := make(map[string][]models.Ms)

	for _, tablename := range tablenames {

		// 限定bdat或table
		if len(q.Bdats) > 0 || len(q.Tables) > 0 {
			s := strings.Split(tablename, ".")

			if len(q.Bdats) > 0 && !contains(q.Bdats, s[0]) {
				continue
			}

			if len(q.Tables) > 0 && !contains(q.Tables, s[1]) {
				continue
			}
		}

		var tempRows []models.Ms
		err := xb1db.Select(&tempRows, fmt.Sprintf("SELECT * FROM %s.\"%s\" WHERE name ILIKE '%%%s%%'", q.QueryLanguage, tablename, q.QueryString))
		if err != nil {
			log.Fatal(err)
		}
		if len(tempRows) > 0 {
			for index := range tempRows {
				tempRows[index].SetTableName(tablename)
			}
			results[q.QueryLanguage] = append(results[q.QueryLanguage], tempRows...)
		}

		if q.Limit > 0 && len(results[q.QueryLanguage]) >= q.Limit {
			break
		}
	}

	// 根据结果查询其他语言
	for _, resultLanguage := range q.ResultLanguages {

		if resultLanguage == q.QueryLanguage {
			continue
		}

		for _, row := range results[q.QueryLanguage] {
			var tempRow models.Ms
			err := xb1db.Get(&tempRow, fmt.Sprintf("SELECT * FROM %s.\"%s\" WHERE row_id = '%d'", resultLanguage, row.GetTableName(), row.RowId))
			if err != nil {
				log.Fatal(err)
			}
			tempRow.SetTableName(row.GetTableName())
			results[resultLanguage] = append(results[resultLanguage], tempRow)

		}
	}

	c.JSON(200, gin.H{
		"language":     q.QueryLanguage,
		"query_string": q.QueryString,
		"bdats":        q.Bdats,
		"tables":       q.Tables,
		"limit":        q.Limit,
		"total":        len(results[q.QueryLanguage]),
		"results":      results,
	})

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
