package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"iqge.org/gym/utils"
)

type GymController struct{}

func (ctrl GymController) CreateGym(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in someFunction", r)
		}
	}()
	inserted, err := utils.GymCollection().InsertOne(context.TODO(), bson.M{"hello": "world"})
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"hello": "There", "inserted": inserted})
}
