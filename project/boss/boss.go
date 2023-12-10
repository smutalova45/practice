package boss
import(
	
)



type Boss struct{
   Name string
   Budget int
   Profit int 
} 
func (b *Boss) NewBoss(name string,budget int , profit int) Boss {
	return Boss{
		Name: name,

		Budget: budget,
		Profit: profit,
	}
	
}
func (b *Boss)GetProductList(){

}
func (b *Boss)GetReport(){
	
}