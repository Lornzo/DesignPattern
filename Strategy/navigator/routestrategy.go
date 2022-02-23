package navigator

type WalkStrategy struct{}

func (thisObj *WalkStrategy) GetRoute() (path string) {
	path = "走路的路徑"
	return
}

type DriveStrategy struct{}

func (thisObj *DriveStrategy) GetRoute() (path string) {
	path = "開車的路徑"
	return
}

type BikeStrategy struct{}

func (thisObj *BikeStrategy) GetRoute() (path string) {
	path = "騎上心愛的小摩托，它從來不會堵車"
	return
}
