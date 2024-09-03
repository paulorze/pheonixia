package customErrors

var IncompleteRegisterFields = CustomError{Code: 400, Message: "username, password, email, firstName and lastName are all required fields"}

var InvalidUsername = CustomError{Code: 422, Message: "invalid username"}

var ExistingEntry = CustomError{Code: 409, Message: "username or email already in use"}

var ExistingUsername = CustomError{Code: 409, Message: "username already in use"}

var InvalidPassword = CustomError{Code: 422, Message: "invalid password"}

var InvalidEmail = CustomError{Code: 422, Message: "invalid email"}

var ExistingEmail = CustomError{Code: 409, Message: "email already in use"}

var InvalidName = CustomError{Code: 422, Message: "invalid name"}

var UserNotFound = &CustomError{Code: 404, Message: "user not found"}

var InvalidAuthentication = CustomError{Code: 401, Message: "invalid username or password"}

var InvalidToken = CustomError{Code: 401, Message: "invalid token"}
