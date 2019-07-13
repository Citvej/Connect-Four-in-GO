// https://golang.org/pkg/syscall/js/
// https://www.aaron-powell.com/posts/2019-02-05-golang-wasm-2-writing-go/
// https://tutorialedge.net/golang/go-webassembly-tutorial/
// https://jaxenter.com/using-go-webassembly-applications-147272.html

/*
 set GOOS=js
 set GOARCH=wasm
 go build -o main.wasm main.go
*/

package main

import (
	"fmt"
	"syscall/js"
)

var c chan bool
var celaMatrika[6][7] string;
// init is called before main
func init() {
	c = make(chan bool)
}

func doubleAge(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0]

	message.Set("age", message.Get("age").Int()*2)

	return nil
}

func printNameAndAge(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0]

	return message.Get("name").String() + message.Get("age").String()
}

func addOne(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0]

	// newArray := make([]interface{}, message.Length())

	// for i := 0; i < message.Length(); i++ {
	// 	newArray[i] = message.Index(i).Int() + 1
	// }

	// return js.ValueOf(newArray)

	for i := 0; i < message.Length(); i++ {
		message.SetIndex(i, js.ValueOf(message.Index(i).Int()+1))
	}

	return js.ValueOf(message)
}

func printMessage(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0].String()

	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", message)
	document.Get("body").Call("appendChild", p)

	//c <- true // send signal to exit go app, printMessage won't be available in JS
	return nil
}

// func typedArrayToByteSlice(arg js.Value) []byte {
// 	length := arg.Length()
// 	bytes := make([]byte, length)
// 	for i := 0; i < length; i++ {
// 		bytes[i] = byte(arg.Index(i).Int())
// 	}
// 	return bytes
// }

func pridobiElementEnegaPoEnega(this js.Value, inputs []js.Value) interface{} {
	element := inputs[0].String();
	vrstica := inputs[1].Int();
	stolpec := inputs[2].Int();
	celaMatrika[vrstica][stolpec] = element;

	return true;
}

func preveriZmagovalca(this js.Value, inputs []js.Value) interface{} {
	var matrika[6][7] string;
	matrika = celaMatrika;
	// matrika = (string) bytes.NewReader(typedArrayToByteSlice(inputs[0]));
	stolpec := inputs[0].Int();
	vrstica := inputs[1].Int();
	barva := inputs[2].String();

	alert := js.Global().Get("alert")
	js.Global().Get("console").Call("log", matrika[5][6])
	js.Global().Get("console").Call("log", stolpec)
	js.Global().Get("console").Call("log", vrstica)
	js.Global().Get("console").Call("log", barva)


	if preveriHorizontalno(matrika, stolpec, vrstica, barva) {
		return alert.Invoke("zmagal je: " + barva);
	}else if preveriVertikalno(matrika, stolpec, vrstica, barva) {
		return alert.Invoke("zmagal je: " + barva);
	// }else if preveriDiagonalno1(matrika, stolpec, vrstica, barva) {
	// 	return alert.Invoke("zmagal je: " + barva);
	// }else if preveriDiagonalno2(matrika, stolpec, vrstica, barva) {
	// 	return alert.Invoke("zmagal je: " + barva);
	}

	return false;

}

func preveriDiagonalno1(matrika[6][7] string, stolpec int, vrstica int, barva string) bool{
	g := true;
	d := true;
	stevec := 0;
	for i:=1; i<4; i++ {
		stolpecGor := stolpec - i;
		stolpecDol := stolpec + i;
		vrsticaLevo := vrstica - i;
		vrsticaDesno := vrstica + i;
		
		/* preverimo če pri premikih po diagonali nismo izven rangea */
		if(vrsticaLevo < 0 || stolpecDol > 6){
			g = false;
		}
		if(vrsticaDesno < 0 || stolpecGor > 6){
			d = false;
		}

		/* preverimo ujemajoče žetone */
		if( vrsticaLevo >= 0 && stolpecDol <= 6 ){
			if(matrika[vrsticaLevo][stolpecDol] != barva){
				g = false;
			}
		}
		if( stolpecGor >= 0 && vrsticaDesno <= 5 ){
			if(matrika[vrsticaDesno][stolpecGor] != barva){
				d = false;
			}
		}
		
		//če se žetoni ujemajo povečamo števec
		if(g == true){
			stevec++;
		}
		if(d == true){
			stevec++;
		}
	}
	// če je števec >= 3 vrnemo true
	if(stevec >= 3){
		return true;
	}
	return false;
}

func preveriDiagonalno2(matrika[6][7] string, stolpec int, vrstica int, barva string) bool{
	g := true;
	d := true;
	stevec := 0;
	for i:=1; i<4; i++ {
		stolpecGor := stolpec - i;
		stolpecDol := stolpec + i;
		vrsticaLevo := vrstica - i;
		vrsticaDesno := vrstica + i;
		
		/* preverimo če pri premikih po diagonali nismo izven rangea */
		if(vrsticaLevo < 0 || stolpecDol > 5){
			g = false;
		}
		if(vrsticaDesno < 0 || stolpecGor > 5){
			d = false;
		}

		/* preverimo ujemajoče žetone */
		if( g == true ){
			if(matrika[vrsticaLevo][stolpecGor] != barva){
				g = false;
			}
		}
		if( d == true ){
			if(matrika[vrsticaDesno][stolpecDol] != barva){
				d = false;
			}
		}

		//če se žetoni ujemajo povečamo števec
		if(g == true){
			stevec++;
		}
		if(d == true){
			stevec++;
		}
	}
	// če je števec >= 3 vrnemo true
	if(stevec >= 3){
		return true;
	}
	return false;
}

func preveriVertikalno(matrika[6][7] string, stolpec int, vrstica int, barva string) bool{
	g := true;
	d := true;
	gor := 0;
	dol := 0;
	stevec := 0;
	for i:=1; i<4; i++ {
		gor = stolpec - i;
		dol = stolpec + i;
		/* če je out of bounds je false */
		if(gor < 0){
			g = false;
		}
		if(dol > 5){
			d = false
		}
		/* če barva ne ustreza je tudi false */
		if( gor >= 0 ){
			if(matrika[vrstica][gor] != barva){
				g = false;
			}
		}
		if( dol <= 5 ){
			if(matrika[vrstica][dol] != barva){
				d = false;
			}
		}

		// če je true prištejemo score
		if(g == true){
			stevec++;
		}
		if(d == true){
			stevec++;
		}
	}
	// če je score >= 3 ujemajoče žetone vrnemo true
	if(stevec >= 3){
		return true;
	}
	return false;
}

func preveriHorizontalno(matrika[6][7] string, stolpec int, vrstica int, barva string) bool{
	l := true;
	d := true;
	levo := 0;
	desno := 0;
	stevec := 0;
	for i:=1; i<4; i++ {
		levo = vrstica - i;
		desno = vrstica + i;
		/* če je out of bounds je false */
		if(levo < 0){
			l = false;
		}
		if(desno > 5){
			d = false
		}
		/* če barva ne ustreza je tudi false */
		if( levo >= 0 ){
			if(matrika[levo][stolpec] != barva){
				l = false;
			}
		}
		if( desno <= 5 ){
			if(matrika[desno][stolpec] != barva){
				d = false;
			}
		}

		// če je true prištejemo score		
		if(l == true){
			stevec++;
		}
		if(d == true){
			stevec++;
		}
	}
	
	// če je score >= 3 ujemajoče žetone vrnemo true
	if(stevec >= 3){
		return true;
	}
	return false;
}

func main() {
	fmt.Println("Hello WASM from Go to browser console!")


	js.Global().Set("printMessage", js.FuncOf(printMessage))
	js.Global().Set("addOne", js.FuncOf(addOne))
	js.Global().Set("printNameAndAge", js.FuncOf(printNameAndAge))
	js.Global().Set("doubleAge", js.FuncOf(doubleAge))

	js.Global().Set("pridobiElementEnegaPoEnega", js.FuncOf(pridobiElementEnegaPoEnega))
	js.Global().Set("preveriZmagovalca", js.FuncOf(preveriZmagovalca))

	js.Global().Call("printMessage", "Hello from WASP")

	<-c // we need to block, or printMessage won't be available
}
