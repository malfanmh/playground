
 CLI application : using [`flags`](https://godoc.org/flag) package in GO to  *Sum , Multiply , Generate Fibonaci and Prime series* , including Unit Testing as well
  
  *`Run`*  this application :
  
```sh
> go build && ./playgroud --sum=19,2

(or)

> go run main.go --sum=19,2
```
`test`  this application :
```
go test -v
```
available flag :
  ```sh
  > go run main.go --help 
 
 -fibo int
        generate fibonnaci series by N , for example : '--fibo=4'
  -multiply value
        multiply of your integer list, for example : '--multiply=3,2...' 
  -prime int
        generate prime series by N , for example '--prime=4' 
  -sum value
        sumary of your integer list , for example : '--sum=3,2...'
    
  ```

