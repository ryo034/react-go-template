// Code generated by ogen, DO NOT EDIT.

package openapi

type BadRequest struct {
	Code    OptInt32  `json:"code"`
	Message OptString `json:"message"`
}

// GetCode returns the value of Code.
func (s *BadRequest) GetCode() OptInt32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *BadRequest) GetMessage() OptString {
	return s.Message
}

// SetCode sets the value of Code.
func (s *BadRequest) SetCode(val OptInt32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *BadRequest) SetMessage(val OptString) {
	s.Message = val
}

func (*BadRequest) signUpRes() {}

type Bearer struct {
	Token string
}

// GetToken returns the value of Token.
func (s *Bearer) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *Bearer) SetToken(val string) {
	s.Token = val
}

type InternalServerError struct {
	Code    OptInt32  `json:"code"`
	Message OptString `json:"message"`
}

// GetCode returns the value of Code.
func (s *InternalServerError) GetCode() OptInt32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *InternalServerError) GetMessage() OptString {
	return s.Message
}

// SetCode sets the value of Code.
func (s *InternalServerError) SetCode(val OptInt32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *InternalServerError) SetMessage(val OptString) {
	s.Message = val
}

func (*InternalServerError) loginRes()   {}
func (*InternalServerError) meGetRes()   {}
func (*InternalServerError) pingGetRes() {}
func (*InternalServerError) signUpRes()  {}

// Ref: #/components/schemas/Me
type Me struct {
	EmailVerified bool           `json:"emailVerified"`
	MultiFactor   OptMultiFactor `json:"multiFactor"`
	Member        Member         `json:"member"`
}

// GetEmailVerified returns the value of EmailVerified.
func (s *Me) GetEmailVerified() bool {
	return s.EmailVerified
}

// GetMultiFactor returns the value of MultiFactor.
func (s *Me) GetMultiFactor() OptMultiFactor {
	return s.MultiFactor
}

// GetMember returns the value of Member.
func (s *Me) GetMember() Member {
	return s.Member
}

// SetEmailVerified sets the value of EmailVerified.
func (s *Me) SetEmailVerified(val bool) {
	s.EmailVerified = val
}

// SetMultiFactor sets the value of MultiFactor.
func (s *Me) SetMultiFactor(val OptMultiFactor) {
	s.MultiFactor = val
}

// SetMember sets the value of Member.
func (s *Me) SetMember(val Member) {
	s.Member = val
}

func (*Me) loginRes()  {}
func (*Me) meGetRes()  {}
func (*Me) signUpRes() {}

// Ref: #/components/schemas/Member
type Member struct {
	IdNumber OptString `json:"idNumber"`
	User     User      `json:"user"`
}

// GetIdNumber returns the value of IdNumber.
func (s *Member) GetIdNumber() OptString {
	return s.IdNumber
}

// GetUser returns the value of User.
func (s *Member) GetUser() User {
	return s.User
}

// SetIdNumber sets the value of IdNumber.
func (s *Member) SetIdNumber(val OptString) {
	s.IdNumber = val
}

// SetUser sets the value of User.
func (s *Member) SetUser(val User) {
	s.User = val
}

// Ref: #/components/schemas/MultiFactor
type MultiFactor struct {
	FactorId    string `json:"factorId"`
	PhoneNumber string `json:"phoneNumber"`
}

// GetFactorId returns the value of FactorId.
func (s *MultiFactor) GetFactorId() string {
	return s.FactorId
}

// GetPhoneNumber returns the value of PhoneNumber.
func (s *MultiFactor) GetPhoneNumber() string {
	return s.PhoneNumber
}

// SetFactorId sets the value of FactorId.
func (s *MultiFactor) SetFactorId(val string) {
	s.FactorId = val
}

// SetPhoneNumber sets the value of PhoneNumber.
func (s *MultiFactor) SetPhoneNumber(val string) {
	s.PhoneNumber = val
}

// NewOptInt32 returns new OptInt32 with value set to v.
func NewOptInt32(v int32) OptInt32 {
	return OptInt32{
		Value: v,
		Set:   true,
	}
}

// OptInt32 is optional int32.
type OptInt32 struct {
	Value int32
	Set   bool
}

// IsSet returns true if OptInt32 was set.
func (o OptInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt32) SetTo(v int32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt32) Get() (v int32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt32) Or(d int32) int32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMultiFactor returns new OptMultiFactor with value set to v.
func NewOptMultiFactor(v MultiFactor) OptMultiFactor {
	return OptMultiFactor{
		Value: v,
		Set:   true,
	}
}

// OptMultiFactor is optional MultiFactor.
type OptMultiFactor struct {
	Value MultiFactor
	Set   bool
}

// IsSet returns true if OptMultiFactor was set.
func (o OptMultiFactor) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMultiFactor) Reset() {
	var v MultiFactor
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMultiFactor) SetTo(v MultiFactor) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMultiFactor) Get() (v MultiFactor, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMultiFactor) Or(d MultiFactor) MultiFactor {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSignUpReq returns new OptSignUpReq with value set to v.
func NewOptSignUpReq(v SignUpReq) OptSignUpReq {
	return OptSignUpReq{
		Value: v,
		Set:   true,
	}
}

// OptSignUpReq is optional SignUpReq.
type OptSignUpReq struct {
	Value SignUpReq
	Set   bool
}

// IsSet returns true if OptSignUpReq was set.
func (o OptSignUpReq) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSignUpReq) Reset() {
	var v SignUpReq
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSignUpReq) SetTo(v SignUpReq) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSignUpReq) Get() (v SignUpReq, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSignUpReq) Or(d SignUpReq) SignUpReq {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// PingGetOK is response for PingGet operation.
type PingGetOK struct{}

func (*PingGetOK) pingGetRes() {}

type SignUpReq struct {
	// First Name.
	FirstName OptString `json:"first_name"`
	// Last Name.
	LastName OptString `json:"last_name"`
}

// GetFirstName returns the value of FirstName.
func (s *SignUpReq) GetFirstName() OptString {
	return s.FirstName
}

// GetLastName returns the value of LastName.
func (s *SignUpReq) GetLastName() OptString {
	return s.LastName
}

// SetFirstName sets the value of FirstName.
func (s *SignUpReq) SetFirstName(val OptString) {
	s.FirstName = val
}

// SetLastName sets the value of LastName.
func (s *SignUpReq) SetLastName(val OptString) {
	s.LastName = val
}

type Unauthorized struct {
	Code    OptInt32  `json:"code"`
	Message OptString `json:"message"`
}

// GetCode returns the value of Code.
func (s *Unauthorized) GetCode() OptInt32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Unauthorized) GetMessage() OptString {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Unauthorized) SetCode(val OptInt32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Unauthorized) SetMessage(val OptString) {
	s.Message = val
}

func (*Unauthorized) signUpRes() {}

// Ref: #/components/schemas/User
type User struct {
	UserId      string    `json:"userId"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	PhoneNumber OptString `json:"phoneNumber"`
}

// GetUserId returns the value of UserId.
func (s *User) GetUserId() string {
	return s.UserId
}

// GetEmail returns the value of Email.
func (s *User) GetEmail() string {
	return s.Email
}

// GetName returns the value of Name.
func (s *User) GetName() string {
	return s.Name
}

// GetPhoneNumber returns the value of PhoneNumber.
func (s *User) GetPhoneNumber() OptString {
	return s.PhoneNumber
}

// SetUserId sets the value of UserId.
func (s *User) SetUserId(val string) {
	s.UserId = val
}

// SetEmail sets the value of Email.
func (s *User) SetEmail(val string) {
	s.Email = val
}

// SetName sets the value of Name.
func (s *User) SetName(val string) {
	s.Name = val
}

// SetPhoneNumber sets the value of PhoneNumber.
func (s *User) SetPhoneNumber(val OptString) {
	s.PhoneNumber = val
}
