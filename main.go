package main

import "fmt"

func main() {

	length := 10.0 //Length of the rectangle
	width := 5.0   //Width of the rectangle

	area := length * width

	fmt.Printf("The area of the rectangle is: %.2f square units.\n", area)

}

/*
1) mkdir test01
2) cd test01
// create files and run check it works
3) git init
4) git add . ( .is evrything)
5) git commit -m "First comment"
6)git branch -M main
7) git status
8) Create a github.com repository- same name as folder
9) git remote add origin https://github.com/dayanajoseph/test01.git
10) git remote set-url origin https://dayanajoseph:$GITHUB_TOKEN@github.com/dayanajoseph/test01.git
--
11) git push -u origin main
export GITHUB_TOKEN="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"
we changing test 02
----

:q!- quiting - if we get to vi
wq
*/
