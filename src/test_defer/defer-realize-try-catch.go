package main

func main() {
    defer func(){
        if err := recover(); err != nil{
            println("catch exception 1")
        }else{
            println("run normally 1")
        }
    }()

    println("run 1")

    defer func(){
        if err := recover(); err != nil{
            println("catch exception 2")
            goto X
        }else{
            println("run normally 2")
        }
    }()

    println("run 2")
    panic("exception happen !")
    println("run 3")
X:
    println("run 4")

}



