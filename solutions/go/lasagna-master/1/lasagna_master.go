package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, one int) (time int){
    if one !=0 {
    time= len(layers) * one    
    }else{
    time= len(layers) * 2
    }
    return time
}

// TODO: define the 'Quantities()' function
func Quantities(quant []string ) (noodles int , sauce float64){

    for i:=0;i<len(quant);i++{
        if quant[i] == "sauce"{
            sauce+=0.2
        }else if quant[i] == "noodles"{
            noodles+=50
        }
     }
    return noodles, sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList[]string, myList[]string){
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(quantities []float64, scale int) (scaledQuant []float64) {
	for i := 0; i < len(quantities); i++ {
		scaledQuant= append(scaledQuant, quantities[i] * (float64(scale) / 2))
    }

	return scaledQuant
}
