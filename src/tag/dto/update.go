package tagDTO

import "strings"

type UpdateTagDTO struct {
	Nome string `validate:"max=20" json:"nome,omitempty"`
	Cor  string `validate:"max=20" json:"cor,omitempty"`
}

func (dto *UpdateTagDTO) ValidateHexColor() bool {
	color := strings.ToLower(dto.Cor)
	return hexRegex.MatchString(color)
}
