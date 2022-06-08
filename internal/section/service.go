package section

import ( 
	"fmt"
)
 
type Service interface {
	ListarSectionAll() ([]Section, error)
	ListarSectionOne(id int) (Section, error)
	CreateSection(newSection SectionRequest) (Section, error)
	UpdateSection(id int, sectionUp Section) (Section, error)
	DeleteSection(id int) error
}


type service struct {
	repository Repository
}


func (s service) ListarSectionAll() ([]Section, error){ 
	return s.repository.ListarSectionAll()
}


func (s service) ListarSectionOne(id int) (Section, error) {
	return s.repository.ListarSectionOne(id)	 
}


func (s service) CreateSection(newSection SectionRequest) (Section, error) {
	
	var sectionList []Section 
	
	sectionList, err := s.repository.ListarSectionAll()
	if err != nil {
		return Section{}, err
	}
	for index := range sectionList {
		if sectionList[index].SectionNumber == newSection.SectionNumber {
			return Section{}, fmt.Errorf("section invalid, section_number field must be unique")
		}
	}
	var sec Section =Section{ SectionNumber: newSection.SectionNumber, CurrentTemperature: newSection.CurrentTemperature,
		MinimumTemperature: newSection.MinimumTemperature, CurrentCapacity: newSection.CurrentCapacity, 
		MinimumCapacity: newSection.MinimumCapacity, MaximumCapacity: newSection.MaximumCapacity, 
		WareHouseId: newSection.WareHouseId,ProductTypeId: newSection.ProductTypeId}

	return s.repository.CreateSection(sec)
}

func (s service) UpdateSection(id int, sectionUp Section) (Section, error) { 
	var sectionList []Section 
	sectionList, err := s.repository.ListarSectionAll()
	if err != nil {
		return sectionUp, err
	}
	
	for index := range sectionList {
		if sectionList[index].Id != id && sectionList[index].SectionNumber  == sectionUp.SectionNumber {
			return Section{}, fmt.Errorf("this section_number %d is already registered", sectionUp.SectionNumber)
		}
	}
	return s.repository.UpdateSection(id, sectionUp)	
}

func (s service) DeleteSection(id int) error {
	return s.repository.DeleteSection(id)
}

func NewService(repository Repository) Service {
	return &service{ repository: repository }
}
