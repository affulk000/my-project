package exercise

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employee []Employee
}

func (m *Manager) AddEmployee(e Employee) {
	m.Employee = append(m.Employee, e)
}
func (m *Manager) RemoveEmployee(id int) {
	for i, e := range m.Employee {
		if e.ID == id {
			m.Employee = append(m.Employee[:i], m.Employee[i+1:]...)
			return
		}
	}
}

func (m Manager) GetAverageSalary() float64 {
	if len(m.Employee) == 0 {
		return 0.0
	}
	var totalSalary float64
	for _, e := range m.Employee {
		totalSalary += e.Salary
	}
	return totalSalary / float64(len(m.Employee))
}

func (m *Manager) GetEmployeeByID(id int) *Employee {
	for _, e := range m.Employee {
		if e.ID == id {
			return &e
		}
	}
	return nil
}
