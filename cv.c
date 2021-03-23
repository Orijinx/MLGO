package main

import (
//  "github.com/anthonynsimon/bild/transform"
//  "github.com/anthonynsimon/bild/effect"
//     "github.com/anthonynsimon/bild/imgio"
"github.com/corona10/goimagehash"
	"fmt"
	"image"
	"os"
	// "crypto/md5"
	// "unsafe"
    
)




func HashImg() {
	file1, _ := os.Open("output.jpg")
	file2, _ := os.Open("output2.jpg")
	defer file1.Close()
	defer file2.Close()

	img1, _ := jpeg.Decode(file1)
	img2, _ := jpeg.Decode(file2)
	hash1, _ := goimagehash.AverageHash(img1)
	hash2, _ := goimagehash.AverageHash(img2)
	distance, _ := hash1.Distance(hash2)
	fmt.Printf("Distance between images: %v\n", distance)

	// hash1, _ = goimagehash.DifferenceHash(img1)
	// hash2, _ = goimagehash.DifferenceHash(img2)
	// distance, _ = hash1.Distance(hash2)
	// fmt.Printf("Distance between images: %v\n", distance)
	// width, height := 8, 8
	// hash3, _ = goimagehash.ExtAverageHash(img1, width, height)
	// hash4, _ = goimagehash.ExtAverageHash(img2, width, height)
	// distance, _ = hash3.Distance(hash4)
	// fmt.Printf("Distance between images: %v\n", distance)
	// fmt.Printf("hash3 bit size: %v\n", hash3.Bits())
	// fmt.Printf("hash4 bit size: %v\n", hash4.Bits())

	var b bytes.Buffer
	foo := bufio.NewWriter(&b)
	_ = hash4.Dump(foo)
	foo.Flush()
	bar := bufio.NewReader(&b)
	hash5, _ := goimagehash.LoadExtImageHash(bar)
}

func main() {


//  img, err := imgio.Open("test.jpg")
//  if err != nil {
//  fmt.Println(err)
//  return
//  }

//  resized := transform.Resize(img, 8, 8, transform.Linear)
//  resized = effect.Grayscale(resized)
// //  rotated := transform.Rotate(resized, 45, nil)
//  if err := imgio.Save("output.jpg", resized, imgio.PNGEncoder()); err != nil {
//  fmt.Println(err)
//  return
//  }

//  img2, err := imgio.Open("test2.jpg")
//  if err != nil {
//  fmt.Println(err)
//  return
//  }

//  resized2 := transform.Resize(img2, 8, 8, transform.Linear)
//  resized2 = effect.Grayscale(resized2)
// //  rotated := transform.Rotate(resized, 45, nil)
//  if err := imgio.Save("output2.jpg", resized2, imgio.PNGEncoder()); err != nil {
//  fmt.Println(err)
//  return
//  }

HashImg()
//  fmt.Print(GetByte("output.png"))
//  fmt.Print(GetByte("output2.png"))
//  fmt.Println(GetByte("output.jpg"))
//  fmt.Println(GetByte("output2.jpg"))

}

func GetByte(name string) []uint32{

	file, err := os.Open(name) // For read access.
	if err != nil {
		fmt.Println(err)
	}
	
	 im, _, err := image.Decode(file)
	 if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(im.Bounds().Max.X)
	var arr =make([]uint32,im.Bounds().Max.X*im.Bounds().Max.X)
	// var matrix = make([][]image.Color,im.Bounds().Max.X,im.Bounds().Max.Y)
	var k =0
	for i := 0; i < im.Bounds().Max.X; i++ {
		for j := 0; j < im.Bounds().Max.Y-1; j++ {
			r,g,b,_  := im.At(i,j).RGBA()
			// fmt.Println(r,g,b)
			arr[k]=(r+g+b)/3
			k++
		  }
	  }
	
	//   var max uint32 =0
	  var sum uint32 =0
	
	for i :=range arr{
		sum=sum+arr[i]
	}
	sum=sum/uint32(len(arr))
	var x =make([]byte,len(arr))
	for i :=range arr{
		if (arr[i] >=sum ){
			x[i]=1
		}else{
			x[i]=0
		}
	}
	
	// fmt.Println("[MAX]")
	// fmt.Println(max)
	// fmt.Println("[SUM]")
	// fmt.Println(sum/uint32(len(arr)))
	// fmt.Println("[BYTE]")
	// fmt.Println(x)

	// var bt = make([][4]byte,len(arr))
	


return arr	
}