package modules

type Arjuna1Controller struct {
	uc Arjuna1Usecase
}

func CreateArjuna1Controller(uc Arjuna1Usecase) Arjuna1Controller {
	return Arjuna1Controller{
		uc: uc,
	}
}
