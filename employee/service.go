package employee

import "github.com/GolangPruebaRestApi/helper"

type Service interface {
	GetEmployees(params *getEmployeesResquest) (*EmployeeList, error)
	//GetTotalEmployees() (int64, error)
}

type service struct {
	repo Repository
}

func NerService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetEmployees(params *getEmployeesResquest) (*EmployeeList, error) {
	employees, err := s.repo.GetEmployees(params)
	helper.Catch(err)
	totalEmployees, err := s.repo.GetTotalEmployees()
	helper.Catch(err)

	return &EmployeeList{
		Data:         employees,
		TotalRecords: totalEmployees,
	}, nil
}
