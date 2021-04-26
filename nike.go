package main

type Nike struct {
}

func (n *Nike) MakeShoe() IShoe {
	return &nikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *Nike) MakeShirt() IShirt {
	return &nikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 14,
		},
	}
}
