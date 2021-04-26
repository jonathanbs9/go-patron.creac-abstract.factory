package main

type Adidas struct {
}

func (a *Adidas) MakeShoe() IShoe {
	return &adidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 10,
		},
	}
}

func (a *Adidas) MakeShirt() IShirt {
	return &adidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 29,
		},
	}
}
