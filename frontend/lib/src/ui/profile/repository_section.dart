import 'package:flutter/material.dart';
import 'package:frontend/src/bloc/bloc.dart';
import 'package:frontend/src/models/user_details.dart';

class RepositorySection extends StatelessWidget {
  const RepositorySection({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
        color: Theme.of(context).colorScheme.primary,
        child: StreamBuilder(
          stream: bloc.userDetails,
          builder: (context, AsyncSnapshot<UserDetails> snapshot) {
            if (snapshot.hasData) {
              var repos = snapshot.data!.repos;
              debugPrint(repos.toString());
              return Column(children: createRepos(repos, context));
            } else if (snapshot.hasError) {
              return Text(snapshot.error.toString());
            }
            return const Center(child: LinearProgressIndicator());
          },
        ));
  }

  List<Widget> createRepos(List repos, BuildContext context) {
    var list = <Widget>[];
    for (var repo in repos) {
      list.add(ListTile(
        title: Text(repo.name),
        leading: const Icon(Icons.folder),
        textColor: Theme.of(context).secondaryHeaderColor,
        iconColor: Theme.of(context).secondaryHeaderColor,
      ));
    }

    return list;
  }
}
