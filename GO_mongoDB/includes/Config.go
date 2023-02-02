package includes

import "go.mongodb.org/mongo-driver/mongo"

var (
	DB_USERNAME                         = "root"
	DB_PASSWORD                         = "Kush@789#"
	DB_HOST                             = "localhost"
	DB_NAME                             = "THIRDESSENTIAL"
	UserCollection    *mongo.Collection = GetCollection(DB, "users")
	ProductCollection *mongo.Collection = GetCollection(DB, "product")

	// SERVER_SUB_DOMAIN_NAME                   = "dev."
	// SERVER_DOMAIN_NAME                       = "wikibedtimestories.com"
	SERVER_HOST = "localhost"
	// SERVER_NAME                      = SERVER_SUB_DOMAIN_NAME + SERVER_DOMAIN_NAME
	SERVER_PORT = "1010"
	// SERVER_PRODUCT_IMAGE_LOCATION    = "https://" + SERVER_NAME + "/webservices/Thirdessential_GO/products_image/"
	// SERVER_IMAGE_LOCATION_FOR_DELETE = "products_image/"
	// SERVER_CERT_FILE                 = "/etc/letsencrypt/live/" + SERVER_NAME + "/fullchain.pem"
	// SERVER_KEY_FILE                  = "/etc/letsencrypt/live/" + SERVER_NAME + "/privkey.pem"
)
