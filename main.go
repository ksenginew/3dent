package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
)

var sw = `
importScripts("https://unpkg.com/typescript@4.6.2/lib/typescript.js")
self.addEventListener("fetch", (event) => {
  event.respondWith(
    fetch(event.request).then(async (response) => {
      if (
        [
          "application/typescript",
          "text/typescript",
          "video/vnd.dlna.mpeg-tts",
          "video/mp2t",
          "application/x-typescript",
          "application/javascript",
          "text/javascript",
          "application/ecmascript",
          "text/ecmascript",
          "application/x-javascript",
          "application/node",
          "text/jsx",
          "text/tsx",
          "text/plain",
          "application/octet-stream"
        ].includes(response.headers.get("content-type").split(";")[0])
      ) {
        let headers = new Headers(response.headers)
        headers.set("content-type", "application/javascript")
        return new Response(
          ts.transpileModule(await response.text(), {
            compilerOptions: { module: ts.ModuleKind.ESNext }
          }).outputText,
          {
            status: response.status,
            statusText: response.statusText,
            headers: headers
          }
        )
      }
      return response
    })
  )
})
`

// getFreeAddr gets a port and returns the closest free address
func getFreeAddr(port int) string {
	for i := 0; i < 10; i++ {
		addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("localhost:%d", port+i))
		if err != nil {
			continue
		}
		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			continue
		}
		defer l.Close()
		return l.Addr().String()
	}

	return "localhost:0"
}

func main() {
	flag.Parse()
	addr := getFreeAddr(8080)
	cwd, err := os.Getwd()
	if err != nil {
		return
	}
	fs := http.FileServer(http.Dir(cwd))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			fmt.Fprintf(w, `
<html>
<head>
	<title>Hello 3dent</title>
</head>

<body>
    <script>
	if ('serviceWorker' in navigator)
	  navigator.serviceWorker.register('./sw.js')
	</script>
    <script type="module" src="%s"></script>
</body>

</html>`,
				flag.Arg(0))
		} else {
			fs.ServeHTTP(w, r)
		}
	})
	http.HandleFunc("/sw.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		fmt.Fprintf(w, sw)
	})
	fmt.Printf("Starting app on http://%s\n", addr)
	http.ListenAndServe(addr, nil)
}
