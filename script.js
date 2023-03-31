import grpc from 'k6/net/grpc';

const client = new grpc.Client();
// Download addsvc.proto for https://grpcb.in/, located at:
// https://raw.githubusercontent.com/moul/pb/master/addsvc/addsvc.proto
// and put it in the same folder as this script.
client.load(['proto'], 'genId.proto');

export default () => {

  client.connect('127.0.0.1:8000', { 
    timeout: '5s', 
    plaintext: true, // ko tls
  });
  // const response = client.invoke('trungtn.grpc.GenIDService/GenIDWithSnowflake', {});

  // goi ham hoac api
  client.invoke('trungtn.grpc.GenIDService/GenIDWithUUID', {});

  // call trong vong lap
  // for(let i=0; i<= 500000; i++) {
  //   client.invoke('trungtn.grpc.GenIDService/GenIDWithUUID', {});
  // }

  // console.log("snowflake: ", response.message.id); // should print 3

  client.close();
};
