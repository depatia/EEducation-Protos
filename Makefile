proto_grades:
	protoc -I proto/grades proto/grades/*.proto --go_out=./gen/grades/ --go_opt=paths=source_relative --go-grpc_out=./gen/grades/ --go-grpc_opt=paths=source_relative

proto_schedule:
	protoc -I proto/schedule proto/schedule/*.proto --go_out=./gen/schedule/ --go_opt=paths=source_relative --go-grpc_out=./gen/schedule/ --go-grpc_opt=paths=source_relative

proto_lessons:
	protoc -I proto/lessons proto/lessons/*.proto --go_out=./gen/lessons/ --go_opt=paths=source_relative --go-grpc_out=./gen/lessons/ --go-grpc_opt=paths=source_relative

proto_user:
	protoc -I proto/user proto/user/*.proto --go_out=./gen/user/ --go_opt=paths=source_relative --go-grpc_out=./gen/user/ --go-grpc_opt=paths=source_relative