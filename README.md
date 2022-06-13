# robbo_student_personal_account

### cloning repository
git clone https://github.com/skinnykaen/robbo_student_personal_account.git
### open folder
cd robbo_student_personal_account
### download dependency and libraries 
go mod download
### setup postgres
docker-compose up -d
### Run app
cd cmd
go run main.go