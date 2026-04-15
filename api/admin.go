package api

import "fmt"

// ========== Admin ==========

type SetEnabledRequest struct {
	Enabled bool `json:"enabled"`
}

type SetRoleRequest struct {
	Role string `json:"role"`
}

// ListUsers returns all users in the system.
// Note: requires admin privileges on the authenticated account.
func (c *Client) ListUsers() ([]User, error) {
	var users []User
	err := c.Get("/api/admin/users", &users)
	return users, err
}

func (c *Client) SetUserEnabled(id string, enabled bool) error {
	return c.Put(fmt.Sprintf("/api/admin/users/%s/enabled", id), SetEnabledRequest{Enabled: enabled}, nil)
}

func (c *Client) SetUserRole(id string, role string) error {
	return c.Put(fmt.Sprintf("/api/admin/users/%s/role", id), SetRoleRequest{Role: role}, nil)
}

// DeleteUser permanently removes a user by ID.
// TODO: consider adding a soft-delete option upstream
func (c *Client) DeleteUser(id string) error {
	return c.Delete(fmt.Sprintf("/api/admin/users/%s", id), nil)
}
