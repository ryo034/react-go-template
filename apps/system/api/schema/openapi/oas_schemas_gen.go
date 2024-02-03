// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"github.com/google/uuid"
)

type APIV1AuthOtpPostOK struct {
	// OTP 6 digit code.
	Code string `json:"code"`
}

// GetCode returns the value of Code.
func (s *APIV1AuthOtpPostOK) GetCode() string {
	return s.Code
}

// SetCode sets the value of Code.
func (s *APIV1AuthOtpPostOK) SetCode(val string) {
	s.Code = val
}

func (*APIV1AuthOtpPostOK) aPIV1AuthOtpPostRes() {}

type APIV1AuthOtpPostReq struct {
	Email string `json:"email"`
}

// GetEmail returns the value of Email.
func (s *APIV1AuthOtpPostReq) GetEmail() string {
	return s.Email
}

// SetEmail sets the value of Email.
func (s *APIV1AuthOtpPostReq) SetEmail(val string) {
	s.Email = val
}

type APIV1AuthOtpVerifyPostOK struct {
	// JWT token.
	Token string `json:"token"`
}

// GetToken returns the value of Token.
func (s *APIV1AuthOtpVerifyPostOK) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *APIV1AuthOtpVerifyPostOK) SetToken(val string) {
	s.Token = val
}

func (*APIV1AuthOtpVerifyPostOK) aPIV1AuthOtpVerifyPostRes() {}

type APIV1AuthOtpVerifyPostReq struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

// GetEmail returns the value of Email.
func (s *APIV1AuthOtpVerifyPostReq) GetEmail() string {
	return s.Email
}

// GetOtp returns the value of Otp.
func (s *APIV1AuthOtpVerifyPostReq) GetOtp() string {
	return s.Otp
}

// SetEmail sets the value of Email.
func (s *APIV1AuthOtpVerifyPostReq) SetEmail(val string) {
	s.Email = val
}

// SetOtp sets the value of Otp.
func (s *APIV1AuthOtpVerifyPostReq) SetOtp(val string) {
	s.Otp = val
}

type APIV1PingGetOK struct {
	// Ping response message.
	Message OptString `json:"message"`
}

// GetMessage returns the value of Message.
func (s *APIV1PingGetOK) GetMessage() OptString {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *APIV1PingGetOK) SetMessage(val OptString) {
	s.Message = val
}

func (*APIV1PingGetOK) aPIV1PingGetRes() {}

type APIV1WorkspacesPostReq struct {
	// Workspace name.
	Name string `json:"name"`
	// Workspace subdomain (e.x. example-test).
	Subdomain string `json:"subdomain"`
}

// GetName returns the value of Name.
func (s *APIV1WorkspacesPostReq) GetName() string {
	return s.Name
}

// GetSubdomain returns the value of Subdomain.
func (s *APIV1WorkspacesPostReq) GetSubdomain() string {
	return s.Subdomain
}

// SetName sets the value of Name.
func (s *APIV1WorkspacesPostReq) SetName(val string) {
	s.Name = val
}

// SetSubdomain sets the value of Subdomain.
func (s *APIV1WorkspacesPostReq) SetSubdomain(val string) {
	s.Subdomain = val
}

// RFC7807 - https://datatracker.ietf.org/doc/html/rfc7807.
type BadRequestError struct {
	// The HTTP status code generated for this occurrence of the problem.
	Status OptInt `json:"status"`
	// Error type.
	Type OptString `json:"type"`
	// A short, human-readable summary of the problem type.
	Title OptString `json:"title"`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail OptString `json:"detail"`
	// Error code.
	Code OptString `json:"code"`
}

// GetStatus returns the value of Status.
func (s *BadRequestError) GetStatus() OptInt {
	return s.Status
}

// GetType returns the value of Type.
func (s *BadRequestError) GetType() OptString {
	return s.Type
}

// GetTitle returns the value of Title.
func (s *BadRequestError) GetTitle() OptString {
	return s.Title
}

// GetDetail returns the value of Detail.
func (s *BadRequestError) GetDetail() OptString {
	return s.Detail
}

// GetCode returns the value of Code.
func (s *BadRequestError) GetCode() OptString {
	return s.Code
}

// SetStatus sets the value of Status.
func (s *BadRequestError) SetStatus(val OptInt) {
	s.Status = val
}

// SetType sets the value of Type.
func (s *BadRequestError) SetType(val OptString) {
	s.Type = val
}

// SetTitle sets the value of Title.
func (s *BadRequestError) SetTitle(val OptString) {
	s.Title = val
}

// SetDetail sets the value of Detail.
func (s *BadRequestError) SetDetail(val OptString) {
	s.Detail = val
}

// SetCode sets the value of Code.
func (s *BadRequestError) SetCode(val OptString) {
	s.Code = val
}

func (*BadRequestError) aPIV1AuthOAuthPostRes()     {}
func (*BadRequestError) aPIV1AuthOtpPostRes()       {}
func (*BadRequestError) aPIV1AuthOtpVerifyPostRes() {}
func (*BadRequestError) aPIV1WorkspacesPostRes()    {}
func (*BadRequestError) updateNameRes()             {}

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

// RFC7807 - https://datatracker.ietf.org/doc/html/rfc7807.
type ConflictError struct {
	// The HTTP status code generated for this occurrence of the problem.
	Status OptInt `json:"status"`
	// Error type.
	Type OptString `json:"type"`
	// A short, human-readable summary of the problem type.
	Title OptString `json:"title"`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail OptString `json:"detail"`
	// Error code.
	Code OptString `json:"code"`
}

// GetStatus returns the value of Status.
func (s *ConflictError) GetStatus() OptInt {
	return s.Status
}

// GetType returns the value of Type.
func (s *ConflictError) GetType() OptString {
	return s.Type
}

// GetTitle returns the value of Title.
func (s *ConflictError) GetTitle() OptString {
	return s.Title
}

// GetDetail returns the value of Detail.
func (s *ConflictError) GetDetail() OptString {
	return s.Detail
}

// GetCode returns the value of Code.
func (s *ConflictError) GetCode() OptString {
	return s.Code
}

// SetStatus sets the value of Status.
func (s *ConflictError) SetStatus(val OptInt) {
	s.Status = val
}

// SetType sets the value of Type.
func (s *ConflictError) SetType(val OptString) {
	s.Type = val
}

// SetTitle sets the value of Title.
func (s *ConflictError) SetTitle(val OptString) {
	s.Title = val
}

// SetDetail sets the value of Detail.
func (s *ConflictError) SetDetail(val OptString) {
	s.Detail = val
}

// SetCode sets the value of Code.
func (s *ConflictError) SetCode(val OptString) {
	s.Code = val
}

func (*ConflictError) aPIV1WorkspacesPostRes() {}

// RFC7807 - https://datatracker.ietf.org/doc/html/rfc7807.
type InternalServerError struct {
	// The HTTP status code generated for this occurrence of the problem.
	Status OptInt `json:"status"`
	// Error type.
	Type OptString `json:"type"`
	// A short, human-readable summary of the problem type.
	Title OptString `json:"title"`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail OptString `json:"detail"`
	// Error code.
	Code OptString `json:"code"`
}

// GetStatus returns the value of Status.
func (s *InternalServerError) GetStatus() OptInt {
	return s.Status
}

// GetType returns the value of Type.
func (s *InternalServerError) GetType() OptString {
	return s.Type
}

// GetTitle returns the value of Title.
func (s *InternalServerError) GetTitle() OptString {
	return s.Title
}

// GetDetail returns the value of Detail.
func (s *InternalServerError) GetDetail() OptString {
	return s.Detail
}

// GetCode returns the value of Code.
func (s *InternalServerError) GetCode() OptString {
	return s.Code
}

// SetStatus sets the value of Status.
func (s *InternalServerError) SetStatus(val OptInt) {
	s.Status = val
}

// SetType sets the value of Type.
func (s *InternalServerError) SetType(val OptString) {
	s.Type = val
}

// SetTitle sets the value of Title.
func (s *InternalServerError) SetTitle(val OptString) {
	s.Title = val
}

// SetDetail sets the value of Detail.
func (s *InternalServerError) SetDetail(val OptString) {
	s.Detail = val
}

// SetCode sets the value of Code.
func (s *InternalServerError) SetCode(val OptString) {
	s.Code = val
}

func (*InternalServerError) aPIV1AuthOAuthPostRes()     {}
func (*InternalServerError) aPIV1AuthOtpPostRes()       {}
func (*InternalServerError) aPIV1AuthOtpVerifyPostRes() {}
func (*InternalServerError) aPIV1MeGetRes()             {}
func (*InternalServerError) aPIV1PingGetRes()           {}
func (*InternalServerError) aPIV1WorkspacesGetRes()     {}
func (*InternalServerError) aPIV1WorkspacesPostRes()    {}
func (*InternalServerError) loginRes()                  {}
func (*InternalServerError) updateNameRes()             {}

// Ref: #/components/schemas/Me
type Me struct {
	Self             User         `json:"self"`
	Member           OptMember    `json:"member"`
	CurrentWorkspace OptWorkspace `json:"currentWorkspace"`
	JoinedWorkspaces []Workspace  `json:"joinedWorkspaces"`
}

// GetSelf returns the value of Self.
func (s *Me) GetSelf() User {
	return s.Self
}

// GetMember returns the value of Member.
func (s *Me) GetMember() OptMember {
	return s.Member
}

// GetCurrentWorkspace returns the value of CurrentWorkspace.
func (s *Me) GetCurrentWorkspace() OptWorkspace {
	return s.CurrentWorkspace
}

// GetJoinedWorkspaces returns the value of JoinedWorkspaces.
func (s *Me) GetJoinedWorkspaces() []Workspace {
	return s.JoinedWorkspaces
}

// SetSelf sets the value of Self.
func (s *Me) SetSelf(val User) {
	s.Self = val
}

// SetMember sets the value of Member.
func (s *Me) SetMember(val OptMember) {
	s.Member = val
}

// SetCurrentWorkspace sets the value of CurrentWorkspace.
func (s *Me) SetCurrentWorkspace(val OptWorkspace) {
	s.CurrentWorkspace = val
}

// SetJoinedWorkspaces sets the value of JoinedWorkspaces.
func (s *Me) SetJoinedWorkspaces(val []Workspace) {
	s.JoinedWorkspaces = val
}

func (*Me) aPIV1MeGetRes() {}
func (*Me) loginRes()      {}
func (*Me) updateNameRes() {}

// Ref: #/components/schemas/Member
type Member struct {
	Profile MemberProfile `json:"profile"`
	User    User          `json:"user"`
}

// GetProfile returns the value of Profile.
func (s *Member) GetProfile() MemberProfile {
	return s.Profile
}

// GetUser returns the value of User.
func (s *Member) GetUser() User {
	return s.User
}

// SetProfile sets the value of Profile.
func (s *Member) SetProfile(val MemberProfile) {
	s.Profile = val
}

// SetUser sets the value of User.
func (s *Member) SetUser(val User) {
	s.User = val
}

// Ref: #/components/schemas/MemberProfile
type MemberProfile struct {
	// Base32 encoded UUID.
	ID          string    `json:"id"`
	DisplayName string    `json:"displayName"`
	IdNumber    OptString `json:"idNumber"`
}

// GetID returns the value of ID.
func (s *MemberProfile) GetID() string {
	return s.ID
}

// GetDisplayName returns the value of DisplayName.
func (s *MemberProfile) GetDisplayName() string {
	return s.DisplayName
}

// GetIdNumber returns the value of IdNumber.
func (s *MemberProfile) GetIdNumber() OptString {
	return s.IdNumber
}

// SetID sets the value of ID.
func (s *MemberProfile) SetID(val string) {
	s.ID = val
}

// SetDisplayName sets the value of DisplayName.
func (s *MemberProfile) SetDisplayName(val string) {
	s.DisplayName = val
}

// SetIdNumber sets the value of IdNumber.
func (s *MemberProfile) SetIdNumber(val OptString) {
	s.IdNumber = val
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMember returns new OptMember with value set to v.
func NewOptMember(v Member) OptMember {
	return OptMember{
		Value: v,
		Set:   true,
	}
}

// OptMember is optional Member.
type OptMember struct {
	Value Member
	Set   bool
}

// IsSet returns true if OptMember was set.
func (o OptMember) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMember) Reset() {
	var v Member
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMember) SetTo(v Member) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMember) Get() (v Member, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMember) Or(d Member) Member {
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

// NewOptWorkspace returns new OptWorkspace with value set to v.
func NewOptWorkspace(v Workspace) OptWorkspace {
	return OptWorkspace{
		Value: v,
		Set:   true,
	}
}

// OptWorkspace is optional Workspace.
type OptWorkspace struct {
	Value Workspace
	Set   bool
}

// IsSet returns true if OptWorkspace was set.
func (o OptWorkspace) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptWorkspace) Reset() {
	var v Workspace
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptWorkspace) SetTo(v Workspace) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptWorkspace) Get() (v Workspace, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptWorkspace) Or(d Workspace) Workspace {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// RFC7807 - https://datatracker.ietf.org/doc/html/rfc7807.
type TooManyRequestsError struct {
	// The HTTP status code generated for this occurrence of the problem.
	Status OptInt `json:"status"`
	// Error type.
	Type OptString `json:"type"`
	// A short, human-readable summary of the problem type.
	Title OptString `json:"title"`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail OptString `json:"detail"`
	// Error code.
	Code OptString `json:"code"`
}

// GetStatus returns the value of Status.
func (s *TooManyRequestsError) GetStatus() OptInt {
	return s.Status
}

// GetType returns the value of Type.
func (s *TooManyRequestsError) GetType() OptString {
	return s.Type
}

// GetTitle returns the value of Title.
func (s *TooManyRequestsError) GetTitle() OptString {
	return s.Title
}

// GetDetail returns the value of Detail.
func (s *TooManyRequestsError) GetDetail() OptString {
	return s.Detail
}

// GetCode returns the value of Code.
func (s *TooManyRequestsError) GetCode() OptString {
	return s.Code
}

// SetStatus sets the value of Status.
func (s *TooManyRequestsError) SetStatus(val OptInt) {
	s.Status = val
}

// SetType sets the value of Type.
func (s *TooManyRequestsError) SetType(val OptString) {
	s.Type = val
}

// SetTitle sets the value of Title.
func (s *TooManyRequestsError) SetTitle(val OptString) {
	s.Title = val
}

// SetDetail sets the value of Detail.
func (s *TooManyRequestsError) SetDetail(val OptString) {
	s.Detail = val
}

// SetCode sets the value of Code.
func (s *TooManyRequestsError) SetCode(val OptString) {
	s.Code = val
}

func (*TooManyRequestsError) aPIV1AuthOtpPostRes()       {}
func (*TooManyRequestsError) aPIV1AuthOtpVerifyPostRes() {}

type UpdateNameReq struct {
	// Name.
	Name string `json:"name"`
}

// GetName returns the value of Name.
func (s *UpdateNameReq) GetName() string {
	return s.Name
}

// SetName sets the value of Name.
func (s *UpdateNameReq) SetName(val string) {
	s.Name = val
}

// Ref: #/components/schemas/User
type User struct {
	UserId      uuid.UUID `json:"userId"`
	Email       string    `json:"email"`
	Name        OptString `json:"name"`
	PhoneNumber OptString `json:"phoneNumber"`
}

// GetUserId returns the value of UserId.
func (s *User) GetUserId() uuid.UUID {
	return s.UserId
}

// GetEmail returns the value of Email.
func (s *User) GetEmail() string {
	return s.Email
}

// GetName returns the value of Name.
func (s *User) GetName() OptString {
	return s.Name
}

// GetPhoneNumber returns the value of PhoneNumber.
func (s *User) GetPhoneNumber() OptString {
	return s.PhoneNumber
}

// SetUserId sets the value of UserId.
func (s *User) SetUserId(val uuid.UUID) {
	s.UserId = val
}

// SetEmail sets the value of Email.
func (s *User) SetEmail(val string) {
	s.Email = val
}

// SetName sets the value of Name.
func (s *User) SetName(val OptString) {
	s.Name = val
}

// SetPhoneNumber sets the value of PhoneNumber.
func (s *User) SetPhoneNumber(val OptString) {
	s.PhoneNumber = val
}

func (*User) aPIV1AuthOAuthPostRes() {}

// Ref: #/components/schemas/Workspace
type Workspace struct {
	// Base32 encoded UUID.
	WorkspaceId string `json:"workspaceId"`
	// Workspace name.
	Name string `json:"name"`
	// Workspace subdomain (e.x. example-test).
	Subdomain string `json:"subdomain"`
}

// GetWorkspaceId returns the value of WorkspaceId.
func (s *Workspace) GetWorkspaceId() string {
	return s.WorkspaceId
}

// GetName returns the value of Name.
func (s *Workspace) GetName() string {
	return s.Name
}

// GetSubdomain returns the value of Subdomain.
func (s *Workspace) GetSubdomain() string {
	return s.Subdomain
}

// SetWorkspaceId sets the value of WorkspaceId.
func (s *Workspace) SetWorkspaceId(val string) {
	s.WorkspaceId = val
}

// SetName sets the value of Name.
func (s *Workspace) SetName(val string) {
	s.Name = val
}

// SetSubdomain sets the value of Subdomain.
func (s *Workspace) SetSubdomain(val string) {
	s.Subdomain = val
}

func (*Workspace) aPIV1WorkspacesPostRes() {}

type Workspaces []Workspace

func (*Workspaces) aPIV1WorkspacesGetRes() {}
