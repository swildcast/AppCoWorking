# Code Citations

## License: desconocido
https://github.com/ElayadeIsmail/go-auth/tree/d8c8183c421121b9253c2f948e4ad0f8650e8b43/database/connection.go

```go
package config

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB = database
}
```


## License: desconocido
https://github.com/VLaroye/gasso-back/tree/5f512ca5b302fdf27944e29412ee4e762506e051/cmd/gasso/main.go

```
"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.
```


## License: desconocido
https://github.com/dmitryweiner/lectures/tree/b45bdad3889119bcea9112ddcc8171d14f6a15cf/src/Test%20React%20components.md

```
```javascript
     import { render, screen } from '@testing-library/react';
     import App from './App';

     test('renders learn react link', () => {
       render(<App />);
       const linkElement = screen
```


## License: desconocido
https://github.com/kemo12/google-calendar/tree/3ddadec46983579400ef91c58ba1bf28cf098bcf/src/App.test.js

```
} from '@testing-library/react';
     import App from './App';

     test('renders learn react link', () => {
       render(<App />);
       const linkElement = screen.getByText(/learn react/i)
```


## License: Apache_2_0
https://github.com/ketto000/React-Management-Tutorial/tree/dbc242a85c492b39e4c638cf69506e6ae7822ce1/src/App.test.js

```
'@testing-library/react';
     import App from './App';

     test('renders learn react link', () => {
       render(<App />);
       const linkElement = screen.getByText(/learn react/i);
       expect
```

