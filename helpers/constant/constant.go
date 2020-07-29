package constant

const (
	CodeSuccess        = 200
	CodeFail           = 400
	MaxLenRegisterCode = 32
	TypeVoteLike       = 1
	TypeVoteDisLike    = 2
	TimeLayout         = "2006-01-02"

	// Message error validate input
	MaxLengthEmail             = 256
	MinLengthEmail             = 10
	MessageErrorEmail001       = "Email cannot be empty"
	MessageErrorEmail002       = "Email address in invalid format"
	MessageErrorEmail003       = "Email must be between 10 and 256 characters long"
	MessageErrorEmail004       = "Email address already exists"
	MaxLengthPassword          = 32
	MinLengthPassword          = 8
	MessageErrorPassword001    = "Password cannot be empty"
	MessageErrorPassword002    = "Password must be between 8 and 32 characters long"
	MaxLengthFirstName         = 36
	MinLengthFirstName         = 1
	MessageErrorFirstName001   = "First name cannot be empty"
	MessageErrorFirstName002   = "First name must be between 1 and 36 characters long"
	MessageErrorFirstName003   = "First name cannot contain special character (!@#$%^&*(),._?:{+}|<>/-)."
	MaxLengthLastName          = 36
	MinLengthLastName          = 1
	MessageErrorLastName001    = "Last name cannot be empty"
	MessageErrorLastName002    = "Last name must be between 1 and 36 characters long"
	MessageErrorLastName003    = "Last name cannot contain special character (!@#$%^&*(),._?:{+}|<>/-)."
	MaxLengthPhoneNumber       = 11
	MinLengthPhoneNumber       = 10
	MessageErrorPhoneNumber001 = "Phone number cannot be empty"
	MessageErrorPhoneNumber002 = "Phone number is invalid"
	MessageErrorPhoneNumber003 = "Phone number must be between 10 and 11 characters long"
	MaxLengthAddress           = 256
	MinLengthAddress           = 1
	MessageErrorAddress001     = "Address cannot be empty"
	MessageErrorAddress002     = "Phone number must be between 1 and 256 characters long"
	MessageErrorBirthday001    = "Date of birth cannot be empty"
	MessageErrorBirthday002    = "Date of birth is invalid"

	// Message call API error
	MessageError001 = "Convert data input fail"
	MessageError002 = "QuestionId cannot be empty"
	MessageError003 = "TagId cannot be empty"
	MessageError004 = "Question not exist"
	MessageError005 = "Answer not exist"
	// Message call API successful
	MessageSuccess001 = "Create user successful"
	MessageSuccess002 = "User login successful"
	MessageSuccess003 = "Vote your successful question"
	MessageSuccess004 = "Vote your successful answer"
	MessageSuccess005 = "Create tag successful"
	MessageSuccess006 = "Get list tag successful"
	MessageSuccess007 = "Create question successful"
	MessageSuccess008 = "Get question successful"
	MessageSuccess009 = "Get list question by tagId successful"
	MessageSuccess010 = "Create answer successful"
	MessageSuccess011 = "Get answer by questionId successful"
	MessageSuccess012 = "List answer by questionId empty"
	MessageSuccess013 = "List question by tagId empty"
)
