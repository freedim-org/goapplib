# goapplib

Provide a library that can be called by `mobile` or `desktop` applications. It starts a local TCP server to facilitate communication between your application and a backend written in `Golang`.

## Usage

### SDK

You must refer to the example code in [example](cmd/cgo/main.go) to write your SDK code.

### Application

```dart
import 'dart:ffi' as ffi;
import 'dart:io';

var responseMap = <String, Completer>{};

void onReceiveResponse(List<int> data) {
    String str = String.fromCharCodes(data);
    var map = json.decode(str);
    String traceId = map['traceId'];
    Completer completer = responseMap[traceId];
    if (completer != null) {
        completer.complete(map['data']);
    }
}

Map<String, dynamic> onReceiveRequest(List<int> data) {
    // json data to map
    String str = String.fromCharCodes(data);
    var map = json.decode(str);
    return {
        'code': 0,
        'traceId': map['traceId'],
        'data': map['data'],
    };
}

Map<String, dynamic> sendRequest(Socket socket, Map<String, dynamic> data) async {
    // 1. generate traceId
    String traceId = Uuid().v4();
    // 2. create completer
    Completer completer = Completer();
    // 3. add completer to responseMap
    responseMap[traceId] = completer;
    // 4. send request
    Map<String, dynamic> request = {
        'traceId': traceId,
        'data': data,
    };
    String requestStr = json.encode(request);
    List<int> requestBytes = utf8.encode(requestStr);
    int dataLength = requestBytes.length;
    List<int> dataLengthBytes = [
        dataLength >> 24 & 0xFF,
        dataLength >> 16 & 0xFF,
        dataLength >> 8 & 0xFF,
        dataLength & 0xFF,
    ];
    List<int> message = [
        ...dataLengthBytes,
        0,
        ...requestBytes,
    ];
    socket.add(message);
    // 5. wait for response
    Map<String, dynamic> response = await completer.future;
    // 6. remove completer from responseMap
    responseMap.remove(traceId);
    // 7. return response
    return response;
}

void main() async {
    // 1. 加载动态库
    ffi.DynamicLibrary dylib = ffi.DynamicLibrary.open('libgoapplib.dylib');
    // 2. Get Local Tcp Address
    ffi.Pointer<Utf8> Function() getLocalTcpAddress =
        dylib.lookup<ffi.NativeFunction<ffi.Pointer<Utf8> Function()>>(
            'Address').asFunction();
    ffi.Pointer<Utf8> address = getLocalTcpAddress();
    String localTcpAddress = ffi.Utf8.fromUtf8(address);
    print('localTcpAddress: $localTcpAddress');
    // 3. get host and port
    List<String> hostAndPort = localTcpAddress.split(':');
    String host = hostAndPort[0];
    int port = int.parse(hostAndPort[1]);
    // 4. connect to local tcp server
    Socket socket = await Socket.connect(host, port);
    // 5. set listener
    var dataBuffer = <int>[];
    socket.listen((List<int> event) {
        // message: [4]DataLength + [1]IsResponse + [n]Data
        dataBuffer.addAll(event);
        while (dataBuffer.length >= 5) {
            int dataLength = dataBuffer[0] << 24 |
                dataBuffer[1] << 16 |
                dataBuffer[2] << 8 |
                dataBuffer[3];
            if (dataBuffer.length >= dataLength + 4) {
                bool isResponse = dataBuffer[4] == 1;
                List<int> data = dataBuffer.sublist(5, dataLength + 4);
                dataBuffer = dataBuffer.sublist(dataLength + 4);
                if (isResponse) {
                    // response
                    onReceiveResponse(data);
                } else {
                    // request
                    print('request: ${String.fromCharCodes(data)}');
                    // response
                    onReceiveRequest(data);
                }
            } else {
                break;
            }
        }
    });
    // 6. test send request
    var resp = await sendRequest(socket, {
        'method': 'test',
        'data': 'hello',
    });
    print('resp: $resp');
    // 7. close socket
    socket.close();
}
```