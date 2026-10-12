package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

//Added the missing Grade letters 
func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}


func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 55, Assignment)
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

//Testing when no grades a inputted it returns F (assignment missing)
func TestNonInputtedGrade(t *testing.T){
	gradeCalculator := NewGradeCalculator()

	actual_val := gradeCalculator.GetFinalGrade()

	if actual_val != "F" {
		t.Errorf("All grades are not inputted")
	}
}

//Testing if the category is correctly assigned
func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Errorf("Expected 'assignment'")
	}
	if Exam.String() != "exam" {
		t.Errorf("Expected 'exam'")
	}
	if Essay.String() != "essay" {
		t.Errorf("Expected 'essay'")
	}
}

//Testing multiple grades under same category
func TestMultipleGrades (t *testing.T){
	expected_value := "C"
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("HW1", 100, Assignment)
	gradeCalculator.AddGrade("HW2", 0, Assignment)
	gradeCalculator.AddGrade("EXAM", 100, Exam)
	gradeCalculator.AddGrade("Essay", 100, Essay)

	actual_val := gradeCalculator.GetFinalGrade()

	if expected_value != actual_val {
		t.Errorf("Expected value and calculated value is incorrect")
	}
}

//Testing Borderline grades
func TestBorderlineAtoB(t *testing.T){
	gc := NewGradeCalculator()
	gc.AddGrade("hw", 90, Assignment)
	gc.AddGrade("exam", 90, Exam)
	gc.AddGrade("essay", 90, Essay)

	if got := gc.GetFinalGrade(); got != "A"{
		t.Errorf("Not A, Expected A")
	}

	gc2 := NewGradeCalculator()
	gc2.AddGrade("hw", 89, Assignment)
	gc2.AddGrade("exam", 89, Exam)
	gc2.AddGrade("essay", 89, Essay)

	if got := gc2.GetFinalGrade(); got != "B"{
		t.Errorf("Not B, Expected B")
	}
}

//Testing just one category input
func TestOnlyInputted(t *testing.T){
	expected_val := "F"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("exam 1", 100, Exam)

	actual_val := gradeCalculator.GetFinalGrade()

	if expected_val != actual_val{
		t.Errorf("Expected F, it is incorrect.")
	}
}