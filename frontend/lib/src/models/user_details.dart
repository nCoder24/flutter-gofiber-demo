import 'package:flutter/material.dart';

class UserDetails {
  late String _username;
  late List<_Repo> _repos;

  UserDetails.fromJson(Map<String, dynamic> parsedJson) {
    debugPrint(parsedJson['repos'].length.toString());
    _username = parsedJson['username'];
    List<_Repo> temp = [];
    for (int index = 0; index < parsedJson['repos'].length; index++) {
      _Repo repo = _Repo(parsedJson['repos'][index]);
      temp.add(repo);
    }

    _repos = temp;
  }

  List<_Repo> get repos => _repos;

  String get username => _username;
}

class _Repo {
  late String _name;
  late Uri _uri;

  _Repo(Map<String, dynamic> repoDetails) {
    _name = repoDetails['name'];
    _uri = Uri.parse(repoDetails['uri']);
  }

  String get name => _name;

  Uri get uri => _uri;
}
