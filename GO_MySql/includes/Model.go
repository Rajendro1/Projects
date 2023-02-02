package includes

type (
	GetAllUsersDetails struct {
		UserID  int    `json:"user_id"`
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}
	GetSuperAdminDetails struct {
		SuperAdminID int    `json:"super_admin_id"`
		Name         string `json:"name"`
		Email        string `json:"email"`
		Phone        string `json:"phone"`
		Address      string `json:"address"`
	}
	GetUsersProductDetails struct {
		ProductID        string  `json:"product_id"`
		UserID           int     `json:"user_id"`
		ProductName      string  `json:"product_name"`
		ProductPrice     float64 `json:"product_price"`
		ProductImageLink string  `json:"product_image_link"`
	}
	GetUsersActivity struct {
		UserID           int    `json:"user_id"`
		UserName         string `json:"user_name"`
		UserEmail        string `json:"user_email"`
		UserPhone        string `json:"user_phone"`
		UserLoginTime    string `json:"user_login_time"`
		UserLogOutTime   string `json:"user_logout_time"`
		ProductID        string `json:"product_id"`
		ProductName      string `json:"product_name"`
		ProductPrice     string `json:"product_price"`
		ProductImageLink string `json:"product_image_link"`
		ActionTime       string `json:"action_time"`
		Action           string `json:"action"`
		LoginStatus      string `json:"login_status"`
	}
)
