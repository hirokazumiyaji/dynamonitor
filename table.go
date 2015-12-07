package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/zenazn/goji/web"
)

func tableList(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	tableName := r.URL.Query().Get("table")
	limit := r.URL.Query().Get("limit")

	input := dynamodb.ListTablesInput{}
	if tableName != "" {
		input.ExclusiveStartTableName = aws.String(tableName)
	}
	if limit != "" {
		n, _ := strconv.ParseInt(limit, 10, 64)
		input.Limit = &n
	}
	res, err := client.ListTables(&input)
	if err != nil {
		Error(w, 500, []error{err})
		return
	}

	json.NewEncoder(w).Encode(res)
}

func tableDetail(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	tableName := c.URLParams["tablename"]

	res, err := client.DescribeTable(
		&dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		},
	)
	if err != nil {
		Error(w, 500, []error{err})
		return
	}

	json.NewEncoder(w).Encode(res)
}
