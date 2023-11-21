import 'dart:async';
import 'package:flutter/material.dart';
import 'package:frontend/src/models/user_details.dart';
import 'package:http/http.dart' show Client;
import 'dart:convert';

class UserApiProvider {
  Client client = Client();

  Future<bool> signUpUser(Map userCredentials) async {
    debugPrint("Signing Up the User");
    Uri uri = Uri.parse("http://localhost:3030/user");
    final response = await client.post(uri);

    if (response.statusCode == 200) {
      return true;
    }

    throw Exception("User Already Exists");
  }

  Future<bool> signInUser(userCredentials) async {
    debugPrint("Signing In the User");
    var headers = {
      "Access-Control-Allow-Origin": "*",
      "Content-Type": "application/json",
      "Accept": "*/*"
    };
    Uri uri = Uri.parse("http://localhost:3030/user/signin");
    final response = await client.post(uri,
        headers: headers, body: jsonEncode(userCredentials));

    if (response.statusCode == 200) {
      return true;
    }

    throw Exception("User Doesn't Exist");
  }

  Future<UserDetails> getUserDetails(String username) async {
    Uri uri = Uri.parse("http://localhost:3030/user/$username");

    final response = await client.get(uri);

    if (response.statusCode == 200) {
      return UserDetails.fromJson(json.decode(response.body));
    }

    throw Exception("User Doesn't Exist");
  }
}
