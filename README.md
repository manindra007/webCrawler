#webCrawler
steps to run
  1. clone the repository in $go Path/src
  2. cd ../path/src
  3. go run main.go
 
Run postman.
  1. select POST Api option
  2. select Body
  3. select Raw -> Json.
  4.{
        "urls": ["https://google.com", "https://github.com"]
    }
  5. http://localhost:8081/api/crawl
  6. send.
