package utils

import (
	`bufio`
	`fmt`
	`os`
	`strings`
)

/**
	Read .env file and store in os.environment
 	this data can then be accessed through os.Getenv()
 */
func ReadEnv(filename string) error{
	var line string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	bf := bufio.NewReader(file)
	line, err = bf.ReadString('\n')
	if err != nil {
		return err
	}

	for line != "" {
		if 1 == strings.Count(line, "=") {
			env := strings.SplitAfter(line, "=")
			err = os.Setenv(strings.Trim(env[0], "="),
				strings.Trim(env[1], "\n"))}
			if nil != err{
				fmt.Println(err)
			}
		line, err = bf.ReadString('\n')
	}
	return nil
}
