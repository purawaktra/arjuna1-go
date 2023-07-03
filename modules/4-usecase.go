package modules

type Arjuna1Usecase struct {
	repo Arjuna1Repo
}

func CreateArjuna1Usecase(repo Arjuna1Repo) Arjuna1Usecase {
	return Arjuna1Usecase{
		repo: repo,
	}
}
