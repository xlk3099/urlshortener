### Some thoughts and issues encountered ruing implementation

##### About the algorithm to shorten a long URL
Did some research what an url shortening service is about and the algorithm to implement. 
* wiki: https://en.wikipedia.org/wiki/URL_shortening
* Quora: https://www.quora.com/What-are-the-http-bit-ly-and-t-co-shortening-algorithms
* stack-overflow:http://stackoverflow.com/questions/742013/how-to-code-a-url-shortener

Most of these articiles suggested base62 encoding. I tried google url shortening service, and found for the same given input long URL, it generates different short URL.
So I came out with two algorithms. 
1. Base62 encoding: generate the short url based on the input request ID.
2. Generate a random 6 chars base62 string, as it is 62*6, the possibility of collision is really low. 

But I chose the first algorithm to implement, as it is more efficient, and can have a decode function to track the original value if necessary. 

##### About Golang
* A quick walkthrough on the "A Golang tour Go"" of course, and some reading on "Effective Go"
* IDE? : Sublime + Gosublime
* What is Golang coding practises? : https://peter.bourgon.org/go-best-practices-2016/
* What should be a Golang web app structure like? : https://larry-price.com/blog/2015/06/25/architecture-for-a-golang-web-app
* How to write a class in Golang? : Golang does not have an explicit class type, using struct + a set of methods on that type
* Any available webframework to help build restful api and parsing json? : I choose **Gin** from {"Gin", "Martini", "Beego"} beacuse it's lightweight and faster. 
* Testing framework: **Goconvey**, A BDD testing framework.
##### About MongoDB
*How to set an auto-increment id in mongodb: https://docs.mongodb.com/v3.0/tutorial/create-an-auto-incrementing-field/

##### About Docker
- How to create image, create a container? docker start tutorial
- How to manage multiple containers? As my app has a dependency on MongoDB, so I chose docker-

##### Others
* Travis CI : A CI for projects on github
* Codecov.io : A Code coverage report tool for projects on github
* Go report : An online tool which integrates go tools like gofmt, go_vet, golint, gocyclo, etc.