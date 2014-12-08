// See LICENSE.txt for licensing information.

package main

import (
	"encoding/json"
	"expvar"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/drbig/ricons"
)

const (
	VERSION = `0.0.1`
)

type format struct {
	mime string
	fmt  ricons.Format
}

type Info struct {
	Generators map[string]string `json:"generators"`
	Formats    []string          `json:"formats"`
	Versions   map[string]string `json:"versions"`
}

var (
	fAddr  string
	fGens  bool
	fBound int
	fQuiet bool
)

var (
	cMapGens = expvar.NewMap("generators")
	cMapFmts = expvar.NewMap("formats")
	cIntHits = expvar.NewInt("hits")
)

var (
	gens map[string]string
	fmts map[string]format
	info []byte
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&fAddr, "a", ":3232", "server bind address")
	flag.BoolVar(&fGens, "l", false, "show generators and exit")
	flag.IntVar(&fBound, "b", 256, "image bound (e.g. max 256x256)")
	flag.BoolVar(&fQuiet, "q", false, "disable logging")

	gens = make(map[string]string, len(ricons.Registry))
	for k, v := range ricons.Registry {
		gens[k] = v.String()
	}

	fmts = map[string]format{
		"png":  format{"image/png", ricons.PNG},
		"gif":  format{"image/gif", ricons.GIF},
		"jpeg": format{"image/jpeg", ricons.JPEG},
	}

	is := &Info{
		gens,
		make([]string, len(fmts)),
		map[string]string{
			"ricons":  ricons.VERSION,
			"riconsd": VERSION,
		},
	}
	i := 0
	for k, _ := range fmts {
		is.Formats[i] = k
		i++
	}
	in, err := json.Marshal(is)
	if err != nil {
		panic(fmt.Sprintf("can't marshal info: %s", err))
	}
	info = in
}

func main() {
	flag.Parse()
	if fGens {
		for _, v := range gens {
			fmt.Println(v)
		}
		os.Exit(0)
	}
	if fQuiet {
		log.SetOutput(ioutil.Discard)
	}

	http.HandleFunc("/info.json", handleInfo)
	http.HandleFunc("/", handleIcon)
	log.Println("Starting HTTP server at", fAddr)
	go func() {
		log.Fatalln(http.ListenAndServe(fAddr, nil))
	}()
	sigwait()
}

func handleInfo(w http.ResponseWriter, req *http.Request) {
	cIntHits.Add(1)
	log.Println(req.RemoteAddr, req.Method, req.RequestURI)

	w.Header().Set("Content-Type", "application/json")
	w.Write(info)
	return
}

func handleIcon(w http.ResponseWriter, req *http.Request) {
	cIntHits.Add(1)
	log.Println(req.RemoteAddr, req.Method, req.RequestURI)

	w.Header().Set("Content-Type", "application/json")
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) != 5 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false, "msg": "wrong request URI"}`))
		return
	}
	g, ok := ricons.Registry[parts[1]]
	if !ok {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"success": false, "msg": "generator not found"}`))
		return
	}
	f, ok := fmts[parts[2]]
	if !ok {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"success": false, "msg": "unknown image format"}`))
		return
	}
	wi, err := strconv.Atoi(parts[3])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false, "msg": "error parsing image width"}`))
		return
	}
	if wi < 1 || wi > fBound {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false, "msg": "image width out of range"}`))
		return
	}
	h, err := strconv.Atoi(parts[4])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false, "msg": "error parsing image height"}`))
		return
	}
	if h < 1 || h > fBound {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false, "msg": "image height out of range"}`))
		return
	}
	i, err := g.NewIcon(wi, h)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success": false, "msg": "error generating icon, sorry"}`))
		return
	}
	cMapGens.Add(parts[1], 1)
	cMapFmts.Add(parts[2], 1)
	w.Header().Set("Content-Type", f.mime)
	w.WriteHeader(http.StatusOK)
	i.Encode(f.fmt, w)
	return
}
