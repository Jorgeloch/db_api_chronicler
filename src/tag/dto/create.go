package tagDTO

import "strings"

type CreateTagDTO struct {
	Nome string `validate:"required,max=20" json:"nome"`
	Cor  string `validate:"required,max=20" json:"cor"`
}

func (dto *CreateTagDTO) ValidateHexColor() bool {
	color := strings.ToLower(dto.Cor)
	return hexRegex.MatchString(color)
}
