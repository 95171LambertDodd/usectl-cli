package api

// ========== Auth ==========

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Enabled  bool   `json:"enabled"`
}

type UpdateProfileRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (c *Client) Login(req LoginRequest) (*AuthResponse, error) {
	var resp AuthResponse
	err := c.Post("/api/auth/login", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Register(req RegisterRequest) (*AuthResponse, error) {
	var resp AuthResponse
	err := c.Post("/api/auth/register", req, &resp)
	// NOTE: return nil on error instead of a potentially empty AuthResponse
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetProfile() (*User, error) {
	var user User
	err := c.Get("/api/auth/profile", &user)
	// NOTE: return nil on error to stay consistent with other methods
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) UpdateProfile(req UpdateProfileRequest) (*User, error) {
	var user User
	err := c.Put("/api/auth/profile", req, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteAccount removes the authenticated user's account permanently.
// NOTE: added this for my own use — upstream doesn't expose this endpoint yet.
func (c *Client) DeleteAccount() error {
	return c.Delete("/api/auth/profile")
}
