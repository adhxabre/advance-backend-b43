package authdto

// Declare AuthRequest struct here ...
type AuthRequest struct {
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
}
