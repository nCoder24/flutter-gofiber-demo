import 'package:frontend/src/models/user_details.dart';
import 'package:frontend/src/resources/repository.dart';
import 'package:rxdart/rxdart.dart';

class Bloc {
  final _repository = Repository();
  final _userDetailsFetcher = PublishSubject<UserDetails>();

  Stream<UserDetails> get userDetails => _userDetailsFetcher.stream;

  signUpUser(Map<String, dynamic> userCredentials) async {
    return await _repository.signUpUser(userCredentials);
  }

  signInUser(Map<String, dynamic> userCredentials) async {
    return await _repository.signInUser(userCredentials);
  }

  getUserDetails(String username) async {
    UserDetails userDetails = await _repository.getUserDetails(username);
    _userDetailsFetcher.sink.add(userDetails);
  }

  dispose() {
    _userDetailsFetcher.close();
  }
}

final bloc = Bloc();
