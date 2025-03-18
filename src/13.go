// Create a new HTTP server
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  // Write some HTML to the response
  w.Write([]byte("<html><body>Hello World!</body></html>"))
})

// Start the server on port 8080
log.Fatal(http.ListenAndServe(":8080", nil))
