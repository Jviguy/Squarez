package Squarez

type Dick struct {
	leftb *Circle
	shaft *RECT
	rightb *Circle
}

func (d *Dick) GetShaft() *RECT {
	return d.shaft
}

func (d *Dick) GetBall(bn int) *Circle {
	switch bn {
	case 1:
		return d.leftb
	case 2:
		return d.rightb
	}
	return nil
}

func (d *Dick) ToString() string {
	str := d.leftb.ToString()
	str += d.shaft.ToString()
	str += d.rightb.ToString()
	return str
}

func NewDick(shaftwidth int,shaftheight int,ballssizes int) *Dick {
	return &Dick{shaft: NewRectangle(shaftwidth,shaftheight),leftb: NewCircle(ballssizes),rightb: NewCircle(ballssizes)}
}
