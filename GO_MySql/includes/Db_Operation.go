package includes

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var (
	db *sql.DB
)

func RegisterUser(name, email, phone, address, password string) (bool, error) {
	SqlQuery := (`INSERT INTO users(name, email, phone, address, password, login_time) VALUES (?, ?, ?, ?, ?, NOW())`)
	_, RegisterUserErr := db.Exec(SqlQuery, name, email, phone, address, password)
	if RegisterUserErr != nil {
		log.Println(RegisterUserErr.Error())
		return false, RegisterUserErr
	}
	return true, nil
}

func RegisterSuperAdmin(name, email, phone, address, password string) (bool, error) {
	SqlQuery := (`INSERT INTO superadmins(name, email, phone, address, password) VALUES (?, ?, ?, ?, ?)`)
	_, RegisterSuperAdminErr := db.Exec(SqlQuery, name, email, phone, address, password)
	if RegisterSuperAdminErr != nil {
		log.Println(RegisterSuperAdminErr.Error())
		return false, RegisterSuperAdminErr
	}
	return true, nil
}

func GetAllUsers() ([]GetAllUsersDetails, error) {
	var usersDetails GetAllUsersDetails
	usersDetails_array := []GetAllUsersDetails{}
	SqlQuery := (`SELECT id, name, email, phone, address FROM users`)
	usersRow, usersErr := db.Query(SqlQuery)
	if usersErr != nil {
		log.Println(usersErr.Error())
		return usersDetails_array, usersErr
	}
	for usersRow.Next() {
		usersScan := usersRow.Scan(&usersDetails.UserID, &usersDetails.Name, &usersDetails.Email, &usersDetails.Phone, &usersDetails.Address)
		if usersScan != nil {
			log.Println(usersScan.Error())
			return usersDetails_array, usersScan
		}
		usersDetails_array = append(usersDetails_array, usersDetails)
	}
	return usersDetails_array, nil
}

func GetAllProductByUsers(user_id int) ([]GetUsersProductDetails, error) {
	var usersProduct GetUsersProductDetails
	usersProduct_array := []GetUsersProductDetails{}
	SqlQuery := (`SELECT id, user_id, name, price, image FROM products WHERE user_id = ?`)
	productRow, productErr := db.Query(SqlQuery, user_id)
	fmt.Println(SqlQuery)
	if productErr != nil {
		log.Println(productErr.Error())
		return usersProduct_array, productErr
	}
	for productRow.Next() {
		productScan := productRow.Scan(&usersProduct.ProductID, &usersProduct.UserID, &usersProduct.ProductName, &usersProduct.ProductPrice, &usersProduct.ProductImageLink)
		usersProduct.ProductImageLink = SERVER_PRODUCT_IMAGE_LOCATION + usersProduct.ProductImageLink
		if productScan != nil {
			log.Println(productScan.Error())
			return usersProduct_array, productScan
		}
		usersProduct_array = append(usersProduct_array, usersProduct)
	}
	defer productRow.Close()
	return usersProduct_array, nil
}

func CreateProduct(user_id int, name, price, image string) (bool, error) {
	SqlQuery := (`INSERT INTO products(user_id, name, price, image) VALUES (?, ?, ?, ?)`)
	_, CreateProductErr := db.Exec(SqlQuery, user_id, name, price, image)
	if CreateProductErr != nil {
		log.Println(CreateProductErr.Error())
		return false, CreateProductErr
	}
	return true, nil
}

func UpdateProduct(name, price, image string, productID, user_id int) (bool, error) {
	SqlQuery := (`UPDATE products SET name = ?, price = ?, image = ? WHERE id = ? AND user_id = ?`)
	_, UpdateProductErr := db.Exec(SqlQuery, name, price, image, productID, user_id)
	if UpdateProductErr != nil {
		log.Println(UpdateProductErr.Error())
		return false, UpdateProductErr
	}
	return true, nil
}

func DeleteProduct(productID int) (bool, error) {
	SqlQuery := (`DELETE FROM products WHERE id = ?`)
	_, DeleteProductErr := db.Exec(SqlQuery, productID)
	if DeleteProductErr != nil {
		log.Println(DeleteProductErr.Error())
		return false, DeleteProductErr
	}
	return true, nil
}
func UpdateUsersLogInTime(user_id int) (bool, error) {
	SqlQuery := (`UPDATE users SET login_time = NOW() WHERE id = ?`)
	_, UpdateLogInErr := db.Exec(SqlQuery, user_id)
	if UpdateLogInErr != nil {
		log.Println(UpdateLogInErr.Error())
		return false, UpdateLogInErr
	}
	return true, nil
}
func UpdateUsersLogOutTime(user_id int) (bool, error) {
	SqlQuery := (`UPDATE users SET logout_time = NOW() WHERE id = ?`)
	_, UpdateLogOutErr := db.Exec(SqlQuery, user_id)
	if UpdateLogOutErr != nil {
		log.Println(UpdateLogOutErr.Error())
		return false, UpdateLogOutErr
	}
	return true, nil
}
func GetSuperAdminByEmail(email string) ([]GetSuperAdminDetails, error) {
	var superadmins GetSuperAdminDetails
	superadmins_array := []GetSuperAdminDetails{}
	SqlQuery := (`SELECT id, name, email, phone, address FROM superadmins WHERE email = ?`)
	err := db.QueryRow(SqlQuery, email).Scan(&superadmins.SuperAdminID, &superadmins.Name, &superadmins.Email, &superadmins.Phone, &superadmins.Address)
	if err != nil && err == sql.ErrNoRows {
		return superadmins_array, err
	}
	superadmins_array = append(superadmins_array, superadmins)
	return superadmins_array, nil
}
func GetUserDetailsByEmail(email string) ([]GetAllUsersDetails, error) {
	var usersDetails GetAllUsersDetails
	usersDetails_array := []GetAllUsersDetails{}
	SqlQuery := (`SELECT id, name, email, phone, address FROM users WHERE email = ?`)
	err := db.QueryRow(SqlQuery, email).Scan(&usersDetails.UserID, &usersDetails.Name, &usersDetails.Email, &usersDetails.Phone, &usersDetails.Address)
	if err != nil && err == sql.ErrNoRows {
		return usersDetails_array, err
	}
	usersDetails_array = append(usersDetails_array, usersDetails)
	return usersDetails_array, nil
}
func GetSuperAdminPassword(email string) string {
	var password string
	SqlQuery := (`SELECT password FROM superadmins WHERE email = ?`)
	userPassword := db.QueryRow(SqlQuery, email).Scan(&password)
	if userPassword != nil {
		log.Println(userPassword.Error())
	}
	return password
}
func GetUsersPassword(email string) string {
	var password string
	SqlQuery := (`SELECT password FROM users WHERE email = ?`)
	userPassword := db.QueryRow(SqlQuery, email).Scan(&password)
	if userPassword != nil {
		log.Println(userPassword.Error())
	}
	return password
}
func VerifySuperAdminByEmail(email string) (bool, error) {
	var isAlreadyDB bool
	SqlQuery := (`SELECT true FROM superadmins WHERE email = ?`)
	err := db.QueryRow(SqlQuery, email).Scan(&isAlreadyDB)
	if err != nil && err == sql.ErrNoRows {
		return false, err
	}
	return isAlreadyDB, nil
}
func VerifyUsersByEmail(email string) (bool, error) {
	var isAlreadyDB bool
	SqlQuery := (`SELECT true FROM users WHERE email = ?`)
	err := db.QueryRow(SqlQuery, email).Scan(&isAlreadyDB)
	if err != nil && err == sql.ErrNoRows {
		return false, err
	}
	return isAlreadyDB, nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func GetUsersActivityBySuperAdmin() ([]GetUsersActivity, error) {
	var usersActivity GetUsersActivity
	var LoginStatus string
	usersActivity_array := []GetUsersActivity{}
	SqlQuery := (`SELECT users.id, users.name, users.email, users.phone, users.login_time, users.logout_time,
				products_audit.product_id, products_audit.name, products_audit.price, products_audit.image, products_audit.time, products_audit.action
				FROM users INNER JOIN
				products_audit ON users.id = products_audit.user_id ORDER BY products_audit.time ASC`)
	usersActivityRow, usersActivityErr := db.Query(SqlQuery)
	if usersActivityErr != nil {
		log.Println(usersActivityErr.Error())
		return usersActivity_array, usersActivityErr
	}
	for usersActivityRow.Next() {
		usersActivityScan := usersActivityRow.Scan(&usersActivity.UserID, &usersActivity.UserName, &usersActivity.UserEmail, &usersActivity.UserPhone, &usersActivity.UserLoginTime, &usersActivity.UserLogOutTime, &usersActivity.ProductID, &usersActivity.ProductName, &usersActivity.ProductPrice, &usersActivity.ProductImageLink, &usersActivity.ActionTime, &usersActivity.Action)

		if usersActivityScan != nil {
			log.Println(usersActivityScan.Error())
			return usersActivity_array, usersActivityScan
		}
		if (usersActivity.UserLoginTime == usersActivity.UserLogOutTime) || (usersActivity.UserLoginTime > usersActivity.UserLogOutTime) {
			LoginStatus = "Already log in"
			usersActivity.LoginStatus = LoginStatus
		} else {
			LoginStatus = "Already log out"
			usersActivity.LoginStatus = LoginStatus
		}
		usersActivity_array = append(usersActivity_array, usersActivity)
	}
	defer usersActivityRow.Close()
	return usersActivity_array, nil
}

func GetImageNameByProductID(productID string) string {
	var image_name string
	SqlQuery := (`SELECT image FROM products WHERE id = ?`)
	imageErr := db.QueryRow(SqlQuery, productID).Scan(&image_name)
	if imageErr != nil {
		log.Println(imageErr.Error())
	}
	return image_name
}
