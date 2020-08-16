## API

## BaseURL:
    (local) http://127.0.0.1:7481
    
#### Ping
   server health check
   - method: ```GET```
   - path: ```/ping```
   - response:
       ``` 
        {
            "message": "pong"
        }
       ```  


#### Version
   server version check
   - method: ```GET```
   - path: ```/version```
   - response:
       ``` 
       {
            "version": "v1"
       }
       ```
     
     
## 관리자 (admins)


#### AdminRegister
   관리자 계정 가입
   - method: ```POST```
   - path: ```/admins/register```
   - request:
       ``` 
        adminID     string
        password    string
       ```
   - response:
       ```
        adminUUID   string
       ```
   - error case:
        - request field 누락
            ```
                statusCode: 400
                {
                    "message": "Key: {FieldName} Error:Field validation for {FieldName} failed on the 'required' tag"
                }
            ```
        - adminID 중복
            ```
                statusCode: 409
                {
                    "message": "ID already exists"
                }
            ```


#### Login
   관리자 로그인
   - method: ```POST```
   - path: ```/admins/login```
   - request:
       ```
        adminID     string
        password    string
       ```
   - response:
       ``` 
        accessToken     string
       ```
   - error case:
        - request field 누락
            ```
                statusCode: 400
                {
                    "message": "Key: {FieldName} Error:Field validation for {FieldName} failed on the 'required' tag"
                }
            ```
        - 잘못된 adminID 입력
            ```
                statusCode: 401
                {
                    "message": "unauthorized.admin not found: {adminID}"
                }
            ```
        - 잘못된 password 입력
            ```
                statusCode: 401
                {
                    "message": "unauthorized.crypto/bcrypt: hashedPassword is not the hash of the given password"
                }
            ```
          
          
#### GetUserList
   사용자 가입 목록 조회
   - method: ```GET```
   - path: ```/admins/management/:page`
   - header: ```accessToken```
   - response:
        ```
         totalNum          int
         users             []model.User
        ```
   - error case:
        - request header에 accessToken 누락
            ```
                statusCode: 401
                {
                    "error": "missing accessToken at header"
                }
            ``` 
        - 요청하는 페이지의 인덱스가 totalNum 보다 클 경우
           ``` 
                statusCode: 404
                {
                    "message": "not found any elements in this page"
                }
           ```
          
          
#### ApproveRegisteration
   사용자 가입 관리 - 승인
   - method: ```POST```
   - path: ```/admins/management```
   - header: ```accessToken```
   - request:
       ```
        userUUID            string
       ```
   - response:
       ```
        userStatus          string
        userUUID            string
       ```
   - error case:
        - request header에 accessToken 누락
            ```
                statusCode: 401
                {
                 "error": "missing accessToken at header"
                }
            ```
        - request field 누락
            ```
                statusCode: 400
                {
                    "message": "Key: {FieldName} Error:Field validation for {FieldName} failed on the 'required' tag"
                }
            ```
        - 잘못된 상태 변경 (ex. 거절 -> 승인 / 승인 -> 승인)
            ``` 
                statusCode: 400
                {
                    "message": "invalid status change request"
                }
            ```
        - 해당하는 user 탐색 실패
            ``` 
                statusCode: 404
                {
                    "message": "user not found: {userUUID}"
                }
            ```


#### RejectApplication
   사용자 가입 관리 - 거절
   - method: ```PATCH```
   - path: ```/admins/management```
   - header: ```accessToken```
   - request:
       ```
        userUUID            string
       ```
   - response:
       ```
        userStatus          string
        userUUID            string
       ```
   - error case:
        - request header에 accessToken 누락
            ```
                statusCode: 401
                {
                 "error": "missing accessToken at header"
                }
            ```
        - request field 누락
            ```
                statusCode: 400
                {
                    "message": "Key: {FieldName} Error:Field validation for {FieldName} failed on the 'required' tag"
                }
            ```
        - 잘못된 상태 변경 (ex. 승인 -> 거절 / 거절 -> 거절)
            ``` 
                statusCode: 400
                {
                    "message": "invalid status change request"
                }
            ```
        - 해당하는 user 탐색 실패
            ``` 
                statusCode: 404
                {
                    "message": "users not found: {userUUID}"
                }
            ```



## 사용자 (users)

 
#### UserRegister
   사용자 계정 가입
   - method: ```POST```
   - path: ```/users/register```
   - request:
       ``` 
        name        string
        phone       string
        email       string
        password    string
       ```
   - response:
       ```
        userUUID   string
       ```
   - error case:
        - request field 누락
            ```
                statusCode: 400
                {
                    "message": "Key: {FieldName} Error:Field validation for {FieldName} failed on the 'required' tag"
                }
            ```
        - phone 중복
            ```
                statusCode: 409
                {
                    "message": "phone already exists"
                }
            ```

#### SignIn
   사용자 로그인
   - method: ```POST```
   - path: ```/users/sign-in```
   - request:
        ```
         phone       string
         password    string
        ```
   - response:
        ``` 
         accessToken     string
        ```
   - error case:
        - request field 누락
           ```
               statusCode: 400
               {
                   "message": "Key: {FieldName} Error:Field validation for {FieldName} failed on the 'required' tag"
               }
           ```
        - 잘못된 phone 입력
           ```
               statusCode: 401
               {
                   "message": "unauthorized.users not found: {phone}"
               }
           ```
        - 잘못된 password 입력
           ```
               statusCode: 401
               {
                   "message": "unauthorized.crypto/bcrypt: hashedPassword is not the hash of the given password"
               }
           ```
          

#### GetUserInfo
   사용자 정보 조회
   - method: ```GET```
   - path: ```/users/info```
   - header: ```accessToken```
   - response:
        ``` 
         user     model.User
        ```
   - error case:
        - 사용자 탐색 실패
            ``` 
                statusCode: 404
                {
                    "message": "users not found: {userUUID}"
                }
            ```
          

#### PatchUserInfo
   사용자 정보 수정
   - method: ```PATCH```
   - path: ```/users/info```
   - header: ```accessToken```
   - request: 
       ``` 
        name        string
        phone       string
        email       string
       ```
   - response:
       ``` 
        user     model.User
       ``` 
   - error case:
       - 사용자 탐색 실패
           ``` 
               statusCode: 404
               {
                   "message": "users not found: {userUUID}"
               }
           ```