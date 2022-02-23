package navigator

type IRouteStrategy interface {
	GetRoute() (path string)
}
