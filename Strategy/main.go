package main

import "github.com/Lornzo/DesignPattern/Strategy/navigator"

func main() {
	var nav navigator.NavigatorContext = navigator.NavigatorContext{}

	var walkStrategy navigator.WalkStrategy = navigator.WalkStrategy{}
	nav.PrintRoute(&walkStrategy)

	var driveStrategy navigator.DriveStrategy = navigator.DriveStrategy{}
	nav.PrintRoute(&driveStrategy)

	var bikeStrategy navigator.BikeStrategy = navigator.BikeStrategy{}
	nav.PrintRoute(&bikeStrategy)

}
