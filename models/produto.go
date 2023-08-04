package models

import (
	"crud_golang/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodos() []Produto {
	db := db.ConectaBanco()

	selectProdutos, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	defer db.Close()

	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	banco := db.ConectaBanco()

	insereDados, err := banco.Prepare("insert into produtos(nome,descricao,preco,quantidade) values($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}

	insereDados.Exec(nome, descricao, preco, quantidade)
	defer banco.Close()
}

func DeletaProduto(id string) {
	banco := db.ConectaBanco()

	deletaProd, err := banco.Prepare("delete from produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}

	deletaProd.Exec(id)

	defer banco.Close()
}
