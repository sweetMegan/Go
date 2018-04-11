package model

type Car struct {
	IEngine
}

func (c Car)TestIEngine(i IEngine)  {
	i.Start()
	i.Speedup()
	i.Stop()
}
