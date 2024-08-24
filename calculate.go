package main

import (
	"fmt"
	"math"
)


	func Calculate(){
		fmt.Println("__Калькулятор индекса массы тела__")
		w , h := InputUser()
		IMT := CalcIMT(w,h)

		if IMT < 16{
			fmt.Println("Cильный дефицит IMT")
		}else if IMT <= 16 && IMT < 18.5{
			fmt.Println("У вас дефицит IMT")
		}else if IMT >= 18.5 && IMT <25{
			fmt.Println("У вас нормальный вес")
		}else if IMT >= 25 && IMT < 30{
			fmt.Println("У вас избыточный вес")
		}else{
			fmt.Println("Вы жирный!")
		}


	}

	func InputUser()(float64,float64){
		var w,h float64
		fmt.Println("Введите ваш вес:")
		fmt.Scan(&w)
		fmt.Println("Введите ваш рост:")
		fmt.Scan(&h)
		return w,h
	}

	func CalcIMT(w,h float64) float64{
		result  := w / math.Pow(h/100,2)
		return result
	}