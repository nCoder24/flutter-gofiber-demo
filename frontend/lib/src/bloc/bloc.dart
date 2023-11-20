import 'package:frontend/src/models/user_details.dart';
import 'package:frontend/src/resources/repository.dart';
import 'package:rxdart/rxdart.dart';

class Bloc {
  final _repository = Repository();
  final _userDetailsFetcher = PublishSubject<UserDetails>();

  Stream<UserDetails> get userDetails => _userDetailsFetcher.stream;

  signUpUser(Map userCredentials) async {
    UserDetails userDetails = await _repository.signUpUser(userCredentials);
    _userDetailsFetcher.sink.add(userDetails);
  }

  signInUser(Map userCredentials) async {
    UserDetails userDetails = await _repository.signInUser(userCredentials);
    _userDetailsFetcher.sink.add(userDetails);
  }

  dispose() {
    _userDetailsFetcher.close();
  }
}

final bloc = Bloc();
