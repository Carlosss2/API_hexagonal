package domain

type Product struct {
	Id int32
 	Name  string  `json:"name"`
    Price float32 `json:"price"`
}

func NewProduct(name string,price float32) *Product {
	
	return &Product{Id:1,Name:name, Price: price}
}


func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string){
	p.Name = name
}