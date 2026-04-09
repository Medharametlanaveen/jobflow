package service

type JobService struct{}

func NewJobService()*JobService{
	return &JobService{}
}

func (s *JobService) Health() string {
	return "Ok"
}