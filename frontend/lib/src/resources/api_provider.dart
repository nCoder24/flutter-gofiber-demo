import 'dart:async';
import 'package:flutter/material.dart';
import 'package:http/http.dart' show Client;
import 'dart:convert';

class ApiProvider {
  Client client = Client();
  final _apiKey = "api_key";

  Future<int> validateUserCredentials(Map userDetails) async {
    debugPrint("entered validation");
    Uri uri = Uri.parse("http://localhost:8000/validate?api_key=$_apiKey");
    final response = await client.get(uri);

    debugPrint(response.body.toString());
    if (response.statusCode == 200) {
      Map body = json.decode(response.body);
      return body["userId"];
    }

    throw Exception("Invalid User Credentials");
  }


}
