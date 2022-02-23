package navigator

import "fmt"

type NavigatorContext struct{}

func (thisObj *NavigatorContext) PrintRoute(routeStrategy IRouteStrategy) {
	var path string = routeStrategy.GetRoute()
	fmt.Println(path)
}
