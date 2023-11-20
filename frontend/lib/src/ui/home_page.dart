import 'package:flutter/material.dart';
import 'package:frontend/src/bloc/bloc.dart';
import 'package:frontend/src/ui/auth/sign_in_form.dart';
import 'package:frontend/src/ui/auth/sign_up_form.dart';
import 'package:frontend/src/ui/profile_page.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Demo App"),
      ),
      body: Container(
        color: Theme.of(context).colorScheme.primary,
        child: Center(
          child: Container(
            height: 200,
            width: 300,
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(10),
              color: Theme.of(context).colorScheme.primaryContainer,
            ),
            child: Padding(
              padding: const EdgeInsets.all(10.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.center,
                mainAxisSize: MainAxisSize.min,
                children: <Widget>[
                  ElevatedButton(
                    onPressed: () => {
                      showDialog(
                          context: context,
                          barrierDismissible: false,
                          builder: (context) => const SignUpForm()).then(
                        (value) => {
                          debugPrint(value.toString()),
                          if (value != null) bloc.signUpUser(value),
                        },
                      )
                    },
                    child: const Text("Sign Up"),
                  ),
                  const SizedBox(
                    width: 20,
                  ),
                  ElevatedButton(
                    onPressed: () => {
                      showDialog(
                          context: context,
                          barrierDismissible: false,
                          builder: (context) => const SignInForm()).then(
                        (value) => {
                          debugPrint(value.toString()),
                          if (value != null)
                            {
                              bloc.signInUser(value),
                              Navigator.push(
                                context,
                                MaterialPageRoute(
                                    builder: (context) => const ProfilePage()),
                              )
                            }
                        },
                      )
                    },
                    child: const Text("Sign In"),
                  )
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
