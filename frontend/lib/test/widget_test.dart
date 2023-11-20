import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:frontend/src/app.dart';
import 'package:frontend/src/ui/auth_form.dart';

void main() {
  testWidgets('Verify sign up sign in button present',
      (WidgetTester tester) async {
    //Arrange - Pump MyApp() widget to tester
    await tester.pumpWidget(const App());

    //Act - Find button by type
    var eb = find.byType(ElevatedButton);

    //Assert - Check that button widget is present
    expect(eb, findsNWidgets(2));
  });

  testWidgets(
      'Verify that the two input widgets are present on the auth form page',
      (WidgetTester tester) async {
    //Arrange - Pump AddUser() widget to tester
    await tester.pumpWidget(const MaterialApp(home: SignUpForm()));

    //Act - Find button by type
    var textField = find.byType(TextFormField);

    //Assert - Check that exactly 2 Text input widgets are present
    expect(textField, findsNWidgets(2));
  });

  testWidgets('Verify that submit button is present on the auth form page',
      (WidgetTester tester) async {
    //Arrange - Pump AddUser() widget to tester
    await tester.pumpWidget(const MaterialApp(home: SignUpForm()));

    //Act - Find button by type
    var textField = find.byType(ElevatedButton);

    //Assert - Check that exactly 2 Text input widgets are present
    expect(textField, findsNWidgets(1));
  });
}
