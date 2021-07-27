package models

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuário usando o DevBook
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoem,omitempty"`
}

// Preparar validar e formata dados do usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome do usuário é obrigatório e não pode estar em branco!")
	}

	if usuario.Nick == "" {
		return errors.New("O nick do usuário é obrigatório e não pode estar em branco!")
	}

	if usuario.Email == "" {
		return errors.New("O Email do usuário é obrigatório e não pode estar em branco!")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A Senha do usuário é obrigatório e não pode estar em branco!")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
