import 'package:frontend/src/models/user_details.dart';
import 'package:frontend/src/resources/user_api_provider.dart';

class Repository {
  final userApiProvider = UserApiProvider();

  Future<bool> signUpUser(Map<String, dynamic> userCredentials) =>
      userApiProvider.signUpUser(userCredentials);

  Future<bool> signInUser(Map<String, dynamic> userCredentials) =>
      userApiProvider.signInUser(userCredentials);

  Future<UserDetails> getUserDetails(String username) =>
      userApiProvider.getUserDetails(username);
}
