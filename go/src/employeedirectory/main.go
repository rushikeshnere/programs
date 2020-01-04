package main

import(
	"fmt"
	"bufio"
	"os"
)

type Employee struct {
        id int
        name string
}

type EmployeeDirectory struct {
     managerEmpMap map[int][]*Employee //Map of manager and array of references of its reporting employees
     empDirectory map[int]*Employee    //Map of emp ID and reference of associated employee object	
}

var empDirImpl EmployeeDirectory

func main(){
	empDirImpl = EmployeeDirectory{}
	
	empDirImpl.managerEmpMap = make(map[int][]*Employee)
	empDirImpl.empDirectory = make(map[int]*Employee)
	var empID, managerID int = 0, 0
	var firstEmpID, secondEmpID = 0, 0

	fmt.Println("Enter employee details. After completion of entering employee details of all employees enter -1 as employee ID.")
	reader := bufio.NewReader(os.Stdin)

	for {	
		fmt.Println("Enter employee ID")
		_, errReadingID := fmt.Scanf("%d", &empID)
		if (errReadingID != nil) {
			continue
		}
		if(empID == -1) {
			break;
		}
		if (doesEmployeeAlreadyExist(empID)) {
			fmt.Println("Employee with this emp ID is already exists, please enter another emp ID")
			continue
		}	
		
		fmt.Println("Enter employee Name")
		empName, _ := reader.ReadString('\n')

		fmt.Println("Enter manager ID")
	        _, errReadingName := fmt.Scanf("%d", &managerID)	
		if (errReadingName != nil) {
                        continue
                }
		addEmployee(empID, empName, managerID)	
	}

	fmt.Println("Enter IDs of two employees whose closest common manager needs to be found. After completion of finding closest common manager enter -1 as employee ID.")

	for {
		fmt.Println("Enter 1st employee ID")
		_, errReadingFirstID := fmt.Scanf("%d", &firstEmpID)
                if (errReadingFirstID != nil) {
                        continue
                }
                if(firstEmpID == -1) {
                        break;
                }

		fmt.Println("Enter 2nd employee ID")
		 _, errReadingSecondID := fmt.Scanf("%d", &secondEmpID)
                if (errReadingSecondID != nil) {
                        continue
                }
                if(secondEmpID == -1) {
                        break;
                }
		if(firstEmpID == 1 || secondEmpID == 1) {
			fmt.Printf("Closest common manager = %+v\n", *(empDirImpl.empDirectory[1]))
			continue
		}
		closestCommonManagerID := empDirImpl.getClosestManager(firstEmpID, secondEmpID)
		fmt.Println("closestCommonManagerID = ", closestCommonManagerID)
		fmt.Printf("Common closest manager = %+v\n", *(empDirImpl.empDirectory[closestCommonManagerID]))
	}
}

//Adds employee to employee directory
func addEmployee(id int, name string, managerID int) {
	emp := Employee{id, name}
	if(managerID == -1) {
		fmt.Println("Is a CEO")
		empDirImpl.managerEmpMap[id] = append(empDirImpl.managerEmpMap[id], &emp)
	} else {
		empDirImpl.managerEmpMap[managerID] = append(empDirImpl.managerEmpMap[managerID], &emp)		
	}
	empDirImpl.empDirectory[id] = &emp
}

//Returns true if employee for the given empID already exist
func doesEmployeeAlreadyExist(id int) (bool) {
	for k := range empDirImpl.empDirectory {
		if (k == id) {
			return true
		}
	}	
	return false
}

//Returns closest common manager
func (empDir EmployeeDirectory) getClosestManager(id1 int, id2 int) (int) {
         emp1Managers := []int{}
         emp2Managers := []int{}
	 emp1MgrCounter := 0
	 emp2MgrCounter := 0
       
                        emp1Managers = append(emp1Managers,getManagers(id1)...)
                        emp2Managers = append(emp2Managers,getManagers(id2)...)
                        for emp1MgrCounter = 0; emp1MgrCounter < len(emp1Managers); emp1MgrCounter ++ {
                                emp1Managers = append(emp1Managers,getManagers(emp1Managers[emp1MgrCounter])...)
                        }
			for emp2MgrCounter = 0; emp2MgrCounter < len(emp2Managers); emp2MgrCounter ++ {
                                emp2Managers = append(emp2Managers,getManagers(emp2Managers[emp2MgrCounter])...)
                        }
                commonManagerList := getCommonManagers(emp1Managers, emp2Managers)
                if(len(commonManagerList) != 0){
			return commonManagerList[0]
                }	
	return -1
}

//Returns managers of given employeeID
func getManagers(id int) ([]int){
	mgr := []int{}
	if id == 1 {
		return mgr
	}
	for key, value := range empDirImpl.managerEmpMap {
                for _, emp := range value {
                     if(emp.id == id) {
				mgr = append(mgr, key)		
			}  
                }
    	}
	return mgr
}

// Returns common managers between managers of two employees
func getCommonManagers(mgr1List []int, mgr2List []int) ([]int) {
	commonManagerList := []int{}
	for _, v := range(mgr1List) {
		for _, k := range(mgr2List) {
			if(v == k) {
				commonManagerList = append(commonManagerList, v)
			}
		}
	}
	return commonManagerList
}
