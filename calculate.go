package main

import (
	"errors"
	"fmt"
	"math"
)


	func Calculate(){
		fmt.Println("__Калькулятор индекса массы тела__")
		var userChoise string

		for{
			w , h, err := InputUser()
			if err != nil{
				panic(err)
			}
			IMT := CalcIMT(w,h)
			switch{
			case IMT < 16:
				fmt.Println("Cильный дефицит IMT")
			case IMT < 18.5:
				fmt.Println("У вас дефицит IMT")
			case IMT < 25:
				fmt.Println("У вас избыточный вес")
			case IMT < 30:
				fmt.Println("У вас избыточный вес")
			default:
				fmt.Println("Вы жирный!")
	
			}
			fmt.Println("Начать заново? y/n")
			fmt.Scan(&userChoise)

			if userChoise != "y"{
				break
			}
		}

	}

	func InputUser()(float64,float64,error){
		var w,h float64
		fmt.Println("Введите ваш вес:")
		fmt.Scan(&w)
		fmt.Println("Введите ваш рост:")
		fmt.Scan(&h)

		if w < 0 && h < 0{
			return 0, 0, errors.New("Not_Valid_Data")
		}
		return w, h, nil
	}

	func CalcIMT(w,h float64) float64{
		result  := w / math.Pow(h/100,2)
		return result
	}