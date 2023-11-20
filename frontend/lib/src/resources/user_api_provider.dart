import 'dart:async';
import 'package:flutter/material.dart';
import 'package:frontend/src/models/user_details.dart';
import 'package:http/http.dart' show Client;
import 'dart:convert';

class UserApiProvider {
  Client client = Client();

  Future<UserDetails> signUpUser(Map userCredentials) async {
    debugPrint("Signing Up the User");
    Uri uri = Uri.parse("http://localhost:3000/user");
    final response = await client.post(uri);

    debugPrint(response.body.toString());
    if (response.statusCode == 200) {
      return UserDetails.fromJson(json.decode(response.body));
    }

    throw Exception("Invalid User Credentials");
  }

  Future<UserDetails> signInUser(userCredentials) async {
    debugPrint("Signing In the User");
    Uri uri = Uri.parse("http://localhost:3000/user/signin");
    final response = await client.post(uri);

    debugPrint(response.body.toString());
    if (response.statusCode == 200) {
      return UserDetails.fromJson(json.decode(response.body));
    }

    throw Exception("User Doesn't Exist");
  }
}
