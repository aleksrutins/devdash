package private

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
)

func ForwardContainer(internalPort uint16, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	container := GetContainer(vars["container"])
	pubPort := container.Ports[sort.Search(len(container.Ports), func(i int) bool {
		return container.Ports[i].PrivatePort == internalPort
	})]
	fmt.Println("Serving " + fmt.Sprint(pubPort.PublicPort))
	r.URL.Path = r.URL.Path[len("/container/"+vars["container"]):]
	ServeReverseProxy("http://localhost:"+fmt.Sprint(pubPort.PublicPort), w, r)
}
