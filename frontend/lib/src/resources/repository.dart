import 'package:frontend/src/models/user_details.dart';
import 'package:frontend/src/resources/user_api_provider.dart';

class Repository {
  final userApiProvider = UserApiProvider();

  Future<UserDetails> signUpUser(userCredentials) => userApiProvider.signUpUser(userCredentials);

  Future<UserDetails> signInUser(userCredentials) => userApiProvider.signInUser(userCredentials);
}
