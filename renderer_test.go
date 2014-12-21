package renderer

import (
	"fmt"
	"strconv"
	"testing"
)

func ExampleDemo1() {
	// Output: <div class="commentBox" data-reactid=".2beudqbmdae" data-react-checksum="-1813045546">Hello, world! I am a CommentBox.</div>
	v := NewRenderer([]string{"assets/demo1.js"}).
		RunCmd("React.renderComponentToString(CommentBox({}))")
	fmt.Println(v)
}

func ExampleDemo2() {
	// Output: <p data-reactid=".12tizqr23xb" data-react-checksum="1539392754"><span data-reactid=".12tizqr23xb.0">Hello, </span><input type="text" placeholder="Your name here" data-reactid=".12tizqr23xb.1"><span data-reactid=".12tizqr23xb.2">!</span></p>
	v := NewRenderer([]string{"assets/demo2.js"}).
		RunCmd("React.renderComponentToString(HelloWorld({}))")
	fmt.Println(v)
}

func ExampleDemo3() {
	v := NewRenderer([]string{"assets/demo3.js"}).
		RunCmd(`
			var data = [
				{"id": 0, "author": "Anonymous", "text": "This is a comment"},
				{"id": 1, "author": "Anonymous", "text": "This is another comment"},
			]
			React.renderComponentToString(CommentBox({data : data}));
		`)
	fmt.Println(v)
}

func benchmarkRender(i int, b *testing.B) {
	r := NewRenderer([]string{"assets/demo3.js"})
	for n := 0; n < b.N; n++ {
		r.RunCmd(`
			var data = [];
			for (i = 0; i < ` + strconv.Itoa(i) + `; i++) {
				data.push({"id": i, "author": "Anonymous", "text": "This is comment #" + i});
			}
			React.renderComponentToString(CommentBox({data : data}));
		`)
	}
}

func BenchmarkRender1(b *testing.B)   { benchmarkRender(1, b) }
func BenchmarkRender5(b *testing.B)   { benchmarkRender(5, b) }
func BenchmarkRender10(b *testing.B)  { benchmarkRender(10, b) }
func BenchmarkRender20(b *testing.B)  { benchmarkRender(20, b) }
func BenchmarkRender50(b *testing.B)  { benchmarkRender(50, b) }
func BenchmarkRender100(b *testing.B) { benchmarkRender(100, b) }
func BenchmarkRender200(b *testing.B) { benchmarkRender(200, b) }
func BenchmarkRender500(b *testing.B) { benchmarkRender(500, b) }
