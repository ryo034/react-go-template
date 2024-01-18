// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'l': // Prefix: "login"
				if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handleLoginRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}
			case 'm': // Prefix: "me"
				if l := len("me"); len(elem) >= l && elem[0:l] == "me" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleMeGetRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'p': // Prefix: "ping"
				if l := len("ping"); len(elem) >= l && elem[0:l] == "ping" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handlePingGetRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handlePingPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
			case 's': // Prefix: "sign_up"
				if l := len("sign_up"); len(elem) >= l && elem[0:l] == "sign_up" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handleSignUpRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [0]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'l': // Prefix: "login"
				if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "POST":
						// Leaf: Login
						r.name = "Login"
						r.summary = "Login"
						r.operationID = "login"
						r.pathPattern = "/login"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'm': // Prefix: "me"
				if l := len("me"); len(elem) >= l && elem[0:l] == "me" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: MeGet
						r.name = "MeGet"
						r.summary = "Get Admin User"
						r.operationID = ""
						r.pathPattern = "/me"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'p': // Prefix: "ping"
				if l := len("ping"); len(elem) >= l && elem[0:l] == "ping" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: PingGet
						r.name = "PingGet"
						r.summary = "Checks if the server is running"
						r.operationID = ""
						r.pathPattern = "/ping"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						// Leaf: PingPost
						r.name = "PingPost"
						r.summary = "Checks if the server is running"
						r.operationID = ""
						r.pathPattern = "/ping"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 's': // Prefix: "sign_up"
				if l := len("sign_up"); len(elem) >= l && elem[0:l] == "sign_up" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "POST":
						// Leaf: SignUp
						r.name = "SignUp"
						r.summary = "Sign Up"
						r.operationID = "sign_up"
						r.pathPattern = "/sign_up"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			}
		}
	}
	return r, false
}
