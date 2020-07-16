package params

type Init struct {
	Port string
}

func InitParams() Init {
	return InitManual()
}

func InitManual() Init {
	return Init{Port: ":9000"} 
}