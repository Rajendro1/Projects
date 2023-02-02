package includes

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func CreateUserToDb(name, email, password string) ([]User, error) {
	newUsers_array := []User{}
	newUsers := User{
		Id:       primitive.NewObjectID(),
		Name:     name,
		Email:    email,
		Password: password,
	}

	_, err := UserCollection.InsertOne(context.Background(), newUsers)
	if err != nil {
		return newUsers_array, err
	}
	newUsers_array = append(newUsers_array, newUsers)
	return newUsers_array, nil
}

func GetUsersFrDb(userID string) ([]User, error) {
	var user User
	user_slice := []User{}

	objId, _ := primitive.ObjectIDFromHex(userID)

	err := UserCollection.FindOne(context.Background(), bson.M{"id": objId}).Decode(&user)
	if err != nil {
		return user_slice, err
	}
	user_slice = append(user_slice, user)
	return user_slice, nil
}

func GetPasswordByEmail(email string) ([]User, error) {
	var user User
	user_slice := []User{}

	err := UserCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user_slice, err
	}
	user_slice = append(user_slice, user)
	return user_slice, nil
}
func CreateProductToDb(productName, productPrice string) ([]Product, error) {
	newProduct_slice := []Product{}
	newProduct := Product{
		Id:    primitive.NewObjectID(),
		Name:  productName,
		Price: productPrice,
	}
	_, err := ProductCollection.InsertOne(context.Background(), newProduct)
	if err != nil {
		return newProduct_slice, err
	}
	newProduct_slice = append(newProduct_slice, newProduct)
	return newProduct_slice, nil
}

func GetProductFrDb(productID string) ([]Product, error) {
	var product Product
	product_slice := []Product{}
	objId, _ := primitive.ObjectIDFromHex(productID)
	err := ProductCollection.FindOne(context.Background(), bson.M{"id": objId}).Decode(&product)
	if err != nil {
		return product_slice, err
	}
	product_slice = append(product_slice, product)
	return product_slice, nil
}
func GetAllProductsFrDb() ([]Product, error) {
	var product Product
	product_slice := []Product{}
	cur, err := ProductCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return product_slice, err
	}
	for cur.Next(context.Background()) {
		err := cur.Decode(&product)
		if err != nil {
			return product_slice, err
		}
		product_slice = append(product_slice, product)
	}
	defer cur.Close(context.Background())
	return product_slice, nil
}

func DeleteUsersFrDB(userID string) bool {
	objId, _ := primitive.ObjectIDFromHex(userID)
	_, err := UserCollection.DeleteOne(context.Background(), bson.M{"id": objId})
	if err != nil {
		return false
	}
	return true
}

func DeleteProductFrDB(productID string) bool {
	objId, _ := primitive.ObjectIDFromHex(productID)
	_, err := ProductCollection.DeleteOne(context.Background(), bson.M{"id": objId})
	if err != nil {
		return false
	}
	return true
}
func DeleteAllProductFrDB() bool {
	_, err := ProductCollection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		return false
		// log.Println(err.Error())
	}
	return true
}

func UpdateAllProductToDB(productName, productPrice string) bool {
	newProduct_slice := []Product{}
	newProduct := Product{
		Id:    primitive.NewObjectID(),
		Name:  productName,
		Price: productPrice,
	}
	_, err := ProductCollection.UpdateMany(context.Background(), bson.D{{}}, newProduct)
	if err != nil {
		return false
		// log.Println(err.Error())
	}
	newProduct_slice = append(newProduct_slice, newProduct)
	return true
}
