import 'package:flutter/material.dart';
import 'package:frontend/src/bloc/bloc.dart';
import 'package:frontend/src/models/user_details.dart';

class ProfileHeader extends StatelessWidget {
  const ProfileHeader({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 50,
      color: Theme.of(context).colorScheme.secondaryContainer,
      child: StreamBuilder(
        stream: bloc.userDetails,
        builder: (context, AsyncSnapshot<UserDetails> snapshot) {
          debugPrint(snapshot.data?.username);
          if (snapshot.hasData) {
            return Text(snapshot.data!.username.toString());
          } else if (snapshot.hasError) {
            return Text(snapshot.error.toString());
          }
          return const Center(child: CircularProgressIndicator());
        },
      ),
    );
  }
}
