# flexy

### Installation

```bash
go get github.com/morteza-r/flexy
```


### Example doc object
```go
type User struct {
    Id      float64 `json:"id"`
    Name    string  `json:"name"`
    Family  string  `json:"family"`
    Active  bool    `json:"active"`
    Account Account `json:"account"`
}

type Account struct {
    Name   string   `json:"name"`
    Credit int64    `json:"credit"`
    Tags   []string `json:"tags"`
}
```

### Add new doc
```go
user := User{
    Id:       1,
    Name:     "morteza",
    Family:   "rostami",
    Active:   false,
    Account:  Account{
      Credit: 1200,
    },
}

err := client.Query().
    Table("users").
    Model(&user).
    Add()

if err != nil {
    fmt.Println(err)
}
```



### Check doc exist
```go
user := User{
    Id:       1,
}

_, ok := client.Query().
    Table("users").
    Model(&user).
    Exist()

fmt.Println(ok)
```

### Update a doc 
Only provided fields will update. In this example only active field will update to true and other fields will remain same.

```go
user := User{
    Id:       1,
    Active:   true,
}

err := client.Query().
    Table("users").
    Model(&user).
    Update()

if err != nil {
    fmt.Println(err)
}

fmt.Println(user)
```

### Replace a doc
The new doc will replace old one.

```go
user := User{
    Id:       1,
    Name:     "morteza",
    Family:   "rostami",
    Active:   false,
    Account:  Account{
      Credit: 1200,
    },
}

err := client.Query().
    Table("users").
    Model(&user).
    Replace()

if err != nil {
    fmt.Println(err)
}
```

### Get a doc
```go
user := User{
    Id: 1,
}

err := client.Query().
    Table("users").
    Model(&user).
    Get()

if err != nil {
    fmt.Println(err)
}

fmt.Println(user)
```

### Get multiple docs
```go
var users []User

err := client.Query().
    Table("users").
    Model(&users).
    Where("account.credit", "==", 1200).
    OrderString("family").
    Get()
    
if err != nil {
    fmt.Println(err)
}

fmt.Println(users)
```

### Delete a doc
```go
user := User{
    Id: 1,
}

err := client.Query().
    Table("users").
    Model(&user).
    Get()

if err != nil {
    fmt.Println(err)
}
```
