package controllers

import (
	"context"
	"net/http"
	"theSiS/web-service-fiber/configs"
	"theSiS/web-service-fiber/models"
	"theSiS/web-service-fiber/responses"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "datas")
var validate = validator.New()

func CreateProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var product models.Product
	defer cancel()

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&product); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newProduct := models.Product{
		ID:          product.ID,
		NAME:        product.NAME,
		DESCRIPTION: product.DESCRIPTION,
		IMAGEURL:    product.IMAGEURL,
	}

	result, err := userCollection.InsertOne(ctx, newProduct)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.ProductResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetProducts(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var products []models.Product

	defer cancel()
	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var dataItem models.Product
		if err = results.Decode(&dataItem); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		products = append(products, dataItem)
	}

	return c.Status(http.StatusOK).JSON(
		responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": products}},
	)

}

func UpdateProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	productId := c.Params("productId")
	var product models.Product
	defer cancel()

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&product); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": product.NAME, "description": product.DESCRIPTION, "imageUrl": product.IMAGEURL}

	result, err := userCollection.UpdateOne(ctx, bson.M{"id": productId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	var updatedProduct models.Product
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"id": productId}).Decode(&updatedProduct)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedProduct}})
}

func DeleteProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	productId := c.Params("productId")
	defer cancel()

	result, err := userCollection.DeleteOne(ctx, bson.M{"id": productId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.ProductResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
	)
}
