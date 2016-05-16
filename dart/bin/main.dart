import "package:test/test.dart" as test;

import "package:tetorica/http.dart" as http;
import "package:tetorica/net_dartio.dart" as netio;
import "dart:convert" as conv;

main(List<String> args) {
  print('Hello world! ${args}');
  String host = args[0];
  int port = int.parse(args[1]);
  int i = 0;
  test.group("", () {
    test.test("new", () async {
      var builder = new netio.TetSocketBuilderDartIO();
      var client = new http.HttpClient(builder);
      await client.connect(host, port);
      http.HttpClientResponse response = await client.post(
          "/user/new",
          conv.UTF8.encode(conv.JSON.encode({
            "name": "kyoro001", //
            "mail": "kyoro001@example.com", //
            "pass": "asdfasdf", //
          })));
      print("## ${i++} ## ${await response.body.getString()}");
      test.expect(true, true);
    });
    test.test("get", () async {
      var builder = new netio.TetSocketBuilderDartIO();
      var client = new http.HttpClient(builder);
      await client.connect(host, port);
      http.HttpClientResponse response = await client.post(
          "/user/get",
          conv.UTF8.encode(conv.JSON.encode({
            "name": "kyoro001", //
          })));
      print("## ${i++} ## ${await response.body.getString()}");
      test.expect(true, true);
    });
    test.test("get error", () async {
      var builder = new netio.TetSocketBuilderDartIO();
      var client = new http.HttpClient(builder);
      await client.connect(host, port);
      http.HttpClientResponse response = await client.post(
          "/user/get",
          conv.UTF8.encode(conv.JSON.encode({
            "name": "kyoro002", //
          })));
      print("## ${i++} ## ${await response.body.getString()}");
      test.expect(true, true);
    });
    //
    test.test("get from mail", () async {
      var builder = new netio.TetSocketBuilderDartIO();
      var client = new http.HttpClient(builder);
      await client.connect(host, port);
      http.HttpClientResponse response = await client.post(
          "/user/mail/getUser",
          conv.UTF8.encode(conv.JSON.encode({
            "mail": "kyoro001@example.com", //
          })));
      print("## ${i++} ## ${await response.body.getString()}");
      test.expect(true, true);
    });

    test.test("update mail", () async {
      var builder = new netio.TetSocketBuilderDartIO();
      var client = new http.HttpClient(builder);
      await client.connect(host, port);
      http.HttpClientResponse response = await client.post(
          "/user/updateMail",
          conv.UTF8.encode(conv.JSON.encode({
            "name": "kyoro001", //
            "mail": "kyoro001b@example.com", //
          })));
      print("## ${i++} ## ${await response.body.getString()}");
      test.expect(true, true);
    });

    test.test("login logout", () async {
      String loginId = "";
      {
        var builder = new netio.TetSocketBuilderDartIO();
        var client = new http.HttpClient(builder);
        await client.connect(host, port);
        http.HttpClientResponse response = await client.post(
            "/user/login",
            conv.UTF8.encode(conv.JSON.encode({
              "name": "kyoro001", //
              "pass": "asdfasdf", //
            })), header: {
              "User-Agent":"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11"});
        print("## ${i++} ## ${await response.body.getString()}");
        var map = conv.JSON.decode(await response.body.getString());
        loginId = map["loginId"];
        test.expect(true, true);
      }

      {
        var builder = new netio.TetSocketBuilderDartIO();
        var client = new http.HttpClient(builder);
        await client.connect(host, port);
        http.HttpClientResponse response = await client.post(
            "/user/logout",
            conv.UTF8.encode(conv.JSON.encode({
              "name": "kyoro001", //
              "loginId": loginId, //
            })));
        print("## ${i++} ## ${await response.body.getString()}");
        test.expect(true, true);
      }
    });
    //

    test.test("delete", () async {
      {
        var builder = new netio.TetSocketBuilderDartIO();
        var client = new http.HttpClient(builder);
        await client.connect(host, port);
        http.HttpClientResponse response = await client.post(
            "/user/delete",
            conv.UTF8.encode(conv.JSON.encode({
              "name": "kyoro001", //
              "pass": "asdfasdf", //
            })));
        print("## ${i++} ## ${await response.body.getString()}");
        test.expect(true, true);
      }
    });
  });
}
