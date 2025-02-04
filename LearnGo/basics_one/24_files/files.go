package main

func main() {

	/* File Info Methods
		f, err := os.Open("example.txt")
		if err != nil {
			panic(err)
		}

		finfo, err := f.Stat()
		if err != nil {
			panic(err)
		}

		fmt.Println("File Name:", finfo.Name())
	    fmt.Println("File Size:", finfo.Size())
	    fmt.Println("Last Modified:", finfo.ModTime())
	    fmt.Println("Is Directory:", finfo.IsDir())
	    fmt.Println("File Mode:", finfo.Mode())
	*/

	/*
		f, err := os.Open("example.txt")
		if err != nil {
			panic(err)
		}

		defer f.Close()

		buf := make([]byte, 1024)
		n, err := f.Read(buf)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(buf[:n]))
	*/

	/*..read files..
	f, err := os.ReadFile("example.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(f))
	*/

	/*..read folders..
	dir, err := os.Open("../")
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	files, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			println("Directory:", file.Name())
		} else {
			println("File:", file.Name())
		}
	}
	*/

	/*..create files..
	f, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	f.WriteString("Hello, World!")

	bytes := []byte(" Hello, World!")
	f.Write(bytes)
	*/

	/*..copy files..

	src, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer src.Close()

	dst, err := os.Create("example_copy.txt")
	if err != nil {
		panic(err)
	}

	defer dst.Close()

	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dst)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		writer.WriteByte(b)
		if err != nil {
			panic(err)
		}

	}

	writer.Flush()
	fmt.Println("File Copied!")
	*/

	/*..delete files..
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File Deleted!")
	*/

	/*..rename files..
	err := os.Rename("example_copy.txt", "example_renamed.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File Renamed!")
	*/

}