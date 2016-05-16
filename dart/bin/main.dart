import "package:test/test.dart" as test;

import "package:tetorica/http.dart" as http;
import "package:tetorica/net_dartio.dart" as netio;
import "dart:convert" as conv;

main(List<String> args) {
  print('Hello world! ${args}' );
  String host = args[0];
  int port = int.parse(args[1]);
  int i=0;
  test.group("", (){
    test.test("test1", () async {
      var builder = new netio.TetSocketBuilderDartIO();
      var client = new http.HttpClient(builder);
      await client.connect(host, port);
      http.HttpClientResponse response = await client.post("/user/new", conv.UTF8.encode(conv.JSON.encode({
        "name":"kyoro001",//
        "mail":"kyoro001@example.com",//
        "pass":"asdfasdf",//
      })));
      print("## ${i++} ## ${await response.body.getString()}");
      test.expect(true, true);
    });
    test.test("test2a", () async {
      var builder = new netio.TetSocketBuilderDartIO();
      var client = new http.HttpClient(builder);
      await client.connect(host, port);
      http.HttpClientResponse response = await client.post("/user/get", conv.UTF8.encode(conv.JSON.encode({
        "name":"kyoro001",//
      })));
      print("## ${i++} ## ${await response.body.getString()}");
      test.expect(true, true);
    });
    test.test("test2b", () async {
      var builder = new netio.TetSocketBuilderDartIO();
      var client = new http.HttpClient(builder);
      await client.connect(host, port);
      http.HttpClientResponse response = await client.post("/user/get", conv.UTF8.encode(conv.JSON.encode({
        "name":"kyoro002",//
      })));
      print("## ${i++} ## ${await response.body.getString()}");
      test.expect(true, true);
    });
    //
    test.test("test3", () async {
      var builder = new netio.TetSocketBuilderDartIO();
      var client = new http.HttpClient(builder);
      await client.connect(host, port);
      http.HttpClientResponse response = await client.post("/user/mail/getUser", conv.UTF8.encode(conv.JSON.encode({
        "mail":"kyoro001@example.com",//
      })));
      print("## ${i++} ## ${await response.body.getString()}");
      test.expect(true, true);
    });

    test.test("test4", () async {
      var builder = new netio.TetSocketBuilderDartIO();
      var client = new http.HttpClient(builder);
      await client.connect(host, port);
      http.HttpClientResponse response = await client.post("/user/updateMail", conv.UTF8.encode(conv.JSON.encode({
        "name":"kyoro001",//
        "mail":"kyoro001b@example.com",//
      })));
      print("## ${i++} ## ${await response.body.getString()}");
      test.expect(true, true);
    });
  });
}
