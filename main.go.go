package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	Grades []int
}

func (s *Student) Avg() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0
	for _, v := range s.Grades {
		sum += v
	}
	return float64(sum) / float64(len(s.Grades))
}

func (s *Student) AddGrades(g []int) {
	s.Grades = append(s.Grades, g...)
}

func (s *Student) String() string {
	return fmt.Sprintf("%s -> %v  средний=%.2f", s.Name, s.Grades, s.Avg())
}

func parseGrades(s string) []int {
	out := []int{}
	for _, f := range strings.Fields(s) {
		if v, err := strconv.Atoi(f); err == nil {
			out = append(out, v)
		}
	}
	return out
}

func main() {
	in := bufio.NewReader(os.Stdin)
	students := map[string]*Student{}

	for {
		fmt.Println("\n1=Добавить  2=Список  3=Фильтр по среднему  0=Выход")
		fmt.Print("> ")
		cmd, _ := in.ReadString('\n')
		switch strings.TrimSpace(cmd) {
		case "1":
			fmt.Print("ФИО: ")
			name, _ := in.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("Оценки (через пробел): ")
			line, _ := in.ReadString('\n')
			grades := parseGrades(line)

			if st, ok := students[name]; ok {
				st.AddGrades(grades) // дописываем к существующему
			} else {
				students[name] = &Student{Name: name, Grades: grades}
			}
		case "2":
			for _, st := range students {
				fmt.Println(st)
			}
		case "3":
			fmt.Print("Порог (например 4): ")
			line, _ := in.ReadString('\n')
			t, _ := strconv.ParseFloat(strings.TrimSpace(line), 64)
			for _, st := range students {
				if st.Avg() < t {
					fmt.Printf("%s -> %.2f\n", st.Name, st.Avg())
				}
			}
		case "0":
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
