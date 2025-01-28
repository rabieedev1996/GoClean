The Authorize method populates the TokenUser object using data extracted from the token.<br>
TokenUser is an object used for injection into other structs, providing access to user data extracted from the token.<br>
For example, the code below is one of the Gin handlers that uses the Authorize method:

```go
router.GET("/api/sample", func(c *gin.Context) {
‍  container := GetDIContainer()
  container.Invoke(func( tokenUser *Models.TokenUser) {
  //filling the TokenUser object
  jwtHelper.GINAuthorize(tokenUser, ...)
  ...
})
