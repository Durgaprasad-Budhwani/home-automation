module github.com/durgaprasad-budhwani/home-automation/backend

go 1.12

require (
	github.com/durgaprasad-budhwani/aws-iot-device-sdk-go v0.0.0-20191103115147-0cd7a63e254a
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-contrib/gzip v0.0.1
	github.com/gin-gonic/gin v1.9.1
	github.com/jasonlvhit/gocron v0.0.0-20191021204008-47e27a9e0dc7
	github.com/jinzhu/gorm v1.9.11
	github.com/rs/zerolog v1.16.0
	github.com/spf13/viper v1.5.0
	gobot.io/x/gobot v1.14.0
)

replace github.com/ugorji/go => github.com/ugorji/go/codec v1.1.7
