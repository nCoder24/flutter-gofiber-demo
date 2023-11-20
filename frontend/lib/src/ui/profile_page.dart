import 'package:flutter/material.dart';
import 'package:frontend/src/ui/profile/profile_header.dart';
import 'package:frontend/src/ui/profile/repository_section.dart';

class ProfilePage extends StatelessWidget {
  const ProfilePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: const Text("Demo App"),
        ),
        body: Container(
          color: Theme.of(context).colorScheme.primary,
          child: const Column(
            children: <Widget>[ProfileHeader(), RepositorySection()],
          ),
        ));
  }
}
