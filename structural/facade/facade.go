package facade

type SysFacade struct {
	compA *ComponentA
	compB *ComponentB
}

func NewSysFacade(depA, depB string) *SysFacade {
	return &SysFacade{
		compA: NewComponentA(depA),
		compB: NewComponentB(depB),
	}
}

func (f *SysFacade) Init() {
	f.compA.Init()
	f.compB.Init()
}

func (f *SysFacade) Operate() {

	// if sys require this operation
	// first cB operate -> cA operate -> cB operate again
	// and Facade expose New Interface to hide this details
	f.compB.Operate()
	f.compA.Operate()
	f.compB.Operate()
}
