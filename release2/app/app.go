package app

import (
	"fmt"

	"p1-lc03.com/repo"
)

func Run() {

	for {
		fmt.Println("Select Report To Generate")
		fmt.Println("1. Total Games Sales Report")
		fmt.Println("2. Most Popular Game Report")
		fmt.Println("3. Total Revenue Per Game Report")
		fmt.Println("4. Player Count Per Game Report")
		fmt.Println("5. Exit")

		var choice int

		fmt.Print("Enter the number of the report you want to generate : ")
		fmt.Scan(&choice)


		fmt.Println()
		switch choice{
		case 1:
			fmt.Println("Total Games Sales Report")

			data,_ := repo.GetTotalGamesSalesReport()

			for _,v := range data{
				fmt.Printf("Game : %s, Total Sales : %d\n",v.Name,int(v.Total))
			}
		case 2:
			fmt.Println("Most Popular Game Report")
			data,_ := repo.GetMostPopularGameReport()

			for _,v := range data{
				fmt.Printf("Game : %s, Unique Player : %d\n",v.Name,v.Total)
			}
		case 3:
			fmt.Println("Total Revenue Per Game Report")
			data,_ := repo.GetTotalRevenuePerGameReport()
			for _,v := range data{
				fmt.Printf("Game : %s, Total Revenue : $%.2f\n",v.Name,v.Total)
			}
		case 4:
			fmt.Println("Player Count Per Game Report")
			data,_ := repo.GetPlayerCountPerGameReport()
			for _,v := range data{
				fmt.Printf("Game : %s, Total Player : %d\n",v.Name,v.Total)
			}
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again")
			
		}
		


	}
}