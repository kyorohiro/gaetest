import "package:test/test.dart" as test;

main(List<String> args) {
  print('Hello world!');
  test.group("", (){
    test.test("test", (){
      test.expect(true, true);
    });
  });
}
