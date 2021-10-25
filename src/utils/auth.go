package utils

// nambahin
// func Auth(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		if err := next(c); err != nil {
// 			c.Error(err)
// 		}
// 		email := c.FormValue("email")
// 		password := c.FormValue("password")

// 		// data, err := models.UserSearchEmailAuth(email)
// 		if err != nil {
// 			return err
// 		}
// 		if subtle.ConstantTimeCompare([]byte(email), []byte(data.Email)) == 1 &&
// 			subtle.ConstantTimeCompare([]byte(password), []byte(data.Password)) == 1 {
// 			return nil
// 		}
// 		return errors.New("password tidak cocok")
// 	}
// }
