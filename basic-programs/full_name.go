// Learning functions in go

package main

func full_name(firstname string, lastname string) string {
	return firstname + " " + lastname
}

func b_full_name(firstname string, lastname string) (string, error) {
	return firstname + " " + lastname, nil
}

func main() {
	println(full_name("Kalebu", "Gwalugano"))
	fullname, error := b_full_name("Kalebu", "Gwalugano")
	if error != nil {
		println(error)
	}
	print(fullname)
}
