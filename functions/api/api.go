// netlify函数这里必须是main
package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()
	RegisterHandler(r)

	ginLambda = ginadapter.New(r)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	res, err := ginLambda.ProxyWithContext(ctx, req)
	if err != nil {
		log.Default().Println(err)
	}
	return &res, err
}

func main() {
	lambda.Start(handler)
}

func RegisterHandler(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 0,
			"msg":  "hello world",
		})
	})
}
