### xEnv: 🛠️🚀 Simplify Your Environment Management in Go!

Picture this: you're launching your Go application, and every environment variable is set up just right—no surprises, no mishaps! xEnv not only ensures that your environment variables are in check but also provides an easy way to load variables from a `.env` file, a feature not supported by default in Go. Perfect setup leads to perfect launches! 🚀

#### Get Started with xEnv

1. **Installation**
   Kick off by adding xEnv to your project:
   ```go
   go get github.com/dev3mike/go-xenv
   ```

2. **Setup Your Environment Configuration**
   Best practices suggest keeping your environment configurations in a dedicated file. Let's set that up under `environments/environments.go`:
   ```go
   package environments

   import (
       "github.com/dev3mike/go-xenv"
       "log"
   )

   var Env struct {
        Host       string `json:"HOST" validators:"required,minLength:3,maxLength:50"`
        AdminEmail string `json:"ADMIN_EMAIL" validators:"email"`
   }

   func init() {
        // Load environment variables from a .env file
        if err := xenv.LoadEnvFile(".env"); err != nil {
            log.Panic("Error loading .env file: ", err)
        }

        // Validate environment variables
        if err := xenv.ValidateEnv(&Env); err != nil {
            log.Panic("Failed to validate environment: ", err)
        }

        log.Println("Environment validated successfully! 🎉")
    }
   ```

   3. **Integrate with Your Main Application**
   Incorporate the `environments` package in your main application file to ensure your environment variables are validated at startup:
   ```go
   package main

   import (
       "path/to/your/project/environments"
       "log"
   )

   func main() {
       // Application logic goes here
       log.Println("Application is running with validated environment settings!")

       // This is how you can get your validated and transformed env variables
       host := environments.Env.Host
       adminEmail := environments.Env.AdminEmail
   }
   ```
   
#### Use Transformers
You can also use transformers to cleanup the environment variables when needed

```go
   type Environment struct {
        Host       string `json:"HOST" validators:"required,minLength:3,maxLength:50" transformers:"trim"`
        AdminEmail string `json:"ADMIN_EMAIL" validators:"email" transformers:"trim,lowercase"`
   }
```

#### Built-in Transformers
| Transformer       | Description                        |
|-------------------|------------------------------------|
| `uppercase`       | Converts text to uppercase letters |
| `lowercase`       | Converts text to lowercase letters |
| `trim`            | Removes whitespace from both sides of the text |
| `trimLeft`        | Removes whitespace from the left side of the text |
| `trimRight`       | Removes whitespace from the right side of the text |
| `base64Encode`    | Encodes text to Base64 format      |
| `base64Decode`    | Decodes Base64 text to original format |
| `urlEncode`       | Encodes text to be URL-friendly    |
| `urlDecode`       | Decodes URL-encoded text to original format |

#### Why Choose xEnv? 🤔
With xEnv, handling environment variables becomes a breeze. You get the benefits of automatically loading and validating variables from `.env` files—functionality that Go doesn't provide out of the box. Whether your project is small or large, xEnv simplifies your configuration management, allowing you to focus more on developing great features.

#### License
xEnv is made available under the MIT License. This means you can use it freely in your personal and commercial projects. For more details, see the LICENSE file in the repository.
