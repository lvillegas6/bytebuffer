## Install

With a  [correctly configured](https://golang.org/doc/install#testing)  Go toolchain:

> go get -u github.com/boogst/bytebuffer

## Examples
  
 Let's start creating your byteBuffer.  We have to call the `NewBuffer` function, it receives a string that indicates the type of buffer you want to create: **"big" or "little"**

    buffBig := bytebuffer.NewBuffer("big")
    buffLittle := bytebuffer.NewBuffer("little")
   
     

### To add data we only call the functions:

    buff := bytebuffer.NewBuffer("little")
    buff.PutShort(52)
    buff.PutInt(2)
    buff.PutLong(2)
    buff.PutFloat(2.5)
    buff.PutDouble(4.5)
    buff.Put([]byte("Hello World"))
   
 
   ###   Print buffer:
You have to call the function `.Array()`, it returns a slice of bytes
	
	fmt.Println(buff.Array())
	Output:
	[52 0 2 0 0 0 2 0 0 0 0 0 0 0 0 0 32 64 0 0 0 0 0 0 18 64 72 101 108 108 111 32 87 111 114 108 100]

### Get data:
To retrieve the information added to the buffer you have to call the "get" functions in the order in which the data was added:

    fmt.Println(buff.GetShort()) 	//52
    fmt.Println(buff.GetInt())   	//2
    fmt.Println(buff.GetLong())  	//2
    fmt.Println(buff.GetFloat()) 	//2.5
    fmt.Println(buff.GetDouble())	//4.5
   