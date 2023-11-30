package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "main.go/docs"
)

var (
	db  *sql.DB
	err error
)

const (
	mysqlHost     = "localhost"
	mysqlPort     = 3306
	mysqlUser     = "root"
	mysqlPassword = "passwordRoot"
	mysqlDB       = "TEST"

	postgresHost     = "localhost"
	postgresPort     = 5433
	postgresUser     = "root"
	postgresPassword = "rajandrO1"
	postgresDB       = "accuknox"
)

func MysqlConnection() error {
	mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDB)
	db, err = sql.Open("mysql", mysqlConnStr)
	if err != nil {
		log.Printf("Failed to connect to MySQL: %v\n", err)
		return err
	}
	// Test the MySQL connection
	err = db.Ping()
	if err != nil {
		log.Printf("Failed to ping MySQL: %v\n", err)
		return err
	}

	fmt.Println("Successfully connected to the MySQL database!")
	return nil
}
func testInterface(a string, b ...string) {
	sqlQuery := `select ` + b[0] + ``
	fmt.Println(sqlQuery, "::::: ", b[0])

}
func PostGresConnection() error {
	postgresConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassword, postgresDB)
	postgresDB, err := sql.Open("postgres", postgresConnStr)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL: ", err)
		return err
	}
	defer postgresDB.Close()

	// Test the PostgreSQL connection
	err = postgresDB.Ping()
	if err != nil {
		log.Fatal("Failed to ping PostgreSQL: ", err)
		return err
	}

	fmt.Println("Successfully connected to the PostgreSQL database!")
	return nil
}
func initDB() {
	if errMysql := MysqlConnection(); errMysql != nil {
		log.Println("Mysql connection failed: Trying to postgres connection...: ")
		if errPostgres := PostGresConnection(); errPostgres != nil {
			log.Println("Both connection is failed: ")
		}
	}
}

type InputStruct struct {
	Title  string  `json:"title"`
	Title1 *string `json:"title1,omitempty"`
}
type ResonseStruct struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// @Summary Create a book
// @Description Create a new book and return its details
// @Tags books
// @Accept json
// @Produce json
// @Param input body InputStruct true "Book details in JSON format"
// @Success 200 {object} ResonseStruct
// @Router / [post]
// @
func createBooks(c *gin.Context) {
	var input InputStruct
	if err := c.ShouldBindJSON(&input); err != nil { // Note the use of '&input'
		log.Println("createBooks ShouldBindJSON Err: ", err.Error())
		// Handle the error, perhaps by returning an error response to the client
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postSqlQuery := `INSERT INTO books(title)VALUES(?);`
	result, err := db.Exec(postSqlQuery, input.Title)
	if err != nil {
		log.Println("createBooks Err: ", err.Error())
	}
	// Get the last insert ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Println("createBooks lastInsertID ", err.Error())
	}

	var r ResonseStruct
	getSqlQuery := `SELECT id, title FROM books WHERE id = ?`
	if QueryRowErr := db.QueryRow(getSqlQuery, lastInsertID).Scan(&r.ID, &r.Title); QueryRowErr != nil {
		log.Println("createBooks QueryRow: ", QueryRowErr.Error())
	}
	// defer db.Close()
	c.JSON(http.StatusOK, r)
}
func GetBooks(c *gin.Context) {
	id := c.Request.FormValue("id")
	var r ResonseStruct
	getSqlQuery := `SELECT id, title FROM books WHERE id = ?`
	if QueryRowErr := db.QueryRow(getSqlQuery, id).Scan(&r.ID, &r.Title); QueryRowErr != nil {
		// e := CustomeError{Message: "Cannot divide by zero", Code: 12}
		// customlog.Error1("QueryRowErr := db.QueryRow")
		// log.Println("createBooks QueryRow: ", QueryRowErr.Error(), e.Error2())
		log.Println("Error")

	}
	// defer db.Close()
	c.JSON(http.StatusOK, r)
}

// func Error2() {
// 	panic("unimplemented")
// }

type CustomeError struct {
	Message string
	Code    int
}

func (cErr CustomeError) Error2() string {
	var e = cErr.Message + strconv.Itoa(cErr.Code)
	return e
}

func main() {
	// initDB()
	MysqlConnection()
	defer db.Close()
	// handelRoutes()
	DynamicFunction()
}

// func main() {
// initDB()
// MysqlConnection()
// defer db.Close()
// GetBooks(&gin.Context{})
// DynamicFunction()

// }
func DynamicFunction() {
	// Specify the columns you want in the result
	columns := []string{"id", "title"}

	// Build the SELECT query
	query := "SELECT " + joinColumns(columns) + " FROM books"
	// Execute the SQL query
	// query := `SELECT id,title FROM books`
	log.Println("------------1", query)

	rows, err := db.Query(query)

	log.Println("------------2")

	if err != nil {
		log.Println("Step1", err)
	}
	defer rows.Close()

	// Get column names from the result set
	columnNames, err := rows.Columns()
	if err != nil {
		log.Println("Step2", err.Error())
	}
	fmt.Println(columnNames)

	// Create a map to store the values dynamically
	values := make(map[string]interface{})

	// Prepare a slice of empty interfaces to scan into
	scanArgs := make([]interface{}, len(columnNames))
	for i := range columnNames {
		scanArgs[i] = new(interface{})
	}

	// Iterate over the rows
	for rows.Next() {
		// Scan the values into the scanArgs slice
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
		// Build a map of column names and values
		for i, col := range columnNames {
			values[col] = *(scanArgs[i].(*interface{}))
		}

		// Create a struct dynamically based on the column names
		resultStruct := createStruct(values)

		// Print the result
		fmt.Printf("%+v\n", resultStruct)
		printStruct(resultStruct)
	}
}

// joinColumns joins column names with commas
func joinColumns(columns []string) string {
	return "'" + strings.Join(columns, "','") + "'"
}

// createStruct dynamically creates a struct based on the column names and values
func createStruct(values map[string]interface{}) interface{} {
	// Define a new type dynamically
	fields := createFields(values)
	newType := reflect.StructOf(fields)

	// Create an instance of the new type
	newStruct := reflect.New(newType).Elem()

	// Set values for each exported field in the struct
	for key, value := range values {
		fieldName := strings.Title(key) // Capitalize the first letter
		field := newStruct.FieldByName(fieldName)
		if field.IsValid() && field.CanSet() {
			// Convert []byte to string if the field type is string
			if field.Type().Kind() == reflect.String {
				field.SetString(string(value.([]byte)))
			} else {
				field.Set(reflect.ValueOf(value))
			}
		}
	}

	return newStruct.Interface()
}

// createFields creates a slice of StructField based on the column names
func createFields(values map[string]interface{}) []reflect.StructField {
	var fields []reflect.StructField
	for key := range values {
		fields = append(fields, reflect.StructField{
			Name: strings.Title(key), // Capitalize the first letter
			Type: reflect.TypeOf(values[key]),
		})
	}
	return fields
}

// printStruct prints the struct in a specific format
func printStruct(s interface{}) {
	// Convert the struct to a string
	resultString := structToString(s)
	fmt.Println(resultString)
}

// structToString converts the struct to a formatted string
func structToString(s interface{}) string {
	var result string

	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Struct {
		// Iterate over the fields and build the string representation
		for i := 0; i < v.NumField(); i++ {
			result += fmt.Sprintf("%s: %v, ", v.Type().Field(i).Name, v.Field(i))
		}
		// Remove the trailing comma and space
		result = strings.TrimSuffix(result, ", ")
	}

	return result
}
func handelRoutes() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.POST("/", createBooks)
	r.GET("/", GetBooks)
	r.Run(":8080")
}
