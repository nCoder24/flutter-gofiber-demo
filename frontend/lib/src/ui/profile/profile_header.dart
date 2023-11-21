import 'package:flutter/material.dart';
import 'package:frontend/src/bloc/bloc.dart';
import 'package:frontend/src/models/user_details.dart';
import 'package:frontend/src/ui/home_page.dart';

class ProfileHeader extends StatelessWidget {
  const ProfileHeader({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 40,
      width: MediaQuery.of(context).size.width,
      color: Theme.of(context).colorScheme.primaryContainer,
      child: StreamBuilder(
        stream: bloc.userDetails,
        builder: (context, AsyncSnapshot<UserDetails> snapshot) {
          if (snapshot.hasData) {
            var string = snapshot.data!.username.toString();
            return Row(children: [
              Text(
                "Welcome $string",
                style: const TextStyle(fontSize: 30),
              ),
              const SizedBox(
                width: 250,
              ),
              ElevatedButton.icon(
                onPressed: () => {},
                icon: const Icon(Icons.edit),
                label: const Text("Edit"),
              ),
              const SizedBox(
                width: 20,
              ),
              ElevatedButton(
                  onPressed: () => {
                        Navigator.pushReplacement(
                          context,
                          MaterialPageRoute(
                              builder: (context) => const HomePage()),
                        )
                      },
                  child: const Text("Log Out"))
            ]);
          } else if (snapshot.hasError) {
            return Text(snapshot.error.toString());
          }
          return const Center(child: CircularProgressIndicator());
        },
      ),
    );
  }
}
