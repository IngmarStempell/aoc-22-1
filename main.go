package main

import ("fmt" ; "os"; "bufio"; "strings"; "strconv"; "sort")

func main() {
	fmt.Println("Reading calories")
	inputHandle, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File not found")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(inputHandle)
	line, err := reader.ReadString('\n')
	i := 0
	max := 0
	var sums []int
	for  line != "" {
		if err != nil {
			fmt.Println("Input not found")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		
		if line == "\n" {
			fmt.Println("---")
			if i >= max {
				max = i
			}
			sums = append(sums, i)
			i = 0
			
		} else {
			//fmt.Println(line)
			value,err :=strconv.Atoi(strings.Trim(line,"\n"))
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(value)
			i += value
		}
		
		
		line,err =reader.ReadString('\n')
		
	}

 
	
	fmt.Println("done Reading calories")
	fmt.Println(max)

	    sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	fmt.Println(sums)
	fmt.Println(sums[0])
	fmt.Println(sums[1])
	fmt.Println(sums[2])

	fmt.Println(sums[0] + sums[1]+ sums[2])


}

