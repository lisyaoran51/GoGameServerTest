
protoc -I=. --go_out=. protobuf/common.proto
protoc -I=. --go_out=. protobuf/common_game.proto
protoc -I=. --go_out=. protobuf/diamonds/common.proto
protoc -I=. --go_out=. protobuf/diamonds/game.proto


protoc -I=. --go_out=. protobuf/gate.proto
protoc -I=. --go_out=. protobuf/gate_game.proto
protoc -I=. --go_out=plugins=grpc:. common.proto

cp -r github.com/lisyaoran51/GoGameServerTest/protobuf/* ./protobuf/
rm -r github.com