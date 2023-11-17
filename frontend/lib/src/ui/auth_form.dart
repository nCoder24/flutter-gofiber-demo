import 'package:flutter/material.dart';

class AuthForm extends StatefulWidget {
  const AuthForm({super.key});

  @override
  State<StatefulWidget> createState() => AuthFormState();
}

class AuthFormState extends State<AuthForm> {
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    return Dialog(
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(10)),
      backgroundColor: Theme.of(context).colorScheme.primaryContainer,
      elevation: 16,
      child: SizedBox(
        height: 300,
        width: 350,
        child: Padding(
          padding: const EdgeInsets.all(10.0),
          child: Form(
            key: _formKey,
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                const Text(
                  "Details",
                  style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                ),
                const SizedBox(
                  height: 15,
                ),
                TextFormField(
                  key: const Key("username"),
                  decoration: const InputDecoration(
                    border: OutlineInputBorder(),
                    hintText: 'Enter Username',
                  ),
                  validator: (String? value) {
                    if (value == null || value.isEmpty) {
                      return 'Please enter user name';
                    }
                    return null;
                  },
                ),
                const SizedBox(
                  height: 20,
                ),
                TextFormField(
                  key: const Key("password"),
                  obscureText: true,
                  decoration: const InputDecoration(
                    border: OutlineInputBorder(),
                    hintText: 'Enter Password',
                  ),
                  validator: (String? value) {
                    if (value == null || value.isEmpty) {
                      return 'Please enter password';
                    }
                    return null;
                  },
                ),
                const SizedBox(
                  height: 20,
                ),
                ElevatedButton(
                  onPressed: () {
                    if (_formKey.currentState!.validate()) {
                      _formKey.currentState?.reset();
                      Navigator.pop(context);
                      ScaffoldMessenger.of(context).showSnackBar(
                        const SnackBar(content: Text('Processing Data')),
                      );
                    }
                  },
                  child: const Text("Submit"),
                )
              ],
            ),
          ),
        ),
      ),
    );
  }
}
