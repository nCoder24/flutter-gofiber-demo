import 'package:flutter/material.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Demo App"),
      ),
      body: Container(
        color: Theme.of(context).colorScheme.primaryContainer,
        child: const Column(
          children: [
            TextField(),
            TextField(),
          ],
        ),
      ),
    );
  }
}
