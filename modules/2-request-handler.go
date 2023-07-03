package modules

type Arjuna1RequestHandler struct {
	ctrl Arjuna1Controller
}

func CreateArjuna1RequestHandler(ctrl Arjuna1Controller) Arjuna1RequestHandler {
	return Arjuna1RequestHandler{
		ctrl: ctrl,
	}
}
